package data

import (
	"java-agro/features/email"
	"time"
)

type Email struct {
	ID        int
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toEmailRecord(email email.Email) Email {
	return Email{
		Email: email.Email,
	}
}

func toEmailCore(eml Email) email.Email {
	return email.Email{
		ID:    int(eml.ID),
		Email: eml.Email,
	}
}

func toEmailCoreList(eml []Email) []email.Email {
	convEmail := []email.Email{}

	for _, emailList := range eml {
		convEmail = append(convEmail, toEmailCore(emailList))
	}
	return convEmail
}
