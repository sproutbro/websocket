package middleware

import (
	"myonly/router"
	"net/http"
)

/*
소켓 미들웨어
현재 upgradeHandler 가 미들웨어에 있지만 나중에 이동예정
아직 웹소켓서버 경로 안정함
*/
func SocketHandler() http.HandlerFunc {
	return upgradeHandler
}

/*
주석쓰는 연습중인데 생각보다 힘드네
냠냠냠 이거 아까 메인에서도 쓴거같은대...
*/
func HttpHandler() http.HandlerFunc {
	return WithCORS(
		WithAuth(
			router.RouterHandler,
		),
	)
}
