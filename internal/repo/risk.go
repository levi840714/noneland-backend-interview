package repo

import (
	"gorm.io/gorm"
	"noneland/backend/interview/internal/biz"
	"noneland/backend/interview/internal/repo/model"
)

type riskRepo struct {
	db *gorm.DB
}

func NewRiskRepo(db *gorm.DB) biz.RiskRepo {
	return &riskRepo{
		db: db,
	}
}

func (r *riskRepo) GetSpotBalance() (balance *model.RiskBalance, err error) {
	balance = &model.RiskBalance{Free: "10000"}

	return balance, err
}

func (r *riskRepo) GetFuturesBalance() (balance *model.RiskBalance, err error) {
	balance = &model.RiskBalance{Free: "20000"}

	return balance, err
}
