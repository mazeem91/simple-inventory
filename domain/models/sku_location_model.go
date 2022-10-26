package models

import (
	"time"

	"gorm.io/gorm"
)

type SkuLocation struct {
	SkuID      uint     `json:"-"`
	Sku        Sku      `json:"sku"`
	LocationID uint     `json:"-"`
	Location   Location `json:"location"`
	Quantity   int      `json:"quantity" binding:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName is Database TableName of this model
func (e *SkuLocation) TableName() string {
	return "sku_location"
}
