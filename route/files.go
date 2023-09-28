package route

import (
	"Album-Backend/codegen"
	"Album-Backend/config"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *api) GetFileInfo(ctx echo.Context, params codegen.GetFileInfoParams) error {
	filePath := params.Path
	fileName := path.Base(filePath)
	ctx.Response().Writer.Header().Add("Content-Disposition", "attachment; filename*=utf-8''"+url.PathEscape(fileName))
	http.ServeFile(ctx.Response().Writer, ctx.Request(), filePath)
	return ctx.JSON(http.StatusOK, "GetFileInfo")
}
func (a *api) GetFiles(ctx echo.Context, params codegen.GetFilesParams) error {
	host := ctx.Request().Host
	dirs := strings.Split(config.Dirs, ",")

	// Slice to hold all file paths
	var files []string

	// Read files from each directory
	for _, dir := range dirs {
		fileInfos, err := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				continue
			}
			files = append(files, filepath.Join(dir, fileInfo.Name()))
		}
	}
	// 创建随机数源
	src := rand.NewSource(time.Now().UnixNano())
	f := []string{}
	// Print 10 random files
	for i := 0; i < 10; i++ {
		// 基于随机数源创建随机数生成器
		r := rand.New(src)
		// 生成随机数
		ro := r.Intn(len(files))
		f = append(f, "http://"+host+"/v1/file?path="+files[ro])
	}
	return ctx.JSON(http.StatusOK, f)
}
func (a *api) GetHealthServices(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "pong")
}
