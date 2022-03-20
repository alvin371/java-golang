package data

import (
	"java-agro/features/highlights"
	"time"
)

type Highlights struct {
	ID            int
	TitleID       string
	DescriptionID string
	TitleEN       string
	DescriptionEN string
	Image         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func toHighlightsRecord(hg highlights.Highlights) Highlights {
	return Highlights{
		ID:            hg.ID,
		TitleID:       hg.TitleID,
		DescriptionID: hg.DescriptionID,
		TitleEN:       hg.TitleEN,
		DescriptionEN: hg.DescriptionEN,
		Image:         hg.Image,
	}
}

func toHighlightsCore(hg Highlights) highlights.Highlights {
	return highlights.Highlights{
		ID:            hg.ID,
		TitleID:       hg.TitleID,
		DescriptionID: hg.DescriptionID,
		TitleEN:       hg.TitleEN,
		DescriptionEN: hg.DescriptionEN,
		Image:         hg.Image,
	}
}

func toHighlightsCoreList(hg []Highlights) []highlights.Highlights {
	convHighlights := []highlights.Highlights{}

	for _, highighlightsList := range hg {
		convHighlights = append(convHighlights, toHighlightsCore(highighlightsList))
	}
	return convHighlights
}
