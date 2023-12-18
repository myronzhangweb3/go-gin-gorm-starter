package models

import (
	"gorm.io/gorm"
)

// StrategyInfo 策略信息
type StrategyInfo struct {
	gorm.Model
	Name string `gorm:"not null;comment:策略名称" json:"name"`
	Type uint   `gorm:"not null;comment:策略类型,例如质押、流动性挖矿等" json:"type,omitempty"`
}

func (StrategyInfo) TableName() string {
	return "strategy_info"
}
