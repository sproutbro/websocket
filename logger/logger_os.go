package logger

import "os"

type FileLogger struct {
	file *os.File
}

func NewFileLogger(path string) (*FileLogger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: f}, nil
}

func (f *FileLogger) Info(msg string) {
	f.file.WriteString("[INFO] " + msg + "\n")
}

func (f *FileLogger) Error(err error) {
	f.file.WriteString("[ERROR] " + err.Error() + "\n")
}
