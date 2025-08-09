package gitlib

type WebHookRequest struct {
	Owner  string   `json:"owner"`
	Repo   string   `json:"repo"`
	Name   string   `json:"name"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Config Config   `json:"config"`
}

type Config struct {
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
}
