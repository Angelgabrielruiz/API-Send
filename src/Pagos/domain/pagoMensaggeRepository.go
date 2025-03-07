package domain

type PagoMessageBroker interface {
	Publish(event string, data interface{}) error
}
