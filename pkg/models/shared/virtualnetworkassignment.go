package shared

type VirtualNetworkAssignmentAttributes struct {
	Description      *string `json:"description,omitempty"`
	ServerID         *int64  `json:"server_id,omitempty"`
	Status           *string `json:"status,omitempty"`
	Vid              *int64  `json:"vid,omitempty"`
	VirtualNetworkID *int64  `json:"virtual_network_id,omitempty"`
}

type VirtualNetworkAssignment struct {
	Attributes *VirtualNetworkAssignmentAttributes `json:"attributes,omitempty"`
	ID         *string                             `json:"id,omitempty"`
}
