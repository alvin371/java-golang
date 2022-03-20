package data

import (
	"errors"
	"fmt"
	"java-agro/features/highlights"

	"gorm.io/gorm"
)

type mySqlHighlightsRepository struct {
	Conn *gorm.DB
}

func NewHighlightsRepository(conn *gorm.DB) highlights.Data {
	return &mySqlHighlightsRepository{
		Conn: conn,
	}
}

func (hr *mySqlHighlightsRepository) SelectAllHighlights(highlights.Highlights) (hg []highlights.Highlights, err error) {
	var record []Highlights
	err = hr.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toHighlightsCoreList(record), nil
}
func (hr *mySqlHighlightsRepository) SelectHighlightsById(id int) (hg highlights.Highlights, err error) {
	var idHighlights Highlights

	err = hr.Conn.First(&idHighlights, id).Error

	if idHighlights.TitleID == "" && idHighlights.ID == 0 {
		return highlights.Highlights{}, errors.New("Highlights not found")
	}

	if err != nil {
		return highlights.Highlights{}, err
	}

	return toHighlightsCore(idHighlights), nil
}
func (hr *mySqlHighlightsRepository) InsertHighlights(data highlights.Highlights) (err error) {
	convData := toHighlightsRecord(data)

	if err := hr.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
func (hr *mySqlHighlightsRepository) UpdateHighlights(id int, hg highlights.Highlights) (result highlights.Highlights, err error) {
	var singleHighlights Highlights
	err = hr.Conn.Model(&singleHighlights).Where("id=?", id).Updates(&hg).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
func (hr *mySqlHighlightsRepository) DestroyHighlights(id int) (hg highlights.Highlights, err error) {
	var singleHighlights Highlights
	fmt.Println("Isi single highlights : ", singleHighlights)
	fmt.Println("id : ", id)
	err = hr.Conn.Model(&singleHighlights).Where("id=?", id).Delete(&singleHighlights).Error

	if err != nil {
		return hg, err
	}

	return toHighlightsCore(singleHighlights), nil
}
