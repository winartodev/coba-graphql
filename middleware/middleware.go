package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandleWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

type Decorator func(handle HandleWithError) HandleWithError

func ApplyDecorators(handle HandleWithError, ds ...Decorator) HandleWithError {
	for _, d := range ds {
		handle = d(handle)
	}

	return handle
}

func HTTP(handle HandleWithError) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r, params)
	}
}
