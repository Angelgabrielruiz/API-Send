package application

import (
    "fmt"
    "Send/src/Pagos/domain"
)

type SendPagoMessageUseCase struct {
    messageBroker domain.PagoMessageBroker
}

func NewSendPagoMessageUseCase(messageBroker domain.PagoMessageBroker) *SendPagoMessageUseCase {
    return &SendPagoMessageUseCase{
        messageBroker: messageBroker,
    }
}


func (uc *SendPagoMessageUseCase) Execute(event string, data interface{}) error {
    err := uc.messageBroker.Publish(event, data)
    if err != nil {
        return err
    }

    fmt.Printf("Mensaje publicado en la cola '%s': %+v\n", event, data)
    return nil
}
