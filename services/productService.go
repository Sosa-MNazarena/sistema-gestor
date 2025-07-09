package services

import (
	"sistema-gestor/models"
	"sistema-gestor/repositories"
)

type ProductService interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Update(id string, updated *models.Product) (*models.Product, error)
	Delete(id string) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s*productService) Update(id string, updated *models.Product) (*models.Product, error) {
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existingProduct.Nombre = updated.Nombre
	existingProduct.Descripcion = updated.Descripcion
	existingProduct.Categoria = updated.Categoria
	existingProduct.Proveedor = updated.Proveedor
	existingProduct.Precio = updated.Precio

	err = s.repo.Update(existingProduct)
	if err != nil {
		return nil, err
	}

	return existingProduct, nil
}

func (s *productService) Delete(id string) error {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(product)
}