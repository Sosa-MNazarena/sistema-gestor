package models

import (
	"time"
	"gorm.io/gorm"
)

// Product representa un producto en el sistema.
// swagger: model Product
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Nombre        string  `gorm:"not null" json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Categoria	string  `json:"categoria"`
	Proveedor	 string  `json:"proveedor"`
	Precio       float64 `gorm:"not null" json:"precio"`
	Stocks    []Stock  `gorm:"foreignKey:ProductID;references:ID" json:"stocks"`
}
