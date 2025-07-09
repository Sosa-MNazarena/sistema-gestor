package repositories

import (
	"sistema-gestor/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(product *models.Product) error
}

type productoRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productoRepository{db: db}
}

func (r *productoRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productoRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Stocks").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productoRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Stocks").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productoRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productoRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}