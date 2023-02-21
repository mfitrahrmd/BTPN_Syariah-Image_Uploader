package main

import (
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/server"
)

func main() {
	srv := server.BuildServer()

	// swagger API documentation
	srv.RouterEngine.StaticFile("/swagger.json", "openapi/spec/swagger.json")
	srv.RouterEngine.Static("/swagger/", "openapi/dist/")
	srv.Logger.Printf("http://%s/swagger for API documentation", srv.GetAddress())

	panic(srv.Run())
}
