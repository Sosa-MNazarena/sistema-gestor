package services

import (
	"sistema-gestor/events"
	"sistema-gestor/models"
	"sistema-gestor/repositories"
	"sistema-gestor/strategy"
	"strconv"
)

type ApiImportService interface {
	ImportFromApi(baseURL string) ([]models.Product, error)
}

type apiImportService struct {
	repo repositories.ProductImportRepository
}

func NewApiImportService(repo repositories.ProductImportRepository) ApiImportService {
	return &apiImportService{repo}
}

func (s *apiImportService) ImportFromApi(baseURL string) ([]models.Product, error) {
	context := strategy.ReaderContext{}
	context.SetReader(&strategy.ApiReader{BaseURL: baseURL})
	data, err := context.ProcessData()
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for _, entry := range data {
		product := models.Product{
			Nombre:      entry["nombre"].(string),
			Descripcion: entry["descripcion"].(string),
			Categoria:   entry["categoria"].(string),
			Proveedor:   entry["proveedor"].(string),
		}

		//en caso de que el precio sea un string, se parsea a float básicamente
		switch v := entry["precio"].(type) {
		case float64:
			product.Precio = v
		case string:
			f, err := strconv.ParseFloat(v, 64)
			if err == nil {
				product.Precio = f
			}
		}
		products = append(products, product)
	}
	err = s.repo.SaveImported(products)
	if err != nil {
		return nil, err
	}
	go events.SendCustomEmail("mariana.sosa@davinci.edu.ar", "Carga desde Excel", "Se han cargado los productos desde la API correctamente.")
	return products, nil
}