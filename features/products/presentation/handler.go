package presentation

import (
	"fmt"
	"java-agro/features/products"
	presentation_response "java-agro/features/products/presentation/rep"
	presentation_request "java-agro/features/products/presentation/req"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductsHandler struct {
	productBussiness products.Bussiness
}

func NewProductsHandler(productBussiness products.Bussiness) *ProductsHandler {
	return &ProductsHandler{
		productBussiness: productBussiness,
	}
}

func (ph *ProductsHandler) GetAllProducts(e echo.Context) error {
	data, err := ph.productBussiness.GetAllProducts(products.Products{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCoreSlice(data),
	})
}
func (ph *ProductsHandler) GetProductsById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := ph.productBussiness.GetProductsById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCore(data),
	})
}
func (ph *ProductsHandler) CreateProducts(c echo.Context) error {
	newProducts := presentation_request.Products{}
	if err := c.Bind(&newProducts); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("ini adalah create products", newProducts)
	if err := ph.productBussiness.CreateProducts(newProducts.ToProductsCore()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newProducts,
	})
}
func (ph *ProductsHandler) UpdateProducts(e echo.Context) error {
	editProducts := presentation_request.Products{}
	e.Bind(&editProducts)
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := ph.productBussiness.EditProducts(id, editProducts.ToProductsCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCore(data),
	})
}
func (ph *ProductsHandler) DeleteProducts(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := ph.productBussiness.DeleteProducts(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCore(data),
	})
}
