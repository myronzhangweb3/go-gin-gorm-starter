package service

import (
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/models"
)

func SaveSimpleStrategy(strategyInfo *models.Demo) error {
	return dao.SaveSimpleStrategy(strategyInfo)
}

func FindSimpleStrategy(strategyInfo *models.Demo) ([]models.Demo, error) {
	return dao.FindSimpleStrategy(strategyInfo)
}

func DeleteSimpleStrategy(strategyInfo *models.Demo) error {
	return dao.DeleteSimpleStrategy(strategyInfo)
}

func UpdateSimpleStrategy(strategyInfo *models.Demo) (*models.Demo, error) {
	return dao.UpdateSimpleStrategy(strategyInfo)
}
