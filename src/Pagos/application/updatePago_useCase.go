package application

import "Send/src/Pagos/domain"

type UpdatePago struct {
	db domain.PagoRepository
}

func NewUpdatePago(db domain.PagoRepository) *UpdatePago {
	return &UpdatePago{db: db}
}

func (up *UpdatePago) Execute(id int, monto int32, pago int32, cambio int32, fecha string) error {
	return up.db.Update(id, monto, pago, cambio, fecha)
}
