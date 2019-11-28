# Mongo Go Sever
Mongo Go Server is a server housing a Mongo Database with the server sided code written in Golang

## Directories
└── mongo-golang
    ├── frontend
    │   ├── APIDoc.yaml
    │   └── httpURLs.txt
    ├── README.md
    └── src
        ├── log.json
        ├── main.go
        └── vendor
            ├── endPoint
            │   └── endPoint.go
            ├── mongodb
            │   └── mongodb.go
            ├── qilog
            │   ├── errorHandling.go
            │   ├── functioncallpath.go
            │   └── logging.go
            └── router
                ├── connect.go
                ├── delete.go
                ├── find.go
                ├── handlers.go
                ├── insert.go
                ├── responseAndRequest.go
                └── update.go

## Dependencies
Install the following packages:
(make sure to export your GOPATH prior to installing the dependencies
and export GOPATH with the same values being assigned now)

```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo/options
go get github.com/op/go-logging 
go get github.com/gorilla/mux
```

## Directory organization
Expected file organization:

.../gocode/
       pkg/
       bin/
       src/
	 /mongo-golang/

considering .../gocode/src as your go workspace and your GOPATH set as ".../gocode"	

## Running the code 
Find swagger config file at mongo-golang/frontend/APIDoc.yaml
Navigate to .../gocode/src/mongo-golang/src on your terminal and run the following commands:

```bash
go build main.go
./main.exe
```
Note: Give permission to the firewall prompt if any appears.


Make http requests from your standalone local swagger or 
make curl requests from terminal as examples given at mongo-golang/frontend/httpURLs.txt:

```bash
curl -X GET "http://localhost:12345/connecttoserver?host=localhost&port=27017" -H "accept: application/json"

curl -X GET "http://localhost:12345/connecttodb?database=Practice&collection=Test1" -H "accept: application/json"

curl -X POST "http://localhost:12345/insertone" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"key1\": \"value1\", \"key2\": \"value2\"}"

curl -X POST "http://localhost:12345/insertmany" -H "accept: application/json" -H "Content-Type: application/json" -d "[ { \"key1\": \"value1\", \"key2\": \"value2\" }, { \"key1\": \"value1\", \"key2\": \"value2\" }]"

curl -X GET "http://localhost:12345/findone?filter=%7B%22key1%22%3A%20%22value1%22%2C%20%22key2%22%3A%20%22value2%22%7D&projection=%7B%22key1%22%3A%201%7D" -H "accept: application/json"

curl -X GET "http://localhost:12345/findall?filter=%7B%22key1%22%3A%20%22value1%22%2C%20%22key2%22%3A%20%22value2%22%7D&projection=%7B%22key1%22%3A%201%7D" -H "accept: application/json"

curl -X PUT "http://localhost:12345/updateone" -H "accept: application/json" -H "Content-Type: application/json" -d "[ { \"key1\": \"value1\", \"key2\": \"value2\" }, { \"$set\": { \"key1\": \"value2\", \"key2\": \"value1\" } }]"

curl -X PUT "http://localhost:12345/updatemany" -H "accept: application/json" -H "Content-Type: application/json" -d "[ { \"key1\": \"value1\", \"key2\": \"value2\" }, { \"$set\": { \"key1\": \"value2\", \"key2\": \"value1\" } }]"

curl -X DELETE "http://localhost:12345/deleteone?filter=%7B%22key1%22%3A%20%22value1%22%2C%20%22key2%22%3A%20%22value2%22%7D" -H "accept: application/json"

curl -X DELETE "http://localhost:12345/deletemany?filter=%7B%22key1%22%3A%20%22value1%22%2C%20%22key2%22%3A%20%22value2%22%7D" -H "accept: application/json"

curl -X DELETE "http://localhost:12345/deleteall" -H "accept: application/json"
```
