package user

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	GetAllUser(User) (user []User, err error)
	GetUserById(id int) (user User, err error)
	CreateUser(data User) (err error)
	EditUser(id int) (usr User, err error)
	LoginUser(user User) (usr User, err error)
}
type Data interface {
	SelectAllUser(User) (user []User, err error)
	SelectUserById(id int) (user User, err error)
	InsertUser(data User) (err error)
	UpdateUser(id int) (usr User, err error)
	CheckAccount(User) (user User, err error)
}
