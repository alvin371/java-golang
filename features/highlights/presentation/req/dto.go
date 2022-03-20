package req

import "java-agro/features/highlights"

type Highlights struct {
	TitleID       string `json:"title_id" form:"title_id"`
	DescriptionID string `json:"description_id" form:"description_id"`
	TitleEN       string `json:"title_en" form:"title_en"`
	DescriptionEN string `json:"description_en" form:"description_en"`
	Image         string `json:"image" form:"image"`
}

func FromCore(core Highlights) highlights.Highlights {
	return highlights.Highlights{
		TitleID:       core.TitleID,
		DescriptionID: core.DescriptionID,
		TitleEN:       core.TitleEN,
		DescriptionEN: core.DescriptionEN,
		Image:         core.Image,
	}
}

func (core *Highlights) ToHighlightsCore() highlights.Highlights {
	return highlights.Highlights{
		TitleID:       core.TitleID,
		DescriptionID: core.DescriptionID,
		TitleEN:       core.TitleEN,
		DescriptionEN: core.DescriptionEN,
		Image:         core.Image,
	}
}
