package application

import (
	"fmt"
	"Send/src/Pagos/domain"

)

type CreatePago struct {
	db            domain.PagoRepository    
	messageBroker domain.PagoMessageBroker
}

func NewCreatePago(db domain.PagoRepository, messageBroker domain.PagoMessageBroker) *CreatePago {
	return &CreatePago{db: db, messageBroker: messageBroker}
}

func (cp *CreatePago) Execute(monto int32, pago int32, cambio int32, fecha string) error {

	err := cp.db.Save(monto, pago, cambio, fecha)
	if err != nil {
		return err
	}


	message := map[string]interface{}{
		"monto":  monto,
		"pago":   pago,
		"cambio": cambio,
		"fecha":  fecha,
	}


	err = cp.messageBroker.Publish("pago_created_queue", message)
	if err != nil {
		return err
	}

	fmt.Println("Pago creado y mensaje publicado en RabbitMQ")
	return nil
}
