package shapeshift

// DepositLimit represents a limit for deposit on specified currency pair
type DepositLimit struct {
	Pair  string `json:"pair,omitempty"`
	Limit string `json:"limit"`
}
