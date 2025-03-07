package application

import "Send/src/Pagos/domain"

type DeletePago struct {
	db domain.PagoRepository
}

func NewDeletePago(db domain.PagoRepository) *DeletePago {
	return &DeletePago{db: db}
}

func (dp *DeletePago) Execute(id int)  {
	dp.db.Delete(id)
}
