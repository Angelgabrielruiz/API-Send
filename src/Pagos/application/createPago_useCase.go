package application

import (
	"fmt"
	"Send/src/Pagos/domain"
	"strconv"
)

type CreatePago struct {
	db         domain.PagoRepository
	rabbitRepo RabbitRepository
}

func NewCreatePago(db domain.PagoRepository, rabbitRepo RabbitRepository) *CreatePago {
	return &CreatePago{db: db, rabbitRepo: rabbitRepo}
}

func (cp *CreatePago) Execute(monto int32, pago int32, cambio int32, fecha string) error {
	
	err := cp.db.Save(monto, pago, cambio, fecha)
	if err != nil {
		return err
	}

	
	message := fmt.Sprintf("Pago creado: Monto %s, Pago %s, Cambio %s, Fecha %s", 
		strconv.Itoa(int(monto)), strconv.Itoa(int(pago)), 
		strconv.Itoa(int(cambio)), fecha)

	
	err = cp.rabbitRepo.SendMessage("pago_created_queue", message)
	if err != nil {
		return err
	}

	return nil
}
