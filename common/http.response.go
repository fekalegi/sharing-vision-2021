package common

type BaseResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
}

func ErrorResponse(err string) BaseResponse {
	return BaseResponse{
		Message: err,
		Status:  false,
	}
}

func BadRequestResponse(errs interface{}) BaseResponse {
	return BaseResponse{
		Message: "Invalid body or query param",
		Status:  false,
		Errors:  errs,
	}
}

func SuccessResponseWithData(data interface{}, msg string) BaseResponse {
	return BaseResponse{
		Data:    data,
		Message: msg,
		Status:  true,
	}
}

func SuccessResponseNoData(msg string) BaseResponse {
	return BaseResponse{
		Message: msg,
		Status:  true,
	}
}
