package handler

import (
	"fmt"
	"net/http"
	"winartodev/coba-graphql/middleware"
	"winartodev/coba-graphql/response"

	"github.com/julienschmidt/httprouter"
)

type Register interface {
	Register(r *httprouter.Router) error
}

func Decorate(handle middleware.HandleWithError) httprouter.Handle {
	return middleware.HTTP(handle)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	response.HttpResponseError(w, r, http.StatusNotFound, "endpoint not found")
}

func NewHandler(register ...Register) http.Handler {
	router := httprouter.New()
	router.HandleMethodNotAllowed = false

	router.GET("/healthz", healthz)

	for _, r := range register {
		r.Register(router)
	}

	router.NotFound = http.HandlerFunc(NotFound)
	return router
}

func healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "OK")
}
