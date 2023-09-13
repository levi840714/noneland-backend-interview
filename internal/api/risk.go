package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"noneland/backend/interview/internal/biz"
	"noneland/backend/interview/internal/pkg/logger"
)

type RiskController struct {
	riskUseCase biz.RiskUseCase
}

func NewRiskController(riskUseCase biz.RiskUseCase) *RiskController {
	return &RiskController{
		riskUseCase: riskUseCase,
	}
}

func (r *RiskController) RiskBalance(c *gin.Context) {
	ctx := context.Background()

	riskBalance, err := r.riskUseCase.GetRiskBalance(ctx)
	if err != nil {
		logger.Zap().Errorf("get risk balance error: %s", err)
		internalError(c)
		return
	}

	jsonResponse(c, riskBalance)
}
