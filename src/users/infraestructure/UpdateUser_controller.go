package infrastructure

import (
	"net/http"
	"strconv"

	"demo/src/users/application"
	"demo/src/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase *application.UpdateUser
}

func NewUpdateUserController(useCase *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{useCase: useCase}
}

func (uuc *UpdateUserController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uuc.useCase.Execute(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}
