package web

import (
	"net/http"

	"kahrersoftware.at/webskeleton/config"

	"github.com/labstack/echo"
)

//EchoController front controller using echo framework
type EchoController struct {
}

//NewEchoController ...
func NewEchoController(env *config.Env) *EchoController {
	return &EchoController{}
}

//InitWeb initialize the web framework
func (c *EchoController) InitWeb(echo *echo.Echo) {
	echo.Renderer = &TemplateRenderer{}

	echo.GET("/health", c.HandleHealth())

	echo.GET("/:name", c.HandlePage())

	echo.POST("/loginuser", c.HandleLogin())

	echo.GET("/logout", c.HandleLogout())

	echo.Static("/assets", "assets")
	// r.Use(c.loggingMiddleware)
	// r.Use(authMiddleware)
}

//HandlePage ...
func (c *EchoController) HandlePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		pageName := c.Param("name")
		c.Render(http.StatusOK, pageName+".html", NewViewData(c.Request()))
		return nil
	}
}

//HandleLogin login user
func (c *EchoController) HandleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		//do login here
		return nil
	}

}

//HandleLogout logout user
func (c *EchoController) HandleLogout() echo.HandlerFunc {
	return func(c echo.Context) error {
		//do logout here
		return nil
	}

}

//HandleHealth ...
func (c *EchoController) HandleHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		vd := NewViewData(c.Request())
		vd["status"] = "ok"
		c.JSON(http.StatusOK, vd)

		return nil
	}

}
