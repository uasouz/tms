package util

type BaseResponse struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

type BaseResponseCompat struct {
	BaseResponse
	Code int `json:"code"`
}

func (response *BaseResponse) Success(data interface{}) {
	response.Status = true
	response.Data = data
}

func (response *BaseResponseCompat) Success(data interface{}) {
	response.Status = true
	response.Data = data
	response.Code = 200
}

func SuccessResponse(data interface{}) (int, *BaseResponse) {
	return 200, &BaseResponse{
		Status: true,
		Data:   data,
	}
}

func (response *BaseResponse) Failure(err interface{}, message string) {
	response.Status = false
	response.Data = map[string]interface{}{}
	response.Error = err
	response.Message = message
}

func (response *BaseResponseCompat) Failure(error interface{}, message string) {
	response.Status = false
	response.Data = map[string]interface{}{}
	response.Error = error
	response.Message = message
}

func FailureResponse(error interface{}, message string) (int, *BaseResponse) {
	return 500, &BaseResponse{
		Status:  false,
		Data:    map[string]interface{}{},
		Error:   error,
		Message: message,
	}
}
