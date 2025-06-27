package logger

import "log"

type JsonLogger struct{}

func (j *JsonLogger) Info(msg string) {
	log.Printf(`{"level":"info","message":"%s"}`, msg)
}

func (j *JsonLogger) Error(err error) {
	log.Printf(`{"level":"error","message":"%s"}`, err.Error())
}
