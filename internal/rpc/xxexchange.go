package rpc

import (
	"encoding/json"
	"noneland/backend/interview/internal/repo/model"
)

type RpcClient interface {
	GetSpotBalance() (balance *model.RiskBalance, err error)
	GetFuturesBalance() (balance *model.RiskBalance, err error)
	GetTransferRecords() (records *model.TransferRecords, err error)
}

type xxExchangeClient struct {
	client interface{}
}

func NewXXExchangeClient() RpcClient {
	return &xxExchangeClient{}
}

// Simulate calls to /spot/balance endpoint
func (c *xxExchangeClient) GetSpotBalance() (balance *model.RiskBalance, err error) {
	resp := []byte(`{"free": "10.12345"}`)
	if err = json.Unmarshal(resp, &balance); err != nil {
		return nil, err
	}

	return
}

// Simulate calls to /futures/balance endpoint
func (c *xxExchangeClient) GetFuturesBalance() (balance *model.RiskBalance, err error) {
	resp := []byte(`{"free": "30.12345"}`)
	if err = json.Unmarshal(resp, &balance); err != nil {
		return nil, err
	}

	return
}

// Simulate calls to /spot/transfer/records endpoint
func (c *xxExchangeClient) GetTransferRecords() (records *model.TransferRecords, err error) {
	resp := []byte(`
		{
		   "rows": [
		      {
		         "amount": "0.10000000",
		         "asset": "BNB",
		         "status": "CONFIRMED",
		         "timestamp": 1566898617,
		         "txId": 5240372201,
		         "type": "IN"
		      },
		      {
		         "amount": "5.00000000",
		         "asset": "USDT",
		         "status": "CONFIRMED",
		         "timestamp": 1566888436,
		         "txId": 5239810406,
		         "type": "OUT"
		      },
		      {
		         "amount": "1.00000000",
		         "asset": "EOS",
		         "status": "CONFIRMED",
		         "timestamp": 1566888403,
		         "txId": 5239808703,
		         "type": "IN"
		      }
		   ],
		   "total": 3
		}`)
	if err = json.Unmarshal(resp, &records); err != nil {
		return nil, err
	}

	return
}
