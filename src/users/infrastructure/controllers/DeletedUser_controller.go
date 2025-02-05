package controllers

import (
	"net/http"
	"strconv"

	"demo/src/users/application"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteUserController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el usuario",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Usuario eliminado exitosamente",
	})
}
