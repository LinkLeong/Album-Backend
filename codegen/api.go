// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// Files defines model for Files.
type Files struct {
	Size *int    `json:"size,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// GetImagesOK defines model for GetImagesOK.
type GetImagesOK = Files

// ResponseInternalServerError defines model for ResponseInternalServerError.
type ResponseInternalServerError = BaseResponse

// GetFileInfoParams defines parameters for GetFileInfo.
type GetFileInfoParams struct {
	// Path Folder path
	Path string `form:"path" json:"path"`
}

// GetFilesParams defines parameters for GetFiles.
type GetFilesParams struct {
	// Type 获取数据的类型
	Type *string `form:"type,omitempty" json:"type,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 获取二进制文件
	// (GET /file)
	GetFileInfo(ctx echo.Context, params GetFileInfoParams) error
	// 获取随机文件信息
	// (GET /files)
	GetFiles(ctx echo.Context, params GetFilesParams) error
	// Get service status
	// (GET /health)
	GetHealthServices(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetFileInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetFileInfo(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFileInfoParams
	// ------------- Required query parameter "path" -------------

	err = runtime.BindQueryParameter("form", true, true, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFileInfo(ctx, params)
	return err
}

// GetFiles converts echo context to params.
func (w *ServerInterfaceWrapper) GetFiles(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFilesParams
	// ------------- Optional query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, false, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFiles(ctx, params)
	return err
}

// GetHealthServices converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealthServices(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealthServices(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/file", wrapper.GetFileInfo)
	router.GET(baseURL+"/files", wrapper.GetFiles)
	router.GET(baseURL+"/health", wrapper.GetHealthServices)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVXW8bRRT9K6sLDyBtvVYqXvYtqDRYFTKCvkEeJrvX66l2Z7Yzs0lDZAlQ6aebGAkq",
	"lFSN0qrwAMUgpBLaCP8YvGvnKX8Bzew6ztqurSLy5vXMPefce87d3QKPRzFnyJQEdwsEypgzieZhBVUt",
	"IgHK+hX96HGmkCn9k8RxSD2iKGfONcmZ/k96TYyI/vW2wAa48JYzxnbyU+lcpiFKaLVaNvgoPUFjDQIu",
	"1K9Ay4ZPCv4aUygYCT9FsY7iAyG4+N8kvE8kjnhmKRlxWzm5lbPrewWCJiiBuFsQCx6jUDSfXIRSksAc",
	"lLGLA0ugSgRD31rbtGROI6mPFm1YqokCLSotwjbBBrxBojhEcAFsEEj8Ogs3wVUiQRvUZqxPpBKUBXkv",
	"+YCnFEn6hZFzCrdUrZ7WU6YwQKENSERYugZNpWLXcTY2NiprhPpJxeMRzGDW3JQ1+HTPg73D9FY77WwP",
	"fu5m28/Snaf/fPn150yDUGU4lsO1JLKWP66BDesoZF63vqQF8RgZiSm4cLFSrVwEG2KimqYnp0FD01SA",
	"app2uP1nuvMwe3i7/+rFYPdm/2V72NtL77zIvv8te/ArGGRh8lPzwdVh16Or6RY0hyARKhQS3M8mkS/z",
	"0EdhaR0lg5xLy1eXHVK5FgeghwEuXE9QaBcZifSNokTg9YQK9Ec2jnM7OdZVu7ySS9XqnD3gnkJ1QSqB",
	"JCrvQ4OLiChwYY0yYhTNiE65y+xOJ723P+x9l+49Tvf+Hty9re14Lxcwa79OhTrzttjsURJFWsXIpLE3",
	"xi2tjgR68kBMMiJUTe5LWNXFxnW5wPbj3Z3s0cscrt87yL7qnhy10/Y32f1f8l5OjtrDn24dP/n25Kh9",
	"vP9H/6/O4Pnd14VCLkpEkTWTrMHuzcHvr9LH90vZEIT5ZnFmxcJ48R9iMN+Fs6/v83Jues4LzGsiCVXz",
	"te5l+8+G3SfZowfpvYPsh27a+TF7/jQ9PBz2OsOD9ix/PjSAWif1jFFvsC/T34032opzmOoKKvM9oB5a",
	"UhGVyDPzzFs9O1Bda7DyUJqXNzig81IUTQ74nav1S/V3x9GbwGzZiwrKpmqmGxcUCVYET+KcsLj4UXFl",
	"jv7V1r8BAAD//0OC5geACAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}