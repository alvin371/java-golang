package factory

import (
	"java-agro/driver"
	_email_bussiness "java-agro/features/email/bussiness"
	_email_data "java-agro/features/email/data"
	_email_presentation "java-agro/features/email/presentation"

	_product_bussiness "java-agro/features/products/bussiness"
	_product_data "java-agro/features/products/data"
	_product_presentation "java-agro/features/products/presentation"

	_user_bussiness "java-agro/features/user/bussiness"
	_user_data "java-agro/features/user/data"
	_user_presentation "java-agro/features/user/presentation"

	_highlights_bussiness "java-agro/features/highlights/bussiness"
	_highlights_data "java-agro/features/highlights/data"
	_highlights_presentation "java-agro/features/highlights/presentation"
)

type Presenter struct {
	UserPresentation       *_user_presentation.UserHandler
	ProductPresentation    *_product_presentation.ProductsHandler
	EmailPresentation      *_email_presentation.EmailHandler
	HighlightsPresentation *_highlights_presentation.HighlightsHandler
}

func Init() Presenter {

	// user
	userData := _user_data.NewMySqlUSer(driver.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewHandlerAccount(userBussiness)

	productData := _product_data.NewProductRepository(driver.DB)
	productBussiness := _product_bussiness.NewBussinessProduct(productData)
	productPresentation := _product_presentation.NewProductsHandler(productBussiness)

	emailData := _email_data.NewEmailRepository(driver.DB)
	emailBussiness := _email_bussiness.NewBussinessEmail(emailData)
	emailPresentation := _email_presentation.NewEmailHandler(emailBussiness)

	highlightsData := _highlights_data.NewHighlightsRepository(driver.DB)
	highlightsBussiness := _highlights_bussiness.NewBussinessHighlights(highlightsData)
	highlightsPresentation := _highlights_presentation.NewHighlightsHandler(highlightsBussiness)
	return Presenter{
		UserPresentation:       userPresentation,
		ProductPresentation:    productPresentation,
		EmailPresentation:      emailPresentation,
		HighlightsPresentation: highlightsPresentation,
	}

}
