package router
import(
 	"net/http"
 	"qilog"
)

func CreateFindOneEndpoint(response http.ResponseWriter, request *http.Request){
	
	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
  findCondition, err := Parse(request, "default", []interface{}{"filter", "projection"})
  if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, errstring := collection.Find(findCondition[0], findCondition[1], 1)
	if errstring != "" {
		qilog.Error("Error while finding", errstring)
		response.Write([]byte("{\"Error while finding\" : \"" + errstring + "\"}"))
	}

	textResponse, err := writeFindResult(response, result)
	qilog.Debug(textResponse)
	if qilog.Check(err){
		qilog.Error("Couldn't unmarshall bson response from mongoDB:", err)
		response.Write([]byte("{\"Error while unmarshall bson response from mongoDB\" : {\"" + err.Error() + "\"}"))
	}
}

func CreateFindAllEndpoint(response http.ResponseWriter, request *http.Request){

	if collection.IsNIL() {
		qilog.Error("Collection handler is nil.")
		response.Write([]byte("Please connect to collection."))
		return
	}

	response.Header().Set("content-type", "application/json")
	findCondition, err := Parse(request, "default", []interface{}{"filter", "projection"})
	if qilog.Check(err){
		qilog.Error("Error while parsing request", err)
		response.Write([]byte("{\"Error while parsing request\" : {\"" + err.Error() + "\"}"))
	}

	result, errstring := collection.Find(findCondition[0], findCondition[1], 2)
	if errstring != "" {
		qilog.Error("Error while finding", errstring)
		response.Write([]byte("{\"Error while finding\" : \"" + errstring + "\"}"))
	}
	
	textResponse, err := writeFindResult(response, result)
	qilog.Debug(textResponse)
	if qilog.Check(err){
		qilog.Error("Couldn't unmarshall bson response from mongoDB:", err)
		response.Write([]byte("{\"Error while unmarshall bson response from mongoDB\" : {\"" + err.Error() + "\"}"))
	}
}