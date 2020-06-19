package usecase

import (
	"github.com/google/uuid"
	"github.com/nodias/golang.templete.api/domain/model"
	"github.com/nodias/golang.templete.api/domain/repository"
	"github.com/nodias/golang.templete.api/domain/service"
	"github.com/nodias/golang.templete.common/models"
	"net/http"
)

type UserUsecase interface {
	ListUser() ([]*User, *models.ResponseError)
	RegisterUser(email string) *models.ResponseError
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{repo: repo, service: service}
}

func (u *userUsecase) ListUser() ([]*User, *models.ResponseError) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, models.NewResponseError(err, http.StatusInternalServerError)
	}
	return toUser(users), nil
}

func (u *userUsecase) RegisterUser(email string) *models.ResponseError {
	uid, err := uuid.NewRandom()
	if err != nil {
		models.NewResponseError(err, http.StatusInternalServerError)
	}
	if err := u.service.Duplicated(email); err != nil {
		return models.NewResponseError(err, http.StatusInternalServerError)
	}
	user := model.NewUser(uid.String(), email)
	if err := u.repo.Save(user); err != nil {
		return models.NewResponseError(err, http.StatusInternalServerError)
	}
	return nil
}

type User struct {
	ID    string
	Email string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:    user.GetID(),
			Email: user.GetEmail(),
		}
	}
	return res
}
