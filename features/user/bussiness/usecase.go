package bussiness

import (
	"fmt"
	"java-agro/features/user"
	middlewares "java-agro/middleware"

	"github.com/go-playground/validator"
)

type UserBussiness struct {
	userData user.Data
	validate *validator.Validate
}

func NewUserBussiness(userData user.Data) user.Bussiness {
	return &UserBussiness{
		userData: userData,
	}
}

func (ub *UserBussiness) GetAllUser(userData user.User) (resp []user.User, err error) {
	user, err := ub.userData.SelectAllUser(userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (ub *UserBussiness) GetUserById(id int) (user user.User, err error) {
	userData, err := ub.userData.SelectUserById(id)
	if err != nil {
		return user, err
	}
	return userData, nil
}
func (ub *UserBussiness) CreateUser(data user.User) (err error) {
	if err := ub.userData.InsertUser(data); err != nil {
		return err
	}
	return nil
}
func (ub *UserBussiness) EditUser(id int) (usr user.User, err error) {
	userData, err := ub.userData.UpdateUser(id)

	if err != nil {
		return user.User{}, err
	}
	return userData, nil
}
func (ub *UserBussiness) LoginUser(user user.User) (usr user.User, err error) {
	accountData, err := ub.userData.CheckAccount(user)
	fmt.Println("ini user bussiness", user)
	if err != nil {
		return user, err
	}
	accountData.Token, err = middlewares.CreateToken(uint(user.ID), user.Name)
	if err != nil {
		return user, err
	}
	return accountData, nil
}
