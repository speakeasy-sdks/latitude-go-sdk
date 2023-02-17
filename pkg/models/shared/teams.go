package shared

type Teams struct {
	Data []Team                 `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
