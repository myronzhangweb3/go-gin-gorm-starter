package service

import (
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/models"
)

func SaveSimpleStrategy(strategyInfo *models.StrategyInfo) error {
	return dao.SaveSimpleStrategy(strategyInfo)
}

func FindSimpleStrategy(strategyInfo *models.StrategyInfo) ([]models.StrategyInfo, error) {
	return dao.FindSimpleStrategy(strategyInfo)
}

func DeleteSimpleStrategy(strategyInfo *models.StrategyInfo) error {
	return dao.DeleteSimpleStrategy(strategyInfo)
}

func UpdateSimpleStrategy(strategyInfo *models.StrategyInfo) (*models.StrategyInfo, error) {
	return dao.UpdateSimpleStrategy(strategyInfo)
}
