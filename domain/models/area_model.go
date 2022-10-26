package models

import (
	"time"

	"gorm.io/gorm"
)

type Area struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Locations []Location     `json:"-"`
}

// TableName is Database TableName of this model
func (e *Area) TableName() string {
	return "areas"
}
