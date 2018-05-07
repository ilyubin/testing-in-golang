package model

type PostResponse struct {
	Args    interface{}       `json:"args"`
	Data    string            `json:"data"`
	Files   interface{}       `json:"files"`
	Form    interface{}       `json:"form"`
	Headers HeadersInResponse `json:"headers"`
	Json    PostRequest       `json:"json"`
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
}
