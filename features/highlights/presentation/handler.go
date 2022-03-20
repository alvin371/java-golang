package presentation

import (
	"fmt"
	"java-agro/features/highlights"
	presentation_response "java-agro/features/highlights/presentation/rep"
	presentation_request "java-agro/features/highlights/presentation/req"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HighlightsHandler struct {
	highlightsBussiness highlights.Bussiness
}

func NewHighlightsHandler(highlightsBussiness highlights.Bussiness) *HighlightsHandler {
	return &HighlightsHandler{
		highlightsBussiness: highlightsBussiness,
	}
}

func (hh *HighlightsHandler) GetAllHighlights(e echo.Context) error {
	data, err := hh.highlightsBussiness.GetAllHighlights(highlights.Highlights{})

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
func (hh *HighlightsHandler) GetHighlightsById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := hh.highlightsBussiness.GetHighlightsById(id)
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
func (hh *HighlightsHandler) CreateHighlights(c echo.Context) error {
	newProducts := presentation_request.Highlights{}
	fmt.Println("testing", c)

	if err := c.Bind(&newProducts); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := hh.highlightsBussiness.CreateHighlights(newProducts.ToHighlightsCore()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newProducts,
	})
}
func (hh *HighlightsHandler) UpdateHighlights(e echo.Context) error {
	editHighlights := presentation_request.Highlights{}
	e.Bind(&editHighlights)
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := hh.highlightsBussiness.EditHighlights(id, editHighlights.ToHighlightsCore())

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
func (hh *HighlightsHandler) DeleteHighlights(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := hh.highlightsBussiness.DeleteHighlights(id)
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
