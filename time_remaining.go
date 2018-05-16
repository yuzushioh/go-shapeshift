package shapeshift

// TimeRemaining represents how many seconds are left before the transaction expires
type TimeRemaining struct {
	// The status can be either "pending" or "expired".
	Status string `json:"status"`

	// If the status is expired then seconds_remaining will show 0
	SecondsRemaining int64 `json:"seconds_remaining"`
}
