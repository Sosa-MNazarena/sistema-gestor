package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"sistema-gestor/services"
	
)

var excelImportService services.ExcelImportService

func InitExcelImportService(service services.ExcelImportService) {
	excelImportService = service
}

//@Summary Leer un archivo Excel y guardarlo en la base de datos
//@Tags Excel
//@Accept  json
//@Produce  json
//@Param path query string true "Ruta del archivo Excel"
//@Success 200 {array} map[string]interface{}
//@Failure 400 {object} string "Error al leer el archivo"
//@Router /excelReader [post]
func ReadExcelFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el parámetro 'path'"})
		return
	}

	products, err := excelImportService.ImportFromExcel(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}