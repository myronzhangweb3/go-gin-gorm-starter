package models

import (
	"gorm.io/gorm"
)

type SimpleStrategy struct {
	gorm.Model
	Name string `gorm:"not null;comment:name" json:"name"`
	Type uint   `gorm:"not null;comment:type" json:"type,omitempty"`
}

func (SimpleStrategy) TableName() string {
	return "simple_strategy"
}
