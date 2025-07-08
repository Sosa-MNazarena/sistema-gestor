package main

import (
	controllers "sistema-gestor/controllers"
	_ "sistema-gestor/docs"
	"sistema-gestor/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Sistema Gestor de Productos API
// @version 1.0
// @description Esta API permite gestionar productos con GORM y Gin.
// @contact.name Sosa, Mariana Nazarena
// @host localhost:8080
// @BasePath /
func main() {
	dsn := "root:@tcp(localhost:3306)/sistemagestor?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
    panic("No se puede conectar a la base de datos: " + err.Error())
}

	if err := db.AutoMigrate(&models.Product{}, &models.Stock{}); err != nil {
		panic("failed to migrate database")
	}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.InitDB(db)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
	router.GET("/excelReader", controllers.ReadExcelFile)


	router.Run(":8080")

}