package main

import (
	"endPoint"
	"qilog"
)

func connectToServer(index int){
	qilog.Info("Server:", index, "runnning at", endPoint.Server[index].Addr)
	qilog.Debugf("Server Details: %#v\n", endPoint.Server[index])
	endPoint.Handlers[index].Route()

	err := endPoint.Server[index].ListenAndServe()
	if qilog.Check(err){
		qilog.Error("Error while listening at port:" + endPoint.Server[index].Addr, err)
	}

	defer func(){
		endPoint.ServerCloseCount++
		if qilog.Check(endPoint.Server[index].Close()){
		qilog.Error("Error while closing server at:", endPoint.Server[index].Addr)
		}
	} ()
}

func main() {
	qilog.Info("Starting application...")
	for index, _ := range endPoint.Server{
		go connectToServer(index)
	}
	for {
		if endPoint.ServerCloseCount == 3 {
			qilog.Info("All servers closed.")
			break;
		}
	}
}
