package controllers

import (
	"demo/src/videojuego/application"
	"demo/src/videojuego/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideojuegoController struct {
	CreateUsecase *application.ManageVideojuegosUsecase
	GetUsecase    *application.GetVideojuegoUsecase
	UpdateUsecase *application.UpdateVideojuegoUsecase
	DeleteUsecase *application.DeleteVideojuegoUsecase
}

func NewVideojuegoController(createUC *application.ManageVideojuegosUsecase, getUC *application.GetVideojuegoUsecase, updateUC *application.UpdateVideojuegoUsecase, deleteUC *application.DeleteVideojuegoUsecase) *VideojuegoController {
	return &VideojuegoController{
		CreateUsecase: createUC,
		GetUsecase:    getUC,
		UpdateUsecase: updateUC,
		DeleteUsecase: deleteUC,
	}
}

func (vc *VideojuegoController) CreateVideojuego(c *gin.Context) {
	var videojuego entities.Videojuego
	if err := c.ShouldBindJSON(&videojuego); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := vc.CreateUsecase.CreateVideojuego(&videojuego); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el videojuego"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Videojuego creado"})
}

func (vc *VideojuegoController) GetVideojuego(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	videojuego, err := vc.GetUsecase.GetVideojuegoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Videojuego no encontrado"})
		return
	}
	c.JSON(http.StatusOK, videojuego)
}

func (vc *VideojuegoController) UpdateVideojuego(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var videojuego entities.Videojuego
	if err := c.ShouldBindJSON(&videojuego); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videojuego.ID = id // Asegurar que el ID en la URL se usa correctamente

	if err := vc.UpdateUsecase.UpdateVideojuego(&videojuego); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el videojuego", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Videojuego actualizado correctamente"})
}

func (vc *VideojuegoController) DeleteVideojuego(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := vc.DeleteUsecase.DeleteVideojuego(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el videojuego"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Videojuego eliminado"})
}
