package router
import(
	"mongodb"
	"qilog"
	"net/http"
	"fmt"
)

func EstablishConn(response http.ResponseWriter, request *http.Request){

	var err error
	defer func (){
		err := recover()
		if qilog.Check(err){
			qilog.Error("Error countered in establishing connection", err)
		}
	}()
	response.Header().Set("content-type", "application/json")	
	params := request.URL.Query()

	hostQuery := params["host"]
	portQuery := params["port"]
	
	client, err = mongodb.Initialize(hostQuery[0], portQuery[0])
	qilog.Debug("Client details:", client)
	if qilog.Check(err){
		qilog.Error("Error found while connecting to mongo client:", err)
		response.Write([]byte("{\"Could not connect to\":  {\"Host\" : \"" + hostQuery[0] + "\", \"Port\" : \"" + portQuery[0] + "\"}}"))
	} else {
		qilog.Info("Connection established to mongo server")
		response.Write([]byte("{\"You are connected to\":  {\"Host\" : \"" + hostQuery[0] + "\", \"Port\" : \"" + portQuery[0] + "\"}}"))
	}
}

func CreateDBSelectEndPoint(response http.ResponseWriter, request *http.Request){
	
	if client.IsNIL() {
		qilog.Error("No connection to any mongo server.")
		response.Write([]byte("Please connect to a mongo server."))
		return
	}

	var err error
	defer func (){
		err := recover()
		if qilog.Check(err){
			qilog.Error("Error countered in getting collection", err, "for client:", fmt.Sprintf("%#v\t", client), "request:", fmt.Sprintf("%#v\n", *request))
		}
	}()
	response.Header().Set("content-type", "application/json")	
	params := request.URL.Query()

	databaseQuery := params["database"]
	collectionQuery := params["collection"]

  collection, err = client.GetCollection(databaseQuery[0], collectionQuery[0])
	qilog.Debug("Collection pointer:", collection)
  if qilog.Check(err){
		qilog.Error("Error found while creating or connecting to mongo collection for db:", databaseQuery[0], "collection:", collectionQuery[0], "with error as:", err)
		response.Write([]byte("{\"Could not connect to\":  {\"Database\" : \"" + databaseQuery[0] + "\", \"Collection\" : \"" + collectionQuery[0] + "\"}}"))
	} else {
		qilog.Info("Connection established to mongo colllection")
		response.Write([]byte("{\"You are connected to\":  {\"Database\" : \"" + databaseQuery[0] + "\", \"Collection\" : \"" + collectionQuery[0] + "\"}}"))
	}
}