package route

import (
	"Album-Backend/codegen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *api) GetFileInfo(ctx echo.Context, params codegen.GetFileInfoParams) error {
	return ctx.JSON(http.StatusOK, "GetFileInfo")
}
func (a *api) GetFiles(ctx echo.Context, params codegen.GetFilesParams) error {
	return ctx.JSON(http.StatusOK, "GetFiles")
}
func (a *api) GetHealthServices(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "pong")
}
