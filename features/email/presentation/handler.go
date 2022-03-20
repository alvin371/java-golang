package presentation

import (
	"fmt"
	"java-agro/features/email"
	presentation_response "java-agro/features/email/presentation/rep"
	presentation_request "java-agro/features/email/presentation/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmailHandler struct {
	emailBussiness email.Bussiness
}

func NewEmailHandler(emailBussiness email.Bussiness) *EmailHandler {
	return &EmailHandler{emailBussiness: emailBussiness}
}

func (eh *EmailHandler) GetAllEmailHandler(c echo.Context) error {
	data, err := eh.emailBussiness.GetAllEmail(email.Email{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCoreSlice(data),
	})
}

func (eh *EmailHandler) CreateEmail(c echo.Context) error {
	newNews := presentation_request.Email{}
	fmt.Println("testing", c)

	if err := c.Bind(&newNews); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := eh.emailBussiness.CreateEmail(newNews.ToEmailCore()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newNews,
	})

}
