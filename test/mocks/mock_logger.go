package mocks

import "errors"

type MockLogger struct {
	Logged []string
}

func (m *MockLogger) Info(msg string) {
	m.Logged = append(m.Logged, "INFO: "+msg)
}

func (m *MockLogger) Error(err error) {
	if err == nil {
		err = errors.New("unknown error")
	}
	m.Logged = append(m.Logged, "ERROR: "+err.Error())
}
