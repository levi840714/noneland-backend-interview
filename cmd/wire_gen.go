// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/biz"
	"noneland/backend/interview/internal/pkg/db"
	"noneland/backend/interview/internal/repo"
	"noneland/backend/interview/internal/router"
)

// Injectors from wire.go:

func InitServer() *Server {
	engine := NewGinEngine()
	gormDB := db.InitGorm()
	riskRepo := repo.NewRiskRepo(gormDB)
	riskUseCase := biz.NewRiskUseCase(riskRepo)
	riskController := api.NewRiskController(riskUseCase)
	routerRouter := router.NewRouter(riskController)
	server := NewServer(engine, routerRouter)
	return server
}