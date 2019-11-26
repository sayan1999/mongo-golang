package router

import(
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"io/ioutil"
	"qilog"
	"encoding/json"
)


func Parse(request *http.Request, ipCount string, method []interface{}) ([]interface{}, error){

	var bsonDocArray []interface{}
	if (method[0].(string) == "body") {	
		responseData, err := ioutil.ReadAll(request.Body)
		if qilog.Check(err) {
	        return nil, err
	  }		
		if ipCount == "single" {			
			var bsonDoc interface{}
			err = bson.UnmarshalExtJSON(responseData, true, &bsonDoc)
			if qilog.Check(err) {
	        return nil, err
	  	}	
			bsonDocArray = append(bsonDocArray, bsonDoc)

		} else if ipCount == "multiple" {
			err = bson.UnmarshalExtJSON(responseData, true, &bsonDocArray)
			if qilog.Check(err) {
	        return nil, err
	  	}	
		}
		return bsonDocArray, nil

	} else {
		params := request.URL.Query()

		firstQuery := params[method[0].(string)]		
	  var first interface{}
	  err := bson.UnmarshalExtJSON([]byte(firstQuery[0]), true, &first)
		if qilog.Check(err) {
	        return nil, err
	  }	
	  bsonDocArray = append(bsonDocArray, first)

	  if len(params) == 1 {
			return bsonDocArray, nil
	  }
    secondQuery := params[method[1].(string)]
    var second interface{}
   	err = bson.UnmarshalExtJSON([]byte(secondQuery[0]), true, &second)
		if qilog.Check(err) {
	    return nil, err
	  }	
    bsonDocArray = append(bsonDocArray, second)
    return bsonDocArray, nil
	}
}

func writeInsertResult(response http.ResponseWriter, id []interface{}) (string, error){

	response.Header().Set("content-type", "application/json")
	jsonString, err := json.Marshal(id)
	if qilog.Check(err) {
		return "", err
  }

	var buf []byte
	buf = append(buf, "{\"Inserted object ID \" : "...)
	buf = append(buf, jsonString...)
	buf = append(buf, '}')
	response.Write(buf)	

	var textString string
	err = json.Unmarshal(buf, &textString) 
	qilog.Error(err)
	return textString, nil
}

func writeFindResult(response http.ResponseWriter, result []interface{}) (string, error){

  response.Header().Set("content-type", "application/json")
	var jsonString []byte

	for key, val := range result {
		res, err := bson.MarshalExtJSON(val, true, true)
		if qilog.Check(err) {
			return "", err
    }
    if key != 0 {
    	jsonString = append(jsonString, byte(','))
    }
    jsonString = append(jsonString , res...)
	}
    
	var buf []byte
	buf = append(buf, "{\"Found result \" : ["...)
	buf = append(buf, jsonString...)
	buf = append(buf, "]}"...)
	response.Write(buf)
	
	var textString string
	err := json.Unmarshal(buf, &textString) 
	qilog.Error(err)
	return textString, nil
}
