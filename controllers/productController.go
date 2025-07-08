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
//@Failure 500 {object} string "Fallo al crear el producto"
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
//@Failure 400 {object} string "Error al obtener productos"
//@Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Preload("Stocks").Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener productos"})
		return
	}
	c.JSON(200, products)
}

//@Summary Obtener un producto por su ID
//@Tags Productos
//@Accept json
//@Produce json
//@Param id path string true "ID del producto"
//@Success 200 {object} models.Product
//@Failure 404 {object} string "Producto no encontrado"
//@Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Preload("Stocks").First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(200, product)
}

//@Summary Actualizar un producto
//@Tags Productos
//@Accept  json
//@Produce json
//@Param id path string true "ID del producto"
//@Param product body models.Product true "Producto a actualizar"
//@Success 200 {object} models.Product
//@Failure 404 {object} string "Producto no encontrado"
//@Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Producto no encontrado"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Model(&product).Updates(updatedProduct)

	c.JSON(200, product)
}

//@Summary Eliminar un producto
//@Tags Productos
//@Accept  json
//@Produce json
//@Param id path string true "ID del producto"
//@Success 204 {string} string ""
//@Failure 404 {object} string "Producto no encontrado"
//@Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Producto no encontrado"})
		return
	}

	db.Delete(&product)
	c.Status(204) //código de que no devuelve contenido
}