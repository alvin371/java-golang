package data

import (
	"java-agro/features/user"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toUserCore(usr User) user.User {
	return user.User{
		ID:       int(usr.ID),
		Name:     usr.Name,
		Password: usr.Password,
		Email:    usr.Email,
		Token:    usr.Token,
	}
}

func toUserCoreList(accList []User) []user.User {
	convUser := []user.User{}

	for _, user := range accList {
		convUser = append(convUser, toUserCore(user))
	}
	return convUser
}

func fromCore(usr user.User) User {
	return User{
		Name:     usr.Name,
		Password: usr.Password,
		Email:    usr.Email,
		Token:    usr.Token,
	}
}
