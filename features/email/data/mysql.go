package data

import (
	"java-agro/features/email"

	"gorm.io/gorm"
)

type mysqlEmailRepository struct {
	Conn *gorm.DB
}

func NewEmailRepository(conn *gorm.DB) email.Data {
	return &mysqlEmailRepository{
		Conn: conn,
	}
}

func (er *mysqlEmailRepository) SelectAllEmail(email.Email) (email []email.Email, err error) {
	var record []Email
	err = er.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toEmailCoreList(record), nil
}
func (er *mysqlEmailRepository) InsertEmail(data email.Email) (err error) {
	convData := toEmailRecord(data)

	if err = er.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
