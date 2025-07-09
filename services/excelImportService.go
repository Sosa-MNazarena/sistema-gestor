package services

import (
	"sistema-gestor/events"
	"sistema-gestor/models"
	"sistema-gestor/repositories"
	"sistema-gestor/strategy"
	"strconv"
)

//var ProductSvc ProductService

type ExcelImportService interface {
	ImportFromExcel(path string) ([]models.Product, error)
}

type excelImportService struct {
	repo repositories.ProductImportRepository
}

func NewExcelImportService(repo repositories.ProductImportRepository) ExcelImportService {
	return &excelImportService{repo}
}

func (s *excelImportService) ImportFromExcel(path string) ([]models.Product, error) {
	context := strategy.ReaderContext{}
	context.SetReader(&strategy.ExcelReader{Path: path})
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
	go events.SendLoadSuccessEmail("mariana.sosa@davinci.edu.ar")
	return products, nil

}