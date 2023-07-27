// @title SPM
// @description Simple password management API
// @version 0.01
// @BasePath /
package main

import (
	"fmt"

	"github.com/jvfrodrigues/simple-password-manager-back/api"
	"github.com/jvfrodrigues/simple-password-manager-back/config"
)

func main() {
	serverPort := config.GetString("HTTP_PORT")
	server := api.NewServer()
	fmt.Println("Server starting at localhost:" + serverPort)
	server.StartServer(serverPort)
}
