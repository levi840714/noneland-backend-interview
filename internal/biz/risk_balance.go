package biz

import (
	"golang.org/x/net/context"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/repo/model"
)

type (
	RiskUseCase interface {
		// get all balance from xx exchange
		GetRiskBalance(ctx context.Context) (*entity.RiskBalance, error)
	}

	RiskRepo interface {
		// get spot balance from xx exchang
		GetSpotBalance() (riskBalance *model.RiskBalance, err error)
		// get futures balance from xx exchang
		GetFuturesBalance() (riskBalance *model.RiskBalance, err error)
	}
)

type riskUseCase struct {
	riskRepo RiskRepo
}

func NewRiskUseCase(riskRepo RiskRepo) RiskUseCase {
	return &riskUseCase{
		riskRepo: riskRepo,
	}
}

func (u *riskUseCase) GetRiskBalance(ctx context.Context) (*entity.RiskBalance, error) {
	spotBalance, err := u.riskRepo.GetSpotBalance()
	if err != nil {
		return nil, err
	}

	futuresBalance, err := u.riskRepo.GetFuturesBalance()
	if err != nil {
		return nil, err
	}

	data := &entity.RiskBalance{
		Spot:    spotBalance.Free,
		Futures: futuresBalance.Free,
	}

	return data, nil
}
