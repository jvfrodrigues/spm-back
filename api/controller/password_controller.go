package api

import (
	"fmt"
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
	err := pc.repository.Create(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error creating password - %v", err))
		return
	}
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

func (pc *PasswordController) DeletePasswordEntry(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.repository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error deleting password - %v", err))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (pc *PasswordController) EditPasswordEntry(ctx *gin.Context) {
	id := ctx.Param("id")
	var request dtos.CreatePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "Error formatting data")
		return
	}
	err := pc.repository.Update(id, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error deleting password - %v", err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
