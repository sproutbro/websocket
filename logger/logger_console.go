package logger

import "fmt"

type ConsoleLogger struct {
	Logged []string
}

func (m *ConsoleLogger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (m *ConsoleLogger) Error(err error) {
	fmt.Println("[ERROR]", err.Error())
}
