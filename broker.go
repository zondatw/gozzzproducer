package gozzzproducer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Broker is broker struct
type Broker struct {
	conn     *RedisConn
	queueKey string
	msgKey   string
	retKey   string
}

// TaskJSONType is register task json schema
type TaskJSONType struct {
	Task     string      `json:"task"`
	Args     interface{} `json:"args"`
	Priority int         `json:"priority"`
}

// NewBroker will initialize a new broker
func NewBroker(address string, password string, db int) *Broker {
	return &Broker{
		conn:     NewRedisConn(address, password, db),
		queueKey: "gozzzworker:task:queue",
		msgKey:   "gozzzworker:task:msg",
		retKey:   "gozzzworker:task:ret",
	}
}

// AddTask add new task to broker
func (broker *Broker) AddTask(taskName string, jsonData interface{}, priority int, delaySec int) (taskID string, retErr error) {
	byteArrayData, err := json.Marshal(&TaskJSONType{
		Task:     taskName,
		Args:     jsonData,
		Priority: priority,
	})
	if err != nil {
		retErr = err
		return
	}
	stringData := string(byteArrayData)
	taskUUID, err := uuid.NewRandom()
	if err != nil {
		retErr = err
		return
	}
	taskID = taskUUID.String()
	timing := time.Now().Add(time.Second * time.Duration(delaySec)).Unix()
	broker.conn.SetHashValue(broker.msgKey, taskID, stringData)
	broker.conn.SetZSetValue(broker.queueKey, taskID, float64(timing))
	return
}
