package protocol

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

type UserServiceServer interface {
	ListUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) ()
	RegisterUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) ()
}

func RegisterUserServiceRouter(router *mux.Router, server UserServiceServer){
	router.Handle("/users", negroni.New(negroni.HandlerFunc(server.ListUser))).Methods("GET")
	router.Handle("/users", negroni.New(negroni.HandlerFunc(server.RegisterUser))).Methods("POST")
}
