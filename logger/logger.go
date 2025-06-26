package logger

import (
	"log"
	"os"
	"sync"
)

// Logger 구조체: 원하는 방식대로 확장 가능
type Logger struct {
	logger *log.Logger
}

var (
	instance *Logger   // 싱글톤 인스턴스
	once     sync.Once // 단 한 번만 실행 보장
)

// GetLogger: 싱글톤 객체 반환
func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			logger: log.New(os.Stdout, "[LOG] ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
	return instance
}

// Info 로그: 일반 정보 출력
func (l *Logger) Info(msg string) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Println(msg)
}

// Error 로그: 에러 메시지 출력
func (l *Logger) Error(msg string) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Println(msg)
}
