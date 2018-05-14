package shapeshift

// MarketInfo represnets market information
type MarketInfo struct {
	Pair     string  `json:"pair"`
	Rate     float64 `json:"rate"`
	Limit    float64 `json:"limit"`
	Min      float64 `json:"min"`
	MinerFee float64 `json:"minerFee"`
}
