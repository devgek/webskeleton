package webecho

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/handler"
	"github.com/labstack/echo"
)

//InitWeb initialize the web framework
func InitWeb(env *config.Env) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// e.Renderer = &TemplateRenderer{}

	e.GET("/health", echo.WrapHandler(handler.HandleHealth()))

	e.POST("/loginuser", echo.WrapHandler(handler.HandleLogin(env)))

	e.GET("/logout", echo.WrapHandler(handler.HandleLogout()))

	e.GET("/favicon.ico", echo.WrapHandler(handler.HandleFavicon()))

	e.Static(web.AssetPattern, web.AssetRoot)

	e.Match([]string{"GET", "POST"}, "/users", echo.WrapHandler(handler.HandleUsers(env)))
	e.POST("/useredit", echo.WrapHandler(handler.HandleUserEdit(env)))
	e.POST("/usernew", echo.WrapHandler(handler.HandleUserNew(env)))
	e.POST("/userdelete", echo.WrapHandler(handler.HandleUserDelete(env)))

	e.Match([]string{"GET", "POST"}, "/:page", DefaultPageHandler())

	e.Use(echo.WrapMiddleware(handler.RequestLoggingMiddleware))
	e.Use(echo.WrapMiddleware(handler.AuthMiddleware))

	return e
}

//DefaultPageHandler ...
func DefaultPageHandler() echo.HandlerFunc {
	return func(ec echo.Context) error {
		page := ec.Param("page")
		web.RenderTemplate(ec.Response().Writer, ec.Request(), page+".html", nil)
		// ec.Render(http.StatusOK, pageName+".html", handler.NewTemplateData(EchoContextData(ec)))
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
// 	th := handler.NewTemplateHandler(name)
// 	return th.Templ.Execute(w, data)
// }
