package model

import (
	"gorm.io/gorm"
)

// AstralEntity representa un cuerpo celeste en la API.
// Se utiliza para almacenar información sobre planetas, asteroides, estrellas, etc.
type AstralEntity struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" json:"name"` // Nombre del objeto astronómico
	Type        string  `gorm:"type:varchar(50);not null" json:"type"`  // Tipo (Ejemplo: "planeta", "estrella", "asteroide")
	Curiosity   string  `gorm:"type:text" json:"curiosity"`             // Dato curioso sobre el cuerpo celeste
	Position    string  `gorm:"type:varchar(100)" json:"position"`      // Ubicación relativa o coordenadas espaciales
	Size        float64 `gorm:"type:decimal(10,2)" json:"size"`         // Tamaño en kilómetros
	Temperature int     `gorm:"type:int" json:"temperature"`            // Temperatura en Kelvin (debe ser >= 0)
	Age         int     `gorm:"type:int" json:"age"`                    // Edad en millones de años
	FileSizeMB  float64 `gorm:"type:decimal(10,2)" json:"file_size_mb"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image_url"` // URL de la imagen en S3
	FileURL     string  `gorm:"type:varchar(255)" json:"file_url"`  // URL del archivo de Unity en S3
}
