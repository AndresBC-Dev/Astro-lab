package repository

import (
	"github.com/AndresBC-Dev/Astro-lab/model"

	"gorm.io/gorm"
)

// Interfaz para el repositorio
type AstralEntityRepository interface {
	Create(entity *model.AstralEntity) error
	GetByID(id int) (*model.AstralEntity, error)
	GetAll() ([]model.AstralEntity, error)
	Update(entity *model.AstralEntity) error
	Delete(id int) error
}

// Implementación concreta con GORM
type AstralEntityRepositoryImpl struct {
	DB *gorm.DB
}

// Constructor con inyección de dependencia (devolver *AstralEntityRepositoryImpl)
func NewAstralEntityRepository(db *gorm.DB) *AstralEntityRepositoryImpl {
	return &AstralEntityRepositoryImpl{DB: db}
}

// Métodos de la implementación
func (r *AstralEntityRepositoryImpl) Create(entity *model.AstralEntity) error {
	return r.DB.Create(entity).Error
}

func (r *AstralEntityRepositoryImpl) GetByID(id int) (*model.AstralEntity, error) {
	var entity model.AstralEntity
	err := r.DB.First(&entity, id).Error
	return &entity, err
}

func (r *AstralEntityRepositoryImpl) GetAll() ([]model.AstralEntity, error) {
	var entities []model.AstralEntity
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *AstralEntityRepositoryImpl) Update(entity *model.AstralEntity) error {
	return r.DB.Save(entity).Error
}

func (r *AstralEntityRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&model.AstralEntity{}, id).Error
}
