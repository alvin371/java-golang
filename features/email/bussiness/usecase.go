package bussiness

import "java-agro/features/email"

type EmailUsecase struct {
	eData email.Data
}

func NewBussinessEmail(emailData email.Data) email.Bussiness {
	return &EmailUsecase{
		eData: emailData,
	}
}

func (eu *EmailUsecase) GetAllEmail(data email.Email) (email []email.Email, err error) {
	news, err := eu.eData.SelectAllEmail(data)
	if err != nil {
		return nil, err
	}
	return news, nil
}
func (eu *EmailUsecase) CreateEmail(data email.Email) (err error) {
	if err := eu.eData.InsertEmail(data); err != nil {
		return err
	}
	return nil
}
