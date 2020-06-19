package v1

import (
	"github.com/gorilla/mux"
	"github.com/nodias/golang.templete.api/interface/rest/v1/protocol"
	"github.com/nodias/golang.templete.api/registry"
	"github.com/nodias/golang.templete.api/usecase"
)

func Apply(router *mux.Router, ctn *registry.Container){
	protocol.RegisterUserServiceRouter(router, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}