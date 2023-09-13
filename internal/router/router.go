package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"noneland/backend/interview/internal/api"
)

var swagHandler gin.HandlerFunc

type Router struct {
	riskController *api.RiskController
}

var RouterProviderSet = wire.NewSet(
	NewRouter,
)

func NewRouter(riskController *api.RiskController) *Router {
	return &Router{
		riskController: riskController,
	}
}

func (r *Router) Register(rootGroup *gin.RouterGroup) {
	if swagHandler != nil {
		rootGroup.GET("v1/swagger/*any", swagHandler)
	}

	r.setRiskRoutes(rootGroup)
}
