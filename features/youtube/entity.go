package youtube

import "time"

type Youtube struct {
	ID            int
	Link          string
	TitleID       string
	DescriptionID string
	TitleEN       string
	DescriptionEN string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Bussiness interface {
	GetAllYoutube(Youtube) (ytb []Youtube, err error)
	GetYoutubeById(id int) (ytb Youtube, err error)
	EditYoutube(id int) (ytb Youtube, err error)
}

type Data interface {
	SelectAllYoutube(Youtube) (ytb []Youtube, err error)
	SelectYoutubeById(id int) (ytb Youtube, err error)
	UpdateYoutube(id int) (ytb Youtube, err error)
}
