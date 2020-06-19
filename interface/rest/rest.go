package rest

import (
	"github.com/gorilla/mux"
	v1 "github.com/nodias/golang.templete.api/interface/rest/v1"
	"github.com/nodias/golang.templete.api/registry"
)

func Apply(router *mux.Router, ctn *registry.Container){
	v1.Apply(router.PathPrefix("/user/api").Subrouter(), ctn)
}
