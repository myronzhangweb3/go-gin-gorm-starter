package models

import (
	"gorm.io/gorm"
)

// Demo demo
type Demo struct {
	gorm.Model
	Name string `gorm:"not null;comment:名称" json:"name"`
	Type uint   `gorm:"not null;comment:类型" json:"type,omitempty"`
}

func (Demo) TableName() string {
	return "demo"
}
