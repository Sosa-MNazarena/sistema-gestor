package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ProductID 	uint	`json:"product_id"`
	Sucursal  	string  `json:"sucursal"`
	Cantidad  	int     `json:"cantidad"`
}