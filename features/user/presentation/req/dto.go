package req

import (
	"java-agro/features/user"
	"time"
)

type User struct {
	ID        uint
	Name      string `json:"name" form:"name" gorm:"unique;not null"`
	Password  string `json:"password" form:"password" gorm:"not null"`
	Email     string `json:"email" form:"email" gorm:"unique;not null"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAuth struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func FromCore(core User) user.User {
	return user.User{
		ID:        int(core.ID),
		Name:      core.Name,
		Email:     core.Email,
		Password:  core.Password,
		Token:     core.Token,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func (core *User) ToUserCore() user.User {
	return user.User{
		ID:        int(core.ID),
		Name:      core.Name,
		Email:     core.Email,
		Password:  core.Password,
		Token:     core.Token,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func (userA *UserAuth) ToUserAuth() user.User {
	return user.User{
		Email:    userA.Email,
		Password: userA.Password,
	}
}
