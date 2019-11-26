package endPoint

import (	
	"mongogolanglib/src/router"
	"github.com/gorilla/mux"
	"net/http"
)

var ServerCloseCount = 0
var Server [3] http.Server
type HandlerType struct{Router *mux.Router}
var Handlers [3] HandlerType

func (Handler HandlerType) Route(){
	
	Handler.Router.HandleFunc("/connecttoserver", router.EstablishConn).Methods("Get")
	Handler.Router.HandleFunc("/connecttodb", router.CreateDBSelectEndPoint).Methods("Get")

	Handler.Router.HandleFunc("/insertone", router.CreateInsertOneEndpoint).Methods("Post")
	Handler.Router.HandleFunc("/insertmany", router.CreateInsertManyEndpoint).Methods("Post")

	Handler.Router.HandleFunc("/findone", router.CreateFindOneEndpoint).Methods("Get")
	Handler.Router.HandleFunc("/findall", router.CreateFindAllEndpoint).Methods("Get")

	Handler.Router.HandleFunc("/updateone", router.CreateUpdateOneEndpoint).Methods("Put")
	Handler.Router.HandleFunc("/updatemany", router.CreateUpdateManyEndpoint).Methods("Put")

	Handler.Router.HandleFunc("/deleteone", router.CreateDeleteOneEndpoint).Methods("Delete")
	Handler.Router.HandleFunc("/deletemany", router.CreateDeleteManyEndpoint).Methods("Delete")
	Handler.Router.HandleFunc("/deleteall", router.CreateDeleteAllEndpoint).Methods("Delete")

}

func init() {

	Handlers[0].Router = mux.NewRouter()
	Handlers[1].Router = mux.NewRouter()
	Handlers[2].Router = mux.NewRouter()

	Server[0] = http.Server{Addr : "localhost:12345",  Handler : Handlers[0].Router}
	Server[1] = http.Server{Addr : "localhost:12344",  Handler : Handlers[1].Router}
	Server[2] = http.Server{Addr : "localhost:54333",  Handler : Handlers[2].Router}
}
