package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/simple-password-manager-back/config"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	setServerMode()
	server := &Server{router: setupRouter()}
	return server
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	return router
}

func setServerMode() {
	mode := config.GetString("API_MODE")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == gin.DebugMode {
		gin.SetMode(mode)
	}
}
