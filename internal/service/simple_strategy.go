package service

import (
	"go-gin-gorm-starter/internal/dao"
	"go-gin-gorm-starter/internal/models"

	"gorm.io/gorm"
)

type SimpleStrategyService struct {
	strategyDao *dao.SimpleStrategyDao
}

func NewSimpleStrategyService(db *gorm.DB) *SimpleStrategyService {
	return &SimpleStrategyService{
		strategyDao: dao.NewSimpleStrategyDao(db),
	}
}

func (s *SimpleStrategyService) SaveSimpleStrategy(strategyInfo *models.SimpleStrategy) error {
	return s.strategyDao.SaveSimpleStrategy(strategyInfo)
}

func (s *SimpleStrategyService) FindSimpleStrategy(strategyInfo *models.SimpleStrategy) ([]models.SimpleStrategy, error) {
	return s.strategyDao.FindSimpleStrategy(strategyInfo)
}

func (s *SimpleStrategyService) DeleteSimpleStrategy(strategyInfo *models.SimpleStrategy) error {
	return s.strategyDao.DeleteSimpleStrategy(strategyInfo)
}

func (s *SimpleStrategyService) UpdateSimpleStrategy(strategyInfo *models.SimpleStrategy) (*models.SimpleStrategy, error) {
	return s.strategyDao.UpdateSimpleStrategy(strategyInfo)
}
