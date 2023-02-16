package main

import (
	"github.com/mfitrahrmd/BTPN_Syariah-Image_Uploader/server"
	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.BuildServer()
	if err != nil {
		logrus.Fatalln(err)
	}

	err = srv.Run()
	if err != nil {
		logrus.Fatalln(err)
	}
}
