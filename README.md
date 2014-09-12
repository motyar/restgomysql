#REST API in Golang with mySql Database

# Install go lang
# Install mysql

# Installation

        git clone https://github.com/motyar/restgomysql
        go get github.com/go-sql-driver/mysql
        cd restgomysql
        go run server.go

And open http://localhost:1234/api or http://<ip>:1234/api

Notes: This requires a valid mysql user account. It also requires a schema. For instance, to use database 'test', create the panda table by running something like this:
- mysql -uuser -p -Dtest < farm.sql

# Nothing but (cute) Pandas

GET /api/ to get all the pandas.
- curl http://localhost:1234/api/

POST /api/ to add new panda {name}
- curl -X POST -d 'name=new' http://localhost:1234/api/

DELETE /api/panda_id to remove that one panda.
- curl -XDELETE "http://localhost:1234/api/1"

PUT /api/ to update details {id and name}
- curl -X PUT -d 'id=1&name=Cool' http://localhost:1234/api/



