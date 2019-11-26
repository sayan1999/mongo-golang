package router

import(
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"qilog"
)

func CreateDeleteOneEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	delCondition, err := Parse(request, "default", []interface{}{"filter"})
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, err := collection.Delete(delCondition[0], 1)
	if qilog.Check(err){
		qilog.Error("Error in deleting", err)
		response.Write([]byte("{\"Could not delete\":  {\"" + err.Error() + "\"}"))
	}	
	toWrite, _ := bson.MarshalExtJSON(result, true, true)
	response.Write(toWrite)	
	qilog.Debug(string(toWrite))
}

func CreateDeleteManyEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	delCondition, err := Parse(request, "default", []interface{}{"filter"})
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, err := collection.Delete(delCondition[0], 2)
	if qilog.Check(err){
		qilog.Error("Error in deleting", err)
		response.Write([]byte("{\"Could not delete\":  {\"" + err.Error() + "\"}"))
	}	
	toWrite, _ := bson.MarshalExtJSON(result, true, true)
	response.Write(toWrite)	
	qilog.Debug(string(toWrite))
}

func CreateDeleteAllEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	result, err := collection.Delete(bson.D{{}}, 2)	
	if qilog.Check(err){
		qilog.Error("Error in deleting", err)		
		response.Write([]byte("{\"Could not delete\":  {\"" + err.Error() + "\"}"))
	}	
	toWrite, _ := bson.MarshalExtJSON(result, true, true)
	response.Write(toWrite)	
	qilog.Debug(string(toWrite))
}