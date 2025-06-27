package httpserver

import (
	"fmt"
	"myonly/middleware"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Router!")
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	loggedHandler := middleware.WithLogging(&HelloHandler{})
	mux.Handle("/", loggedHandler)

	return mux
}
