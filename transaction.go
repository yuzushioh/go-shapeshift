package shapeshift

// Transaction represents a transaction on shapeshift
type Transaction struct {
	CurIn     string  `json:"curIn,omitempty"`
	CurOut    string  `json:"curOut,omitempty"`
	Timestamp float64 `json:"timestamp,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
	TxID      int64   `json:"txid,omitempty"`
	Error     string  `json:"error,omitempty"`
}
