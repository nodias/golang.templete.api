package repository

import "github.com/nodias/golang.templete.api/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}
