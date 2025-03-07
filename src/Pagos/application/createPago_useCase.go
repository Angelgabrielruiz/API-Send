package application

import "Send/src/Pagos/domain"

type CreatePago struct {
	db domain.PagoRepository
}

func NewCreatePago(db domain.PagoRepository) *CreatePago {
	return &CreatePago{db: db}
}

func (cp *CreatePago) Execute(monto int32, pago int32, cambio int32, fecha string) error {
	return cp.db.Save(monto, pago, cambio, fecha)
}

