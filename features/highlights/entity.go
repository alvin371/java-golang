package highlights

import "time"

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

type Bussiness interface {
	GetAllHighlights(Highlights) (hg []Highlights, err error)
	GetHighlightsById(id int) (hg Highlights, err error)
	CreateHighlights(data Highlights) (err error)
	EditHighlights(id int, data Highlights) (hg Highlights, err error)
	DeleteHighlights(id int) (hg Highlights, err error)
}

type Data interface {
	SelectAllHighlights(Highlights) (hg []Highlights, err error)
	SelectHighlightsById(id int) (hg Highlights, err error)
	InsertHighlights(data Highlights) (err error)
	UpdateHighlights(id int, data Highlights) (hg Highlights, err error)
	DestroyHighlights(id int) (hg Highlights, err error)
}
