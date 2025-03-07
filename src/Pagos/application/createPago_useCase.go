package application

import "Send/src/Pagos/domain"


type CreatePago struct {
	db             domain.PagoRepository
	rabbitRepo     RabbitRepository
}

func NewCreatePago(db domain.PagoRepository, rabbitRepo RabbitRepository) *CreatePago {
	return &CreatePago{db: db, rabbitRepo: rabbitRepo}
}

func (cp *CreatePago) Execute(monto int32, pago int32, cambio int32, fecha string) error {
	// Guardar el pago en la base de datos
	err := cp.db.Save(monto, pago, cambio, fecha)
	if err != nil {
		return err
	}

	// Crear el mensaje para enviar a RabbitMQ
	message := "Pago creado: Monto " + string(monto) + ", Pago " + string(pago) + ", Cambio " + string(cambio) + ", Fecha " + fecha

	// Enviar el mensaje a la cola de RabbitMQ
	err = cp.rabbitRepo.SendMessage("pago_created_queue", message)
	if err != nil {
		return err
	}

	return nil
}
