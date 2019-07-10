package router

import( 
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"qilog"
)

func CreateInsertOneEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	dataToInsert, err := Parse(request, "single", []interface{}{"body"})	
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	insertedResult, err := collection.Insert(dataToInsert)
	if qilog.Check(err){
		qilog.Error("Error while inserting", err)
		response.Write([]byte("{\"Error while inserting\" : {\"" + err.Error() + "\"}"))
	}
	toWrite, _ := bson.MarshalExtJSON(insertedResult, true, true)
	response.Write(toWrite)	
	qilog.Debug(string(toWrite))
}

func CreateInsertManyEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	dataToInsert, err := Parse(request, "multiple", []interface{}{"body"})	
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	insertedResult, err := collection.Insert(dataToInsert)
	if qilog.Check(err){
		qilog.Error("Error while inserting", err)
		response.Write([]byte("{\"Error while inserting\" : {\"" + err.Error() + "\"}"))
	}	
	toWrite, _ := bson.MarshalExtJSON(insertedResult, true, true)
	response.Write(toWrite)	
	qilog.Debug(string(toWrite))
}