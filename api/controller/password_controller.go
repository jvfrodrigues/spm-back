package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/simple-password-manager-back/data"
	dtos "github.com/jvfrodrigues/simple-password-manager-back/domain/dtos"
)

type PasswordController struct {
	repository *data.PasswordRepository
}

func NewPasswordController() *PasswordController {
	return &PasswordController{
		repository: data.NewPasswordRepository(),
	}
}

func (pc *PasswordController) CreatePasswordEntry(ctx *gin.Context) {
	var request dtos.CreatePasswordRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Error formatting data")
		return
	}

	pc.repository.Create(request)

	ctx.Status(http.StatusNoContent)
}

func (pc *PasswordController) GetAll(ctx *gin.Context) {
	passwords, err := pc.repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error getting passwords")
		return
	}
	ctx.JSON(http.StatusOK, passwords)
}
