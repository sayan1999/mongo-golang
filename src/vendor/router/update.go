package router

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"qilog"
)

func CreateUpdateOneEndpoint(response http.ResponseWriter, request *http.Request) {

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	updateSpec, err := Parse(request, "multiple", []interface{}{"body"})
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, err := collection.Update(updateSpec[0], updateSpec[1], 1)
	if qilog.Check(err){
		qilog.Error("Error while updating", err)
		response.Write([]byte("{\"Error while updating\" : {\"" + err.Error() + "\"}"))
	}
	toWrite, _ := bson.MarshalExtJSON(result, true, true)
	response.Write(toWrite)
	qilog.Debug(string(toWrite))
}

func CreateUpdateManyEndpoint(response http.ResponseWriter, request *http.Request) {

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	updateSpec, err := Parse(request, "multiple", []interface{}{"body"})
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, err := collection.Update(updateSpec[0], updateSpec[1], 2)
	if qilog.Check(err){
		qilog.Error("Error while updating", err)
		response.Write([]byte("{\"Error while updating\" : {\"" + err.Error() + "\"}"))
	}
	toWrite, _ := bson.MarshalExtJSON(result, true, true)
	response.Write(toWrite)
	qilog.Debug(string(toWrite))
}
