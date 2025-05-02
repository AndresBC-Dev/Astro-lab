package main

import (
	"github.com/AndresBC-Dev/Astro-lab/controller"
	"github.com/AndresBC-Dev/Astro-lab/database"
	"github.com/AndresBC-Dev/Astro-lab/repository"
	"github.com/AndresBC-Dev/Astro-lab/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	database.ConnectDB()

	// Crear instancia del repositorio, servicio y controlador para DI
	repo := repository.NewAstralEntityRepository(database.DB)
	service := service.NewAstralEntityService(repo)
	controller := controller.NewAstralEntityController(service)

	// Rutas con Gin
	r := gin.Default()
	r.POST("/planets", controller.Create)
	r.GET("/planets/:id", controller.GetByID)
	r.GET("/planets", controller.GetAll)
	r.PUT("/planets", controller.Update)
	r.DELETE("/planets/:id", controller.Delete)

	// Iniciar servidor
	r.Run(":8080")
}
