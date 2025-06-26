package router

import "net/http"

// 라우터 인터페이스
type Handler interface {
	Serve(w http.ResponseWriter, r *http.Request)
}

/*
	CreateHandler는 주어진 경로 문자열에 따라 알맞은 핸들러를 생성합니다.
	내부적으로 switch 문을 통해 각 기능별 핸들러로 분기하며,
	핸들러는 공통 미들웨어(CORS, Auth 등)로 감싸집니다.
*/
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

/*
위에는 챗지피티가 써줜느대 이거 도저히 뭐라써야할지 모르것네
CreateHandler 의 응답을 라우터에게 넘기는애??
*/
func RouterHandler(w http.ResponseWriter, r *http.Request) {
	h := CreateHandler(r.URL.Path)
	h.Serve(w, r)
}
