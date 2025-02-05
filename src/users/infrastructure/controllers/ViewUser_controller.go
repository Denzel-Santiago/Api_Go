package controllers

import (
	"net/http"
	"strconv"

	"demo/src/users/application"
	"github.com/gin-gonic/gin"
)

type ViewUserController struct {
	useCase *application.ViewUser
}

func NewViewUserController(useCase *application.ViewUser) *ViewUserController {
	return &ViewUserController{useCase: useCase}
}

func (vuc *ViewUserController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := vuc.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}
