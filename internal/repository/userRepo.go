package repository

import "github.com/Nikik0/dataCollectorBot/internal/model"

func SaveUser(u *model.User) bool {
	return true
}

func FindByTgId(id string) model.User {
	return model.User{}
}

func FindByIdOrCreateNew(id int64) model.User {
	return *model.NewUser(id)
}
