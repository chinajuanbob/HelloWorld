package service

import (
	"github.com/chinajuanbob/helloworld/pkg/common"
	"github.com/gin-gonic/gin"
)

type HealthService struct {
	common.HttpService
}

func NewHealthService() *HealthService {
	s := HealthService{}
	s.HttpService.GetStatus = s.GetStatus
	s.HttpService.GetHealthz = s.GetHealthz
	s.Init()
	return &s
}

func (s *HealthService) GetHealthz() int {
	return 1
}

func (s *HealthService) GetStatus(detail bool) gin.H {
	return gin.H{
		"a": 100,
	}
}
