package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name" binding:"required"`
	AreaID    uint   `json:"-"`
	Area      Area   `json:"area"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName is Database TableName of this model
func (e *Location) TableName() string {
	return "locations"
}
