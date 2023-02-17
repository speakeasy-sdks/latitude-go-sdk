package shared

type DeployConfigDataAttributes struct {
	Hostname        *string       `json:"hostname,omitempty"`
	OperatingSystem *string       `json:"operating_system,omitempty"`
	Raid            *string       `json:"raid,omitempty"`
	SSHKeys         []interface{} `json:"ssh_keys,omitempty"`
}

type DeployConfigData struct {
	Attributes *DeployConfigDataAttributes `json:"attributes,omitempty"`
	ID         *string                     `json:"id,omitempty"`
}

type DeployConfig struct {
	Data *DeployConfigData `json:"data,omitempty"`
}
