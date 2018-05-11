package shapeshift

// Rate represents a model for rate api response
type Rate struct {
	Pair  string `json:"pair,omitempty"`
	Value string `json:"rate"`
}
