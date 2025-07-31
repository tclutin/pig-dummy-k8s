package internal

import "time"

type API struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type apiHandler struct {
	api *API
}

type postgresSamples struct {
	ID      int
	Message string
}

type rabbitMQSamples struct {
	ID        int64
	Timestamp time.Time
	Message   string
}

type kafkaSamples struct {
	ID        int64
	Timestamp time.Time
	Message   string
}
