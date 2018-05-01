package model

type GetResponse struct {
	Args    interface{}       `json:"args"`
	Headers HeadersInResponse `json:"headers"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}

type HeadersInResponse struct {
	AcceptEncoding string `json:"Accept-Encoding"`
	Connection     string `json:"Connection"`
	ContentType    string `json:"Content-Type"`
	Host           string `json:"Host"`
	UserAgent      string `json:"User-Agent"`
}
