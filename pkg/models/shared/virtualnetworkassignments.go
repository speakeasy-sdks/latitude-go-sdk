package shared

type VirtualNetworkAssignments struct {
	Data []VirtualNetworkAssignment `json:"data,omitempty"`
	Meta map[string]interface{}     `json:"meta,omitempty"`
}
