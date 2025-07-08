package controllers

import (
	"net/http"
	"sistema-gestor/strategy"
	"github.com/gin-gonic/gin"
	
)

//@Summary Leer un archivo Excel
//@Tags Excels
//@Accept  json
//@Produce json
//@Param path query string true "Ruta del archivo (absoluta)"
//@Success 200 {array} map[string]interface{}
//@Failure 400 {object} string "Error al leer el archivo"
//@Router /excelReader [get]
func ReadExcelFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el parámetro 'path'"})
		return
	}

	context := strategy.ReaderContext{}
	context.SetReader(&strategy.ExcelReader{Path: path})
	data, err := context.ProcessData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}