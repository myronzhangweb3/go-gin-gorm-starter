package service

import (
	"go-gin-gorm-starter/dao"
	"go-gin-gorm-starter/moduls/moduls_db"
)

func SaveSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) error {
	return dao.SaveSimpleStrategy(strategyInfo)
}

func FindSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) ([]moduls_db.StrategyInfo, error) {
	return dao.FindSimpleStrategy(strategyInfo)
}

func DeleteSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) error {
	return dao.DeleteSimpleStrategy(strategyInfo)
}

func UpdateSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) (*moduls_db.StrategyInfo, error) {
	return dao.UpdateSimpleStrategy(strategyInfo)
}
