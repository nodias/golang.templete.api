package postgres

import (
	"database/sql"
	"github.com/nodias/golang.templete.api/domain/model"
	"github.com/nodias/golang.templete.common/shared/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() *userRepository {
	db := repository.NewOpenDB()
	return &userRepository{db: db}
}

func (u userRepository) FindAll() (users []*model.User, err error) {
	rows, err := u.db.Query("SELECT id, email FROM public.users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id string
	var email string

	for rows.Next() {
		err := rows.Scan(&id, &email)
		if err != nil {
			return nil, err
		}
		user := model.User{
			ID:    id,
			Email: email,
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u userRepository) FindByEmail(email string) (*model.User, error) {
	panic("implement me")
}

func (u userRepository) Save(*model.User) error {
	panic("implement me")
}

type User struct {
	ID    string
	Email string
}
