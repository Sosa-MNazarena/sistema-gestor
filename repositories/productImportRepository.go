package repositories

import (
	"sistema-gestor/models"
	"gorm.io/gorm"
)

type ProductImportRepository interface {
	SaveImported(products []models.Product) error
}

type productImportRepository struct {
	db *gorm.DB
}

func NewProductImportRepository(db *gorm.DB) ProductImportRepository {
	return &productImportRepository{db: db}
}

func (r *productImportRepository) SaveImported(products []models.Product) error {
	for _, product := range products {
		if err := r.db.Create(&product).Error; err != nil {
			return err
		}
	}
	return nil
}