package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AstralEntityController struct {
}

func (AEC *AstralEntityController) UploadFile(c *gin.Context) {
	// Obtener archivo del request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo recibir el archivo"})
		return
	}

	// Guardar el archivo en disco (puedes cambiar la ruta)
	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Archivo subido exitosamente", "filename": file.Filename})
}
