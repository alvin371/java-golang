package bussiness

import (
	"java-agro/features/highlights"
)

type HighlightsUsecase struct {
	hData highlights.Data
}

func NewBussinessHighlights(hiData highlights.Data) highlights.Bussiness {
	return &HighlightsUsecase{
		hData: hiData,
	}
}

func (hu *HighlightsUsecase) GetAllHighlights(data highlights.Highlights) (hg []highlights.Highlights, err error) {
	highlights, err := hu.hData.SelectAllHighlights(data)
	if err != nil {
		return nil, err
	}
	return highlights, nil
}
func (hu *HighlightsUsecase) GetHighlightsById(id int) (hg highlights.Highlights, err error) {
	highlightsData, err := hu.hData.SelectHighlightsById(id)

	if err != nil {
		return highlights.Highlights{}, err
	}
	return highlightsData, nil
}

func (hu *HighlightsUsecase) CreateHighlights(data highlights.Highlights) (err error) {
	if err := hu.hData.InsertHighlights(data); err != nil {
		return err
	}
	return nil
}

func (hu *HighlightsUsecase) EditHighlights(id int, hg highlights.Highlights) (result highlights.Highlights, err error) {
	highlightsData, err := hu.hData.UpdateHighlights(id, hg)

	if err != nil {
		return result, err
	}
	return highlightsData, nil
}

func (hu *HighlightsUsecase) DeleteHighlights(id int) (hg highlights.Highlights, err error) {
	highlightsData, err := hu.hData.DestroyHighlights(id)

	if err != nil {
		return hg, err
	}
	return highlightsData, nil
}
