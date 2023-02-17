package shared

type LazySideloadMeta struct {
	Included *bool `json:"included,omitempty"`
}

type LazySideload struct {
	Meta *LazySideloadMeta `json:"meta,omitempty"`
}
