package ws

// Message 는 클라이언트가 서버에 보내는 JSON 메시지 형식입니다.
// to: 받을 사람 ID, msg: 보낼 실제 메시지
type Message struct {
	To  string `json:"to"`
	Msg string `json:"msg"`
}
