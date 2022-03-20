package req

import "java-agro/features/email"

type Email struct {
	Email string `json: "email"`
}

func ToCore(core Email) email.Email {
	return email.Email{
		Email: core.Email,
	}
}

func (core *Email) ToEmailCore() email.Email {
	return email.Email{
		Email: core.Email,
	}
}
