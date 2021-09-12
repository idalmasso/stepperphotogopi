package server

type errorMessage struct {
	Message string `json:"message"`
}
type valueResponse struct {
	Value string `json:"value"`
}

type valueListStringResponse struct {
	Value []string `json:"value"`
}
