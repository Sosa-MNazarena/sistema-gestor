package controllers

import (
	"github.com/gin-gonic/gin"
	"sistema-gestor/strategy"
)

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
		c.JSON(400, gin.H{"error": "Falta el parámetro 'baseURL'"})
		return
	}

	context := strategy.ReaderContext{}
	context.SetReader(&strategy.ApiReader{BaseURL: baseURL})

	data, err := context.ProcessData()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}