package rep

import "java-agro/features/user"

type User struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
	Token    string `json:"token"`
}

type UserLogin struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToUserCore(req user.User) User {
	return User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
	}
}

func ToUserCoreSlice(core []user.User) []User {
	var userArray []User
	for key := range core {
		userArray = append(userArray, ToUserCore(core[key]))
	}
	return userArray
}

func ToUserLoginResponse(user user.User) UserLogin {
	return UserLogin{
		ID:    uint(user.ID),
		Email: user.Email,
		Token: user.Token,
	}
}
