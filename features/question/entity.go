package question

import "time"

type Question struct {
	ID        int
	Name      string
	Email     string
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	GetAllQuestion(Question) (quest []Question, err error)
	CreateQuestion(data Question) (err error)
}

type Data interface {
	SelectAllQuestion(Question) (quest []Question, err error)
	InsertQuestion(data Question) (err error)
}
