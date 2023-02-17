package shared

type Servers struct {
	Data []ServerData           `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
