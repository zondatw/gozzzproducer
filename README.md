# gozzzproducer

[![Go Report Card](https://goreportcard.com/badge/github.com/zondatw/gozzzproducer)](https://goreportcard.com/report/github.com/zondatw/gozzzproducer)

gozzzproducer is task producer with [gozzzworker](http://github.com/zondatw/gozzzworker).  

## Installation

To install  
`go get github.com/zondatw/gozzzproducer`  

To import  
`import "github.com/zondatw/gozzzproducer"`  

## Quickstart

Push new task:  
```go
producerObj.AddTask("Task Name", interface{}, delay sec)
```

ProducerSetting:  
```go
&gozzzproducer.ProducerSetting{
    Address:  "localhost:6379",     // Redis path
    Password: "",                   // Redis password, set empty string if no password
    DB:       0,                    // Redis DB number
}
```

Example quicker start:  
```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/zondatw/gozzzproducer"
)

func main() {
	type ArgType struct {
		A int    `json:"a"`
		B string `json:"b"`
	}

	p := gozzzproducer.NewProducer(&gozzzproducer.ProducerSetting{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	})

	m, err := p.AddTask("Task 1", &ArgType{A: 1, B: "producer~~~~"}, 10)
	if err == nil {
		// block: true, timeout: 60sec
		status, jsonMsg, err := m.GetRetMessage(true, 60)
		// When status is Success, task is complete execution and msg is task function return message
		// When status is Fail, task have error and msg is task error message
		log.Println("status:", status, "msg:", msg, "err:", err)
	}
}
```
