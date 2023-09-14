//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types,server,spec -package codegen api.yml > codegen/api.go"
package main

import (
	"Album-Backend/route"
	"fmt"
)

func main() {

	//logger.LogInit(config.AppInfo.LogPath, config.AppInfo.LogSaveName, config.AppInfo.LogFileExt)

	// service
	//services := service.NewServices(&repository)
	r := route.InitRouter()
	fmt.Println(r.Start("0.0.0.0:8081"))
	// listener, err := net.Listen("tcp", "0.0.0.0:8081")
	// if err != nil {
	// 	panic(err)
	// }
	// server := &http.Server{
	// 	Handler:           ,
	// 	ReadHeaderTimeout: 5 * time.Second,
	// }

	// err = server.Serve(listener)
	// fmt.Println(err)
}
