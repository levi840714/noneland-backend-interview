//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/db"
	"noneland/backend/interview/internal/entity"
	repo "noneland/backend/interview/internal/repo/gorm"
	"sync"
)

var cg *configs.Config
var configOnce sync.Once

func NewConfig() *configs.Config {
	configOnce.Do(func() {
		cg = configs.NewConfig()
	})

	return cg
}

var database *gorm.DB
var dbOnce sync.Once

func NewDB() *gorm.DB {
	dbOnce.Do(func() {
		database = db.NewDb()
	})

	return database
}

func NewRepo() (entity.Repository, error) {
	panic(wire.Build(repo.NewRepository, NewDB, NewConfig))
}
