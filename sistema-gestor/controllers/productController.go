package controllers

import (
	"gorm.io/gorm"
	"sistema-gestor/models"
	"github.com/gin-gonic/gin"
	
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

//@Summary Crear un producto nuevo
//@Tags Productos
//@Accept json
//@Produce json
//@Param product body models.Product true "Producto a crear"
//@Success 201 {object} models.Product
//@Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Fallo al crear el producto"})
		return
	}

	c.JSON(201, product)
}

//@Summary Obtener todos los productos
//@Tags Productos	
//@Accept json
//@Produce json
//@Success 200 {array} models.Product
//@Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Preload("Stocks").Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener productos"})
		return
	}
	c.JSON(200, products)
}
