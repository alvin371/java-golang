package rep

import "java-agro/features/products"

type Products struct {
	ID            int    `json:"id"`
	NameID        string `json:"name_id"`
	DescriptionID string `json:"description_id"`
	NameEN        string `json:"name_en"`
	DescriptionEN string `json:"description_en"`
	Category      string `json:"category"`
	Image         string `json:"image"`
}

func ToCore(req products.Products) Products {
	return Products{
		ID:            req.ID,
		NameID:        req.NameID,
		DescriptionID: req.DescriptionID,
		NameEN:        req.NameEN,
		DescriptionEN: req.DescriptionEN,
		Category:      req.Category,
		Image:         req.Image,
	}
}

func ToCoreSlice(core []products.Products) []Products {
	var productsArray []Products
	for key := range core {
		productsArray = append(productsArray, ToCore(core[key]))
	}
	return productsArray
}
