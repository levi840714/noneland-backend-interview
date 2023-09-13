//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/biz"
	"noneland/backend/interview/internal/pkg/db"
	"noneland/backend/interview/internal/repo"
	"noneland/backend/interview/internal/router"
	"noneland/backend/interview/internal/rpc"
)

func InitServer() *Server {
	wire.Build(
		db.InitGorm,
		api.ApiProviderSet,
		rpc.RpcProviderSet,
		repo.RepositoryProviderSet,
		biz.BizProviderSet,
		router.RouterProviderSet,
		NewGinEngine,
		NewServer)

	return &Server{}
}
