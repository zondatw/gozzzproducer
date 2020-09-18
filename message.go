package gozzzproducer

import (
	"encoding/json"
	"time"
)

// Message is message struct
type Message struct {
	taskID string
	broker *Broker
}

// taskMsgType is return message json schema
type retMsgType struct {
	Status string `json:"status"` // complete execution
	Msg    string `json:"msg"`    // return message json type
}

// NewMessage will initialize a new message
func NewMessage(taskID string, broker *Broker) *Message {
	return &Message{
		taskID: taskID,
		broker: broker,
	}
}

// GetRetMessage get task return message
func (m *Message) GetRetMessage(block bool, timeoutSec int) (status string, msg string, err error) {
	var maxIndex int = 1
	var value string
	if block {
		maxIndex = timeoutSec
	}
	for i := 0; i < maxIndex; i++ {
		value, err = m.broker.conn.GetHashValue(m.broker.retKey, m.taskID)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
	if err == nil {
		var retMsg retMsgType
		json.Unmarshal([]byte(value), &retMsg)
		status = retMsg.Status
		msg = retMsg.Msg
	}
	return
}
