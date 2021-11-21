package http

const (
	Scheme        = "http"
	Host          = "localhost"
	Port          = "8000"
	TimeoutSecond = 5
)

type Config struct {
	Scheme      string
	Host        string
	Port        string
	URL         string
	AccessToken string
}
