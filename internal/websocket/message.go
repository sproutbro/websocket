package websocket

import (
	"encoding/json"
	"errors"
)

type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Body string `json:"body"`
}

func ParseMessage(data []byte) (*Message, error) {
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}

	if msg.Type == "" || msg.User == "" || msg.Body == "" {
		return nil, errors.New("메시지 필수 항목 누락")
	}

	return &msg, nil
}
