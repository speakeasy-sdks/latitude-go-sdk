package shared

type VirtualNetworks struct {
	Data []VirtualNetwork       `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
