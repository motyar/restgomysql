// server.go
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run server.go

package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var apiport *int = flag.Int("apiport", 1234, "The port to listen on for the api.")
var dbhost = flag.String("dbhost", "localhost", "The mysql hostname/ip address.")
var dbport *int = flag.Int("dbport", 3306, "The mysql port number.")
var dbuser = flag.String("dbuser", "root", "The mysql username to use to access the database.")
var dbpass = flag.String("dbpass", "", "The mysql password to use to access the database.")
var dbname = flag.String("dbname", "test", "The mysql database name.")

type Panda struct {
	Id   int
	Name string
}

//Handle all requests
func Handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

// Respond to URLs of the form /generic/...
func APIHandler(response http.ResponseWriter, request *http.Request) {

	//Connect to database
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *dbuser, *dbpass, *dbhost, *dbport, *dbname)
	db, e := sql.Open("mysql", connectString)
	if e != nil {
		fmt.Print(e)
	}

	//set mime type to JSON
	response.Header().Set("Content-type", "application/json")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	//can't define dynamic slice in golang
	var result = make([]string, 1000)

	switch request.Method {
	case "GET":
		st, err := db.Prepare("select * from pandas limit 10")
		if err != nil {
			fmt.Print(err)
		}
		rows, err := st.Query()
		if err != nil {
			fmt.Print(err)
		}
		i := 0
		for rows.Next() {
			var name string
			var id int
			err = rows.Scan(&id, &name)
			panda := &Panda{Id: id, Name: name}
			b, err := json.Marshal(panda)
			if err != nil {
				fmt.Println(err)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	case "POST":
		name := request.PostFormValue("name")
		st, err := db.Prepare("INSERT INTO pandas(name) VALUES(?)")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(name)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]

	case "PUT":
		name := request.PostFormValue("name")
		id := request.PostFormValue("id")

		st, err := db.Prepare("UPDATE pandas SET name=? WHERE id=?")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(name, id)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]
	case "DELETE":
		id := strings.Replace(request.URL.Path, "/api/", "", -1)
		st, err := db.Prepare("DELETE FROM pandas WHERE id=?")
		if err != nil {
			fmt.Print(err)
		}
		res, err := st.Exec(id)
		if err != nil {
			fmt.Print(err)
		}

		if res != nil {
			result[0] = "true"
		}
		result = result[:1]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}

func main() {

	flag.Parse() // parse the command line args

	port := *apiport
	var err string
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
