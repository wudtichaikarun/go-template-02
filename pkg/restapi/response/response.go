package restapi_response

import (
	"net/http"

	restapi_context "github.com/wudtichaikarun/poc-go-template-02/pkg/restapi/context"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	// Meta    *paginate.Meta `json:"meta,omitempty"`
}

type FailResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (f *FailResponse) Error() string {
	return f.Message
}

func NewFail(message string, code int) error {
	return &FailResponse{
		Code:    code,
		Status:  FailStatus,
		Message: message,
		Errors:  nil,
	}
}

func NotContent(ctx restapi_context.Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusNoContent,
	})
}

func Created(ctx restapi_context.Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusCreated,
	})
}

func Data(ctx restapi_context.Context, data interface{}) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func Success(ctx restapi_context.Context, r *SuccessResponse) {
	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" {
		r.Message = ""
	}

	ctx.JSON(r.Code, r)
}

func UnprocessableEntity(ctx restapi_context.Context, errors interface{}) {
	Fail(ctx, &FailResponse{
		Code:    http.StatusUnprocessableEntity,
		Status:  FailStatus,
		Message: "",
		Errors:  errors,
	})
}

func BadRequest(ctx restapi_context.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadRequest))
}

func Forbidden(ctx restapi_context.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusForbidden))
}

func NotFound(ctx restapi_context.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusNotFound))
}

func InternalServerError(ctx restapi_context.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusInternalServerError))
}

func BadGateway(ctx restapi_context.Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadGateway))
}

func Fail(ctx restapi_context.Context, err error) {
	var response *FailResponse

	if e, ok := err.(*FailResponse); err != nil && ok {
		response = e
	} else {
		var message string
		if err == nil {
			message = "Server error occurred"
		} else {
			message = err.Error()
		}

		response = &FailResponse{
			Code:    500,
			Status:  FailStatus,
			Message: message,
		}
	}

	ctx.JSON(response.Code, response)
}
