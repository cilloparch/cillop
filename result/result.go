package result

import (
	"net/http"
)

type Result struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

type DetailResult struct {
	Result
	Detail any `json:"detail"`
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DetailResult) Error() string {
	return r.Message
}

func Success(m string, c ...int) *Result {
	code := http.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func Error(m string, c ...int) *Result {
	code := http.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func SuccessDetail(m string, d any, c ...int) *DetailResult {
	code := http.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}

func ErrorDetail(m string, d any, c ...int) *DetailResult {
	code := http.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}
