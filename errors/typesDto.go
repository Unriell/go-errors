package errors

// ErrorDto defines an extended error
type ErrorDto struct {
	ID           string `json:"id"`
	Status       int    `json:"status"`
	Msg          string `json:"msg"`
	ComponentMsg string `json:"componentMsg"`
	SentryCode   string `json:"sentryCode"`
}
