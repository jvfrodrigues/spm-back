package api

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jvfrodrigues/simple-password-manager-back/api/controller"
)

var basicPasswordRoute = v1BasicRoute() + "/password-cards"

func SetupPasswordRoutes(router *gin.Engine) {
	controller := controllers.NewPasswordController()
	router.GET(basicPasswordRoute, controller.GetAll)
	router.POST(basicPasswordRoute, controller.CreatePasswordEntry)
	router.PUT(basicPasswordRoute+"/:id", controller.EditPasswordEntry)
	router.DELETE(basicPasswordRoute+"/:id", controller.DeletePasswordEntry)
}
