package webecho

import (
	"github.com/labstack/echo"
	"kahrersoftware.at/webskeleton/config"
	"kahrersoftware.at/webskeleton/web"
)

//InitWeb initialize the web framework
func InitWeb(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// e.Renderer = &TemplateRenderer{}

	e.GET("/health", echo.WrapHandler(web.HandleHealth()))

	e.POST("/loginuser", echo.WrapHandler(web.HandleLogin(env)))

	e.GET("/logout", echo.WrapHandler(web.HandleLogout()))

	e.Static(web.AssetPattern, web.AssetRoot)

	e.Match([]string{"GET", "POST"}, "/:page", DefaultPageHandler())

	e.Use(echo.WrapMiddleware(web.LoggingMiddleware))
	e.Use(echo.WrapMiddleware(web.AuthMiddleware))

	return e
}

//DefaultPageHandler ...
func DefaultPageHandler() echo.HandlerFunc {
	return func(ec echo.Context) error {
		page := ec.Param("page")
		web.RenderTemplate(ec.Response().Writer, ec.Request(), page+".html", nil)
		// ec.Render(http.StatusOK, pageName+".html", web.NewTemplateData(EchoContextData(ec)))
		return nil
	}
}

// TemplateRenderer is a custom html/template renderer for Echo framework
// damit man echo.Context.Render aufrufen kann
// type TemplateRenderer struct {
// }

// // Render renders a template document
// func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	log.Println("render", name)
// 	th := web.NewTemplateHandler(name)
// 	return th.Templ.Execute(w, data)
// }
