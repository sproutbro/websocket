// File: logger.go
// Author: bbcode
// Date: 2025-06-27
//
// Description:
//   - 개발로그: ConsoleLogger
//   - 서버로그: FileLogger
//   - JSON로그: JsonLogger
//   - 테스트로그: MockLogger
//
// Related files:
//   - package/logger
//
// Test Strategy:
//   - 아직안했다 환경변수 적용도 남았다
//   - 오늘 머리털이 빠지는거 같냐..
//
// Last Modified: 2025-06-28 by JackieChan
package logger

type Logger interface {
	Info(msg string)
	Error(err error)
}

func NewLogger(env string) Logger {
	switch env {
	case "test":
		return &MockLogger{}
	case "dev":
		return &ConsoleLogger{}
	case "prod":
		logger, _ := NewFileLogger("./app.log")
		return logger
	case "cloud":
		return &JsonLogger{}
	default:
		return &ConsoleLogger{}
	}
}
