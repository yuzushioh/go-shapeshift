package shapeshift

// Transaction represents a transaction on shapeshift
type Transaction struct {
	CurIn     string  `json:"curIn"`
	CurOut    string  `json:"curOut"`
	Timestamp float64 `json:"timestamp"`
	Amount    float64 `json:"amount"`
	TxID      int64   `json:"txid"`
}
