# gozzzproducer

gozzzproducer is task producer with [gozzzWorker](http://github.com/zondatw/gozzzWorker).  

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

	p.AddTask("Task 1", &ArgType{A: 1, B: "producer~~~~"}, 10)
}
```