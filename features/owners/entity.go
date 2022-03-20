package owners

import "time"

type Owners struct {
	ID            int
	Image         string
	TitleID       string
	DescriptionID string
	TitleEN       string
	DescriptionEN string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Bussiness interface {
	GetAllOwners(Owners) (own []Owners, err error)
	GetOwnersById(id int) (own Owners, err error)
	EditOwners(id int) (own Owners, err error)
}

type Data interface {
	SelectAllOwners(Owners) (own []Owners, err error)
	SelectOwnerById(id int) (own Owners, err error)
	UpdateOwners(id int) (own Owners, err error)
}
