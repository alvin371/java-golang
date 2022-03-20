package email

import "time"

type Email struct {
	ID         int
	Email      string
	Created_At time.Time
	Updated_At time.Time
}

type Bussiness interface {
	// Create Data
	// Get All Data
	GetAllEmail(Email) (email []Email, err error)
	CreateEmail(data Email) (err error)
}

type Data interface {
	SelectAllEmail(Email) (email []Email, err error)
	InsertEmail(data Email) (err error)
}
