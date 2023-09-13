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
		// get all transfer records from xx exchange
		GetTransferRecords(ctx context.Context) ([]*entity.TransferRecord, error)
	}

	RiskRepo interface {
		// get spot balance from xx exchange
		GetSpotBalance() (riskBalance *model.RiskBalance, err error)
		// get futures balance from xx exchange
		GetFuturesBalance() (riskBalance *model.RiskBalance, err error)
		// get transfer records from xx exchange
		GetTransferRecords() (records *model.TransferRecords, err error)
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

func (u *riskUseCase) GetTransferRecords(ctx context.Context) ([]*entity.TransferRecord, error) {
	transferRecords, err := u.riskRepo.GetTransferRecords()
	if err != nil {
		return nil, err
	}

	return entity.RecordsModelToEntity(transferRecords), nil
}
