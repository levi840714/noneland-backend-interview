package gorm

import (
	"gorm.io/gorm"
	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
)

type repository struct {
	db     *gorm.DB
	config *configs.Config
}

func NewRepository(db *gorm.DB, config *configs.Config) entity.Repository {
	return &repository{
		db:     db,
		config: config,
	}
}
