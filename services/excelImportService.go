package services

import (
	"fmt"
	"sistema-gestor/events"
	"sistema-gestor/models"
	"sistema-gestor/repositories"
	"sistema-gestor/strategy"
	"strconv"
)

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
	duplicados := 0
	for i, entry := range data {
		nombre, ok1 := entry["nombre"].(string)
		descripcion, ok2 := entry["descripcion"].(string)
		categoria, ok3 := entry["categoria"].(string)
		proveedor, ok4 := entry["proveedor"].(string)
		precioRaw, ok5 := entry["precio"]

		if !ok1 || nombre == "" || !ok2 || !ok3 || !ok4 || !ok5 {
			return nil, fmt.Errorf("producto %d: campos faltantes o vacíos", i+1)
		}

		var precio float64
		switch v := precioRaw.(type) {
		case float64:
			precio = v
		case string:
			p, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("producto %d: precio inválido", i+1)
			}
			precio = p
		default:
			return nil, fmt.Errorf("producto %d: tipo de precio desconocido", i+1)
		}

		if precio <= 0 {
			return nil, fmt.Errorf("producto %d: el precio debe ser mayor a cero", i+1)
		}

		//duplicado de prod combo proveedor+nombre
		exists, err := s.repo.ExistsDuplicate(nombre, proveedor)
		if err != nil {
			return nil, fmt.Errorf("producto %d: error al verificar duplicados", i+1)
		}
		if exists {
			duplicados++
			continue
		}

		product := models.Product{
			Nombre:      nombre,
			Descripcion: descripcion,
			Categoria:   categoria,
			Proveedor:   proveedor,
			Precio:      precio,
		}
		products = append(products, product)
	}

	if err := s.repo.SaveImported(products); err != nil {
		return nil, err
	}

	msg := "Se han leído los datos del Excel y actualizado el sistema correctamente."
	if duplicados > 0 {
	msg += fmt.Sprintf(" Se descartaron %d productos duplicados.", duplicados)
	}
	go events.SendCustomEmail("mariana.sosa@davinci.edu.ar", "Carga desde Excel", msg)
	return products, nil
}
