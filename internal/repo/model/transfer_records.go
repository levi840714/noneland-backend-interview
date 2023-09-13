package model

type (
	Record struct {
		Amount    string `json:"amount"`
		Asset     string `json:"asset"`
		Status    string `json:"status"`
		Timestamp int    `json:"timestamp"`
		TxID      int64  `json:"txId"`
		Type      string `json:"type"`
	}

	TransferRecords struct {
		Rows  []Record `json:"rows"`
		Total int      `json:"total"`
	}
)
