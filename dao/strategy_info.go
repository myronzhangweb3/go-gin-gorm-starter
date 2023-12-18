package dao

import (
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
)

func SaveSimpleStrategy(strategyInfo *models.StrategyInfo) error {
	return global.DB.Save(strategyInfo).Error
}

func DeleteSimpleStrategy(strategyInfo *models.StrategyInfo) error {
	return global.DB.Where(strategyInfo).Delete(strategyInfo).Error
}

func FindSimpleStrategy(info *models.StrategyInfo) ([]models.StrategyInfo, error) {
	var infos []models.StrategyInfo
	err := global.DB.Where(info).Find(&infos).Error
	return infos, err
}

func UpdateSimpleStrategy(info *models.StrategyInfo) (*models.StrategyInfo, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
