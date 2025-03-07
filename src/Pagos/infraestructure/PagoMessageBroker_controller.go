package infraestructure

import "log"

type SimpleMessageBroker struct{}

func NewSimpleMessageBroker() *SimpleMessageBroker {
    return &SimpleMessageBroker{}
}

func (s *SimpleMessageBroker) Publish(event string, data interface{}) error {
    log.Printf("Publishing event: %s with data: %v", event, data)
    return nil
}