package repositories

import (
	"sistema-gestor/models"

	"gorm.io/gorm"
)

type ProductImportRepository interface {
	SaveImported(products []models.Product) error
	ExistsDuplicate(nombre, proveedor string) (bool, error)
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

func (r *productImportRepository) ExistsDuplicate(nombre, proveedor string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Product{}).
		Where("nombre = ? AND proveedor = ?", nombre, proveedor).
		Count(&count).Error
	return count > 0, err
}
