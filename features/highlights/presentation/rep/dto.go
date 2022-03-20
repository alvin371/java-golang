package rep

import "java-agro/features/highlights"

type Highlights struct {
	ID            int    `json:"id"`
	TitleID       string `json:"title_id"`
	DescriptionID string `json:"description_id"`
	TitleEN       string `json:"title_en"`
	DescriptionEN string `json:"description_en"`
	Image         string `json:"image"`
}

func ToCore(req highlights.Highlights) Highlights {
	return Highlights{
		ID:            req.ID,
		TitleID:       req.TitleID,
		DescriptionID: req.DescriptionID,
		TitleEN:       req.TitleEN,
		DescriptionEN: req.DescriptionEN,
		Image:         req.Image,
	}
}

func ToCoreSlice(core []highlights.Highlights) []Highlights {
	var highlightsArray []Highlights
	for key := range core {
		highlightsArray = append(highlightsArray, ToCore(core[key]))
	}
	return highlightsArray
}
