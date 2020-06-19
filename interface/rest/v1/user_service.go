package v1

import (
	"encoding/json"
	"github.com/nodias/golang.templete.api/usecase"
	"github.com/nodias/golang.templete.common/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

func (s userService) ListUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) () {
	users, rErr := s.userUsecase.ListUser()
	if rErr != nil {
		logrus.WithError(rErr).Error("Failed to load user")
		w.WriteHeader(rErr.Code)
	}
	err := json.NewEncoder(w).Encode(models.Response{
		Id: models.ID("123"),
		Body: users,
		Error: rErr,
	})
	if err != nil {
		logrus.WithError(err).Error("Failed encode to json")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s userService) RegisterUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)() {
	panic("implement me")
}

func NewUserService(userUsecase usecase.UserUsecase) *userService{
	return &userService{userUsecase: userUsecase}
}

