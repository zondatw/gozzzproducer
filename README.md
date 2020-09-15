# goZzzProducer

goZzzProducer is task producer with [goZzzWorker](http://github.com/zondatw/goZzzWorker).  

## Installation

To install  
`go get github.com/zondatw/goZzzProducer`  

To import  
`import "github.com/zondatw/goZzzProducer"`  

## Quickstart

Push new task:  
```go
producerObj.AddTask("Task Name", interface{}, delay sec)
```

ProducerSetting:  
```go
&goZzzProducer.ProducerSetting{
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

	"github.com/zondatw/goZzzProducer"
)

func main() {
	type ArgType struct {
		A int    `json:"a"`
		B string `json:"b"`
	}

	p := goZzzProducer.NewProducer(&goZzzProducer.ProducerSetting{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
	})

	p.AddTask("Task 1", &ArgType{A: 1, B: "producer~~~~"}, 10)
}
```