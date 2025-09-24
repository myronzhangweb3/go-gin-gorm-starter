package controller

import (
	"go-gin-gorm-starter/internal/models"
	"go-gin-gorm-starter/internal/service"
	"go-gin-gorm-starter/pkg/gin2"
	"go-gin-gorm-starter/pkg/time2"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SimpleStrategyController struct {
	log             *zap.Logger
	strategyService *service.SimpleStrategyService
}

func NewSimpleStrategyController(log *zap.Logger, db *gorm.DB) *SimpleStrategyController {
	return &SimpleStrategyController{
		log:             log,
		strategyService: service.NewSimpleStrategyService(db),
	}
}

func (s *SimpleStrategyController) GetSimpleStrategy(ctx *gin.Context) {
	defer time2.TimeConsume(time.Now())

	address := ctx.Query("address")

	data, err := s.strategyService.FindSimpleStrategy(&models.SimpleStrategy{
		Name: address,
	})
	if err != nil {
		gin2.HttpResponse(s.log, ctx, "", err)
		return
	}

	gin2.HttpResponse(s.log, ctx, data, err)
}

func (s *SimpleStrategyController) SaveSimpleStrategy(ctx *gin.Context) {
	var (
		reqUser models.SimpleStrategy
	)

	if err := ctx.ShouldBindJSON(&reqUser); err != nil {
		gin2.HttpResponse(s.log, ctx, nil, gin2.CommonError)
		return
	}

	err := s.strategyService.SaveSimpleStrategy(&reqUser)
	if err != nil {
		gin2.HttpResponse(s.log, ctx, "", err)
		return
	}

	gin2.HttpResponse(s.log, ctx, err == nil, err)
}

func (s *SimpleStrategyController) DeleteSimpleStrategy(ctx *gin.Context) {
	var (
		params models.SimpleStrategy
	)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		gin2.HttpResponse(s.log, ctx, "", err)
		return
	}

	params.ID = uint(id)

	err = s.strategyService.DeleteSimpleStrategy(&params)
	if err != nil {
		gin2.HttpResponse(s.log, ctx, "", err)
		return
	}

	gin2.HttpResponse(s.log, ctx, err == nil, err)
}
