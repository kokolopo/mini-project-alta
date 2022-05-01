package helper

type Response struct {
	MetaMessage Meta        `json:"meta"`
	Data        interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	metaData := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		MetaMessage: metaData,
		Data:        data,
	}

	return jsonResponse
}
