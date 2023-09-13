package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"noneland/backend/interview/config"
	"noneland/backend/interview/internal/router"
	"time"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *router.Router
}

func (s *Server) Start() *http.Server {
	appRouter := s.engine.Group("/app")
	s.apiRouter.Register(appRouter)

	return &http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.Conf.Host, config.Conf.Port),
		Handler:        s.engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}
func NewServer(engine *gin.Engine, apiRouter *router.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

func NewGinEngine() *gin.Engine {
	var ginMode string
	if config.Conf.Debug {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
	router := gin.New()
	router.Use(gin.Recovery())

	return router
}
