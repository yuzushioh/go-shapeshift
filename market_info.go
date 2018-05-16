package shapeshift

// MarketInfo represnets market information
type MarketInfo struct {
	Pair     string  `json:"pair,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Limit    float64 `json:"limit,omitempty"`
	Min      float64 `json:"min,omitempty"`
	MinerFee float64 `json:"minerFee,omitempty"`
	Error    string  `json:"error,omitempty"`
}
