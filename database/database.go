package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global para la conexión
var DB *gorm.DB

// Cargar variables de entorno desde .env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

// Conectar a PostgreSQL con GORM
func ConnectDB() {
	LoadEnv()

	// Crear cadena de conexión desde variables de entorno
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	// Establecer conexión con GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL")
}
