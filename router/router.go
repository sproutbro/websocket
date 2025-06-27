package router

import (
	"net/http"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {

	h := CreateHandler(r.URL.Path)
	h.Serve(w, r)
}

func CreateHandler(path string) Handler {
	switch path {
	case "/users":
		return &UserHandler{}
	case "/posts":
		return &PostHandler{}
	default:
		return &NotFoundHandler{}
	}
}

type Handler interface {
	Serve(w http.ResponseWriter, r *http.Request)
}

type PostHandler struct{}

func (h *PostHandler) Serve(w http.ResponseWriter, r *http.Request) {
}

type UserHandler struct{}

func (h *UserHandler) Serve(w http.ResponseWriter, r *http.Request) {
}

type NotFoundHandler struct{}

func (h *NotFoundHandler) Serve(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
