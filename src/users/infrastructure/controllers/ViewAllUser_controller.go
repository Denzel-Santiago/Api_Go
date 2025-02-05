package controllers

import (
	"net/http"

	"demo/src/users/application"
	"github.com/gin-gonic/gin"
)

type ViewAllUsersController struct {
	useCase *application.ViewAllUsers
}

func NewViewAllUsersController(useCase *application.ViewAllUsers) *ViewAllUsersController {
	return &ViewAllUsersController{useCase: useCase}
}

func (vuc *ViewAllUsersController) Execute(c *gin.Context) {
	users, err := vuc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, users)
}
