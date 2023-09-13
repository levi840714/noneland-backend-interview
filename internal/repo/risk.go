package repo

import (
	"gorm.io/gorm"
	"noneland/backend/interview/internal/biz"
	"noneland/backend/interview/internal/repo/model"
	"noneland/backend/interview/internal/rpc"
)

type riskRepo struct {
	db            *gorm.DB
	xxExchangeRpc rpc.RpcClient
}

func NewRiskRepo(db *gorm.DB, xxExchangeRpc rpc.RpcClient) biz.RiskRepo {
	return &riskRepo{
		db:            db,
		xxExchangeRpc: xxExchangeRpc,
	}
}

func (r *riskRepo) GetSpotBalance() (balance *model.RiskBalance, err error) {
	balance, err = r.xxExchangeRpc.GetSpotBalance()
	if err != nil {
		return nil, err
	}

	return balance, err
}

func (r *riskRepo) GetFuturesBalance() (balance *model.RiskBalance, err error) {
	balance, err = r.xxExchangeRpc.GetFuturesBalance()
	if err != nil {
		return nil, err
	}

	return balance, err
}

func (r *riskRepo) GetTransferRecords() (records *model.TransferRecords, err error) {
	records, err = r.xxExchangeRpc.GetTransferRecords()
	if err != nil {
		return nil, err
	}

	return records, err
}
