package gozzzproducer

// Producer is producer struct
type Producer struct {
	broker *Broker
}

// ProducerSetting is producer setting
type ProducerSetting struct {
	Address  string
	Password string
	DB       int
}

// NewProducer will initialize a new producer
func NewProducer(setting *ProducerSetting) *Producer {
	broker := NewBroker(setting.Address, setting.Password, setting.DB)
	return &Producer{
		broker: broker,
	}
}

// AddTask add new task to broker
func (p *Producer) AddTask(taskName string, jsonData interface{}, delaySec int) (retErr error) {
	if err := p.broker.AddTask(taskName, jsonData, delaySec); err != nil {
		retErr = err
	}
	return
}
