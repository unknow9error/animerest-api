package models

type HttpSuccess struct {
	Message    string      `json:"message"`
	Status     interface{} `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type HttpError struct {
	Reason     string      `json:"reason"`
	Status     interface{} `json:"status"`
	StatusCode int         `json:"statusCode"`
}
