package middlewares

import (
	"TSACodingChallengeAPI/src/common"
	"errors"
	"net/http"
	"runtime"
)

type ErrorMiddleware struct {
	StackSize int
}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{1024 * 4}
}

func (emw *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			stack := make([]byte, emw.StackSize)
			stack = stack[:runtime.Stack(stack, false)]
			http.Error(w, common.ErrorResponse{errors.New("Unknow server error")}, http.StatusInternalServerError)
		}
	}()

	next(w, r)
}
