package api

import (
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web/api/env"
	apihandler "github.com/devgek/webskeleton/web/api/handler"
	"github.com/devgek/webskeleton/web/common"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

//InitEchoApi initialize the echo web framework for serving an api
func InitEchoApi(apiEnv *env.ApiEnv) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	if config.IsServerDebug() {
		e.Debug = true
		// e.Use(middleware.Recover())
	}

	e.HTTPErrorHandler = common.LoggingDefaultHTTPErrorHandler

	// api
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		ContextKey: "token",
		Skipper:    common.JWTAuthSkipper,
	}))
	apiGroup.Use(common.TokenLoggingMiddleware)

	apiGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead},
	}))

	// OPTIONS because of CORS Preflight requests sent from axios
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/login", apihandler.HandleAPILogin)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitylist:entity", apihandler.HandleAPIEntityList)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/optionlist:entity", apihandler.HandleAPIOptionList)

	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitynew:entity", apihandler.HandleAPICreateEntity)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entityedit:entity", apihandler.HandleAPIUpdateEntity)
	apiGroup.Match([]string{"OPTIONS", "POST"}, "/entitydelete:entity/:id", apihandler.HandleAPIDeleteEntity)

	apiGroup.PUT("/allnew:entity", apihandler.HandleAPICreateAll)
	apiGroup.GET("/health", apihandler.HandleAPIHealth)

	e.GET("/favicon.ico", common.HandleFavicon)

	e.Use(apihandler.ApiEnvContextMiddleware)
	e.Use(common.RequestLoggingMiddleware)

	return e
}
