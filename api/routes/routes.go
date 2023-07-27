package api

import (
	"github.com/gin-gonic/gin"
	api "github.com/jvfrodrigues/simple-password-manager-back/api/routes/v1"
)

func SetupRoutes(router *gin.Engine) {
	api.SetupPasswordRoutes(router)
}
