package faq

import "time"

type Faq struct {
	ID         int
	QuestionID string
	AnswerID   string
	QuestionEN string
	AnswerEN   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Bussiness interface {
	// Create Data
	// Get All Data
	// Delete Data
	// Edit Data
	GetAllFaq(Faq) (faq []Faq, err error)
	CreateFaq(data Faq) (err error)
	EditFaq(id int) (faq Faq, err error)
	DeleteFaq(id int) (faq Faq, err error)
}

type Data interface {
	SelectAllFaq(Faq) (faq []Faq, err error)
	InsertFaq(data Faq) (err error)
	UpdateFaq(id int) (faq Faq, err error)
	DestroyFaq(Faq) (faq Faq, err error)
}
