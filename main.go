package main

import (
	"SoftwareEngine/internal/server"
	"os"
)

// go:generate swag init --parseDependency --parseInternal
// @title online practice server
// @version 1.0
// @description 向前端提供api

// @contact.name 软工小组
// @contact.url http://www.swagger.io/support
// @contact.email 2020111058@email.szu.edu.cn

// @host localhost:8080
// @BasePath /v1
func main() {

	command := server.NewGoServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
