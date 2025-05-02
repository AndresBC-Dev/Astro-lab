package service

import (
	"fmt"
	"mime/multipart"

	"github.com/AndresBC-Dev/Astro-lab/model"
	"github.com/AndresBC-Dev/Astro-lab/repository"
)

// Definir la interfaz del servicio
type AstralEntityService interface {
	Create(entity *model.AstralEntity, file *multipart.FileHeader) error
	GetByID(id int) (*model.AstralEntity, error)
	GetAll() ([]model.AstralEntity, error)
	Update(entity *model.AstralEntity) error
	Delete(id int) error
}

// Implementación del servicio usando la interfaz de repositorio
type AstralEntityServiceImpl struct {
	Repo repository.AstralEntityRepository
}

// Constructor con inyección de dependencia
func NewAstralEntityService(repo repository.AstralEntityRepository) AstralEntityService {
	return &AstralEntityServiceImpl{Repo: repo}
}

func (s *AstralEntityServiceImpl) Create(entity *model.AstralEntity, file *multipart.FileHeader) error {
	if entity.Name == "" {
		return fmt.Errorf("El nombre no puede estar vacío")
	}

	// Calcular tamaño del archivo en MB
	entity.FileSizeMB = float64(file.Size) / (1024 * 1024) // Convertir de bytes a MB

	return s.Repo.Create(entity)
}

func (s *AstralEntityServiceImpl) GetByID(id int) (*model.AstralEntity, error) {
	return s.Repo.GetByID(id)
}

func (s *AstralEntityServiceImpl) GetAll() ([]model.AstralEntity, error) {
	return s.Repo.GetAll()
}

func (s *AstralEntityServiceImpl) Update(entity *model.AstralEntity) error {
	if entity.ID == 0 {
		return fmt.Errorf("ID inválido para actualizar")
	}
	return s.Repo.Update(entity)
}

func (s *AstralEntityServiceImpl) Delete(id int) error {
	return s.Repo.Delete(id)
}
