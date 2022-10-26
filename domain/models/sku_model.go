package models

import (
	"time"

	"gorm.io/gorm"
)

type Sku struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName is Database TableName of this model
func (e *Sku) TableName() string {
	return "skus"
}
