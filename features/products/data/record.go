package data

import (
	"java-agro/features/products"
	"time"
)

type Products struct {
	ID            int
	NameID        string
	DescriptionID string
	NameEN        string
	DescriptionEN string
	Category      string
	Image         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func toProductsRecord(prd products.Products) Products {
	return Products{
		ID:            prd.ID,
		NameID:        prd.NameID,
		DescriptionID: prd.DescriptionID,
		NameEN:        prd.NameEN,
		DescriptionEN: prd.DescriptionEN,
		Category:      prd.Category,
		Image:         prd.Image,
	}
}

func toProductsCore(prd Products) products.Products {
	return products.Products{
		ID:            prd.ID,
		NameID:        prd.NameID,
		DescriptionID: prd.DescriptionID,
		NameEN:        prd.NameEN,
		DescriptionEN: prd.DescriptionEN,
		Category:      prd.Category,
		Image:         prd.Image,
	}
}

func toProductsCoreList(prd []Products) []products.Products {
	convProducts := []products.Products{}

	for _, productList := range prd {
		convProducts = append(convProducts, toProductsCore(productList))
	}
	return convProducts
}
