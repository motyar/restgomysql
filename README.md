#REST API in Golang with mySql Database

# Install go lang
# Install mysql

# Installation

        git clone https://github.com/motyar/restgomysql
        go get github.com/go-sql-driver/mysql
        cd restgomysql
        go run server.go

And open http://localhost:1234/api or http://<ip>:1234/api

Notes: This requires a valid mysql user account. It also requires a schema, which has not yet been provided.

# Nothing but (cute) Pandas

GET /api/ to get all the pandas.

POST /api/ to add new panda {name}

DELETE /api/panda_id to remove that one panda.

PUT /api/ to update details {id and name}



