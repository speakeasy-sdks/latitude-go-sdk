package shared

type SchemeBearer struct {
	APIKey string `security:"name=Authorization"`
}
