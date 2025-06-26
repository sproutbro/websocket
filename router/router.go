package router

import (
	"fmt"
	"net/http"
)

/*
나중에 다 각자 뺄꺼임
언젠가 언젠가
이것은 패토리 패턴
그리고 팩토리 패턴만이 아니라 go 언어의 행동지향 인터페이스던가??
*/
type PostHandler struct{}

func (h *PostHandler) Serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "게시글 목록입니다.")
}

// /users 핸들러
type UserHandler struct{}

func (h *UserHandler) Serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "유저 목록입니다.")
}

// 기본 NotFound 핸들러
type NotFoundHandler struct{}

func (h *NotFoundHandler) Serve(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
