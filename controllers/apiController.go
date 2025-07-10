package controllers

import (
	"net/http"
	"sistema-gestor/services"
	"github.com/gin-gonic/gin"
)

var apiImportService services.ApiImportService

func InitApiImportService(service services.ApiImportService) {
	apiImportService = service
}

//@Summary Leer un archivo de datos de una API
//@Tags APIs
//@Accept  json
//@Produce json
//@Param baseURL query string true "URL de la API"
//@Success 200 {array} map[string]interface{}
//@Failure 400 {object} string "Error al leer el archivo"
//@Router /apiReader [get]
func ReadApiFile(c *gin.Context) {
	baseURL := c.Query("baseURL")
	if baseURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el parámetro 'baseURL'"})
		return
	}

	products, err := apiImportService.ImportFromApi(baseURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}