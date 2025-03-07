package application

import "Send/src/Pagos/domain"

type GetPago struct {
	db domain.PagoRepository
}

func NewGetPago(db domain.PagoRepository) *GetPago {
	return &GetPago{db: db}
}

func (gp *GetPago) Execute() ([]map[string]interface{}, error) {
	return gp.db.GetAll()
}
