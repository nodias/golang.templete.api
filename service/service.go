package service

import (
	"../database"
	"context"
	"github.com/nodias/golang.templete.common/logger"
	"github.com/nodias/golang.templete.common/model"
)

func GetUserInfo(ctx context.Context, id string) (*model.User, *model.ResponseError) {
	log := logger.New(ctx)

	db := database.NewOpenDB()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, model.NewResponseError(err, 500)
	}
	user := model.User{}
	defer db.Close()
	row := tx.QueryRowContext(ctx, "SELECT * FROM schema_user.user WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.WithError(err).Debug("There is no corresponding user information.")
		return nil, model.NewResponseError(err, 500)
	}
	log.WithField("user", user).Debug("User information retrieval was successful.")
	return &user, nil
}
