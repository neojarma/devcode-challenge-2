package response

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type emptyDataReponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

func NewResponse(status string, message string, data any) *response {
	return &response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewEmptyDataResponse(status string, message string) *emptyDataReponse {
	return &emptyDataReponse{
		Status:  status,
		Message: message,
	}
}
