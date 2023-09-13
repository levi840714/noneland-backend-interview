package router

import "github.com/gin-gonic/gin"

func (r *Router) setRiskRoutes(router *gin.RouterGroup) {
	riskV1Group := router.Group("v1/risk")

	riskV1Group.POST("risk_balance", r.riskController.RiskBalance)
	riskV1Group.POST("transfer_records", r.riskController.TransferRecords)
}
