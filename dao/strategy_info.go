package dao

import (
	"go-gin-gorm-starter/global"
	"go-gin-gorm-starter/moduls/moduls_db"
)

func SaveSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) error {
	return global.DB.Save(strategyInfo).Error
}

func DeleteSimpleStrategy(strategyInfo *moduls_db.StrategyInfo) error {
	return global.DB.Where(strategyInfo).Delete(strategyInfo).Error
}

func FindSimpleStrategy(info *moduls_db.StrategyInfo) ([]moduls_db.StrategyInfo, error) {
	var infos []moduls_db.StrategyInfo
	err := global.DB.Where(info).Find(&infos).Error
	return infos, err
}

func UpdateSimpleStrategy(info *moduls_db.StrategyInfo) (*moduls_db.StrategyInfo, error) {
	err := global.DB.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
