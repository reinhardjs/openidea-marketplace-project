package response

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
