package dao

import (
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/models"
)

func SaveSimpleStrategy(strategyInfo *models.Demo) error {
	return global.DB.Save(strategyInfo).Error
}

func DeleteSimpleStrategy(strategyInfo *models.Demo) error {
	return global.DB.Where(strategyInfo).Delete(strategyInfo).Error
}

func FindSimpleStrategy(info *models.Demo) ([]models.Demo, error) {
	var infos []models.Demo
	err := global.DB.Where(info).Find(&infos).Error
	return infos, err
}

func UpdateSimpleStrategy(info *models.Demo) (*models.Demo, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
