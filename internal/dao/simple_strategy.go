package dao

import (
	"go-gin-gorm-starter/internal/models"
	"gorm.io/gorm"
)

type SimpleStrategyDao struct {
	db *gorm.DB
}

func NewSimpleStrategyDao(db *gorm.DB) *SimpleStrategyDao {
	return &SimpleStrategyDao{
		db: db,
	}
}

func (d *SimpleStrategyDao) SaveSimpleStrategy(strategyInfo *models.SimpleStrategy) error {
	return d.db.Save(strategyInfo).Error
}

func (d *SimpleStrategyDao) DeleteSimpleStrategy(strategyInfo *models.SimpleStrategy) error {
	return d.db.Where(strategyInfo).Delete(strategyInfo).Error
}

func (d *SimpleStrategyDao) FindSimpleStrategy(info *models.SimpleStrategy) ([]models.SimpleStrategy, error) {
	var infos []models.SimpleStrategy
	err := d.db.Where(info).Find(&infos).Error
	return infos, err
}

func (d *SimpleStrategyDao) UpdateSimpleStrategy(info *models.SimpleStrategy) (*models.SimpleStrategy, error) {
	err := d.db.Where("id = ?", info.ID).Updates(info).Error
	return info, err
}
