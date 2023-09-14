package route

import (
	"Album-Backend/codegen"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

var (
	APIPath string
)

type api struct{}

func init() {
	swagger, err := codegen.GetSwagger()
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	APIPath = strings.TrimRight(u.Path, "/")
}

func NewAPIService() codegen.ServerInterface {
	return &api{}
}

func InitRouter() *echo.Echo {
	apiService := NewAPIService()

	e := echo.New()

	e.Use((echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderXCSRFToken, echo.HeaderContentType, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderConnection, echo.HeaderOrigin, echo.HeaderXRequestedWith},
		ExposeHeaders:    []string{echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders},
		MaxAge:           172800,
		AllowCredentials: true,
	})))

	e.Use(echo_middleware.Gzip())
	e.GET("/aaaa", func(c echo.Context) error {
		return c.JSON(200, "aaaa")
	})
	// e.Use(echo_middleware.RequestLoggerWithConfig(echo_middleware.RequestLoggerConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		return strings.Contains(c.Request().URL.Path, "status")
	// 	},
	// }))

	codegen.RegisterHandlersWithBaseURL(e, apiService, "/v1")

	return e
}
