package rep

import "java-agro/features/email"

type Email struct {
	Email string `json: "email"`
}

func ToCore(req email.Email) Email {
	return Email{
		Email: req.Email,
	}
}

func ToCoreSlice(core []email.Email) []Email {
	var emailArray []Email
	for key := range core {
		emailArray = append(emailArray, ToCore(core[key]))
	}
	return emailArray
}
