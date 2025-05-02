package controller

import (
	"net/http"
	"strconv"

	"github.com/AndresBC-Dev/Astro-lab/model"
	"github.com/AndresBC-Dev/Astro-lab/service"
	"github.com/gin-gonic/gin"
)

// Definir la interfaz del controlador
type AstralEntityController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// Implementación concreta del controlador
type AstralEntityControllerImpl struct {
	Service service.AstralEntityService
}

// Constructor con inyección de dependencia
func NewAstralEntityController(service service.AstralEntityService) AstralEntityController {
	return &AstralEntityControllerImpl{Service: service}
}

// Crear entidad astral
func (ctrl *AstralEntityControllerImpl) Create(c *gin.Context) {
	var entity model.AstralEntity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Obtener el archivo del request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo requerido"})
		return
	}

	// Pasar el archivo al servicio para calcular su peso
	if err := ctrl.Service.Create(&entity, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entity)
}

// Obtener por ID
func (ctrl *AstralEntityControllerImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	entity, err := ctrl.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado"})
		return
	}

	c.JSON(http.StatusOK, entity)
}

// Obtener todos
func (ctrl *AstralEntityControllerImpl) GetAll(c *gin.Context) {
	entities, err := ctrl.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities)
}

// Actualizar entidad astral
func (ctrl *AstralEntityControllerImpl) Update(c *gin.Context) {
	var entity model.AstralEntity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := ctrl.Service.Update(&entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity)
}

// Eliminar entidad astral
func (ctrl *AstralEntityControllerImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.Service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Eliminado correctamente"})
}
