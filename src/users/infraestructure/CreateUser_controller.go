package infrastructure

import (
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	CreateUserUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		CreateUserUseCase: createUserUseCase,
	}
}

func (ctrl *CreateUserController) Run(c *gin.Context) {
	var user entities.User

	if errJSON := c.ShouldBindJSON(&user); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del Usuario inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	userCreado, errAdd := ctrl.CreateUserUseCase.Run(&user)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el usuario",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "El usuario ha sido agregado",
		"empleado": userCreado,
	})
}
