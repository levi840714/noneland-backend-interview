package entity

import (
	"noneland/backend/interview/internal/repo/model"
	"time"
)

type (
	RiskBalance struct {
		Spot    string `json:"spot"`
		Futures string `json:"futures"`
	}

	TransferRecord struct {
		Amount    string    `json:"amount"`
		Asset     string    `json:"asset"`
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
		TxID      int64     `json:"txId"`
		Type      string    `json:"type"`
	}
)

// records status enum
const (
	Pending   = "PENDING"
	Confirmed = "CONFIRMED"
	Failed    = "FAILED"
)

var StatusEnum = map[string]string{
	Pending:   "等待",
	Confirmed: "成功",
	Failed:    "失敗",
}

func RecordsModelToEntity(records *model.TransferRecords) []*TransferRecord {
	entityRecords := make([]*TransferRecord, 0, len(records.Rows))
	for _, record := range records.Rows {
		entityRecords = append(entityRecords, &TransferRecord{
			Amount:    record.Amount,
			Asset:     record.Asset,
			Status:    StatusEnum[record.Status],
			Timestamp: time.Unix(int64(record.Timestamp), 0),
			TxID:      record.TxID,
			Type:      record.Type,
		})
	}

	return entityRecords
}
