package api

import (
	"github.com/gin-gonic/gin"
)

var basicPasswordRoute = v1BasicRoute() + "/password-cards"

func setupPasswordRoutes(router *gin.Engine) {
	router.GET(basicPasswordRoute)
	router.POST(basicPasswordRoute)
}
