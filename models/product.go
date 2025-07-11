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
	Nombre        string  `gorm:"not null;size:200" json:"nombre"`
	Descripcion string  `gorm:"not null;size:500" json:"descripcion"`
	Categoria	string  `gorm:"not null;size:100" json:"categoria"`
	Proveedor	 string  `gorm:"not null" json:"proveedor"`
	Precio       float64 `gorm:"not null;check:Precio>0" json:"precio"`
	Stocks    []Stock  `gorm:"foreignKey:ProductID;references:ID" json:"stocks"`
}
