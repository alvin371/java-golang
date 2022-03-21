package route

import (
	"java-agro/config"
	"java-agro/factory"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.CORS())
	jwt := e.Group("")

	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// User Credential
	e.POST("/user/register", presenter.UserPresentation.CreateUserHandler)
	e.POST("/user/login", presenter.UserPresentation.LoginUserHandler)
	jwt.GET("/user", presenter.UserPresentation.GetAllUserHandler)
	jwt.GET("/user/:id", presenter.UserPresentation.GetUserById)
	e.PUT("/user/:id", presenter.UserPresentation.UpdateAccountHandler)

	// product
	e.GET("/products", presenter.ProductPresentation.GetAllProducts)
	e.GET("/products/:id", presenter.ProductPresentation.GetProductsById)
	jwt.POST("/products/create", presenter.ProductPresentation.CreateProducts)
	jwt.PATCH("/products/edit/:id", presenter.ProductPresentation.UpdateProducts)
	jwt.PATCH("/products/delete/:id", presenter.ProductPresentation.DeleteProducts)

	// email
	e.POST("/email/create", presenter.EmailPresentation.CreateEmail)
	jwt.GET("/email", presenter.EmailPresentation.GetAllEmailHandler)

	// highlights
	e.GET("/highlights", presenter.HighlightsPresentation.GetAllHighlights)
	e.GET("/highlights/:id", presenter.HighlightsPresentation.GetHighlightsById)
	jwt.POST("/highlights/create", presenter.HighlightsPresentation.CreateHighlights)
	jwt.PATCH("/highlights/edit/:id", presenter.HighlightsPresentation.UpdateHighlights)
	jwt.DELETE("/highlights/delete/:id", presenter.HighlightsPresentation.DeleteHighlights)
	return e
}
