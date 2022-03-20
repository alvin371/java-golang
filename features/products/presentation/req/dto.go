package req

import "java-agro/features/products"

type Products struct {
	NameID        string `json:"name_id" form:"name_id"`
	DescriptionID string `json:"description_id" form:"description_id"`
	NameEN        string `json:"name_en" form:"name_en"`
	DescriptionEN string `json:"description_en" form:"description_en"`
	Category      string `json:"category" form:"category"`
	Image         string `json:"image" form:"image"`
}

func FromCore(core Products) products.Products {
	return products.Products{
		NameID:        core.NameID,
		DescriptionID: core.DescriptionID,
		NameEN:        core.NameEN,
		DescriptionEN: core.DescriptionEN,
		Category:      core.Category,
		Image:         core.Image,
	}
}

func (core *Products) ToProductsCore() products.Products {
	return products.Products{
		NameID:        core.NameID,
		DescriptionID: core.DescriptionID,
		NameEN:        core.NameEN,
		DescriptionEN: core.DescriptionEN,
		Category:      core.Category,
		Image:         core.Image,
	}
}
