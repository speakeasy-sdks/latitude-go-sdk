package shared

type ErrorObjectErrors struct {
	Code   *string                `json:"code,omitempty"`
	Detail *string                `json:"detail,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Status *string                `json:"status,omitempty"`
	Title  *string                `json:"title,omitempty"`
}

type ErrorObject struct {
	Errors []ErrorObjectErrors `json:"errors,omitempty"`
}
