package helpers

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  Errors `json:"errors"`
}

type Errors interface{}

type Data interface{}
