package user_repository

import (
	"context"

	"github.com/VanessaPellegrini/CMS_headless/app/infrastructure/mongodb"
	m "github.com/VanessaPellegrini/CMS_headless/app/models"
)

var collection = mongodb.GetCollection("users")
var ctx = context.Background()

func Create(user m.User) error {
	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {

	var users m.Users

	return users, nil
}

func Update(user m.User, userID string) error {
	return nil
}

func Delete(userId string) error {
	return nil
}
