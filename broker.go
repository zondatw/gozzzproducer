package goZzzProducer

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
}

// TaskJsonType is register task json schema
type TaskJsonType struct {
	Task string      `json:"task"`
	Args interface{} `json:"args"`
}

// NewBroker will initialize a new broker
func NewBroker(address string, password string, db int) *Broker {
	return &Broker{
		conn:     NewRedisConn(address, password, db),
		queueKey: "goZzzWorker:task:queue",
		msgKey:   "goZzzWorker:task:msg",
	}
}

// AddTask add new task to broker
func (broker *Broker) AddTask(taskName string, jsonData interface{}, delaySec int) (retErr error) {
	byteArrayData, err := json.Marshal(&TaskJsonType{
		Task: taskName,
		Args: jsonData,
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
	taskID := taskUUID.String()
	timing := time.Now().Add(time.Second * time.Duration(delaySec)).Unix()
	broker.conn.SetHashValue(broker.msgKey, taskID, stringData)
	broker.conn.SetZSetValue(broker.queueKey, taskID, float64(timing))
	return
}
