package entities

type Pago struct {
	id     int32
	monto  int32
	pago   int32
	cambio int32
	fecha  string
}

func NewTicket(monto int32, pago int32, cambio int32, fecha string) *Pago {
	return &Pago{ monto:  monto,pago:   pago, cambio: cambio, fecha:  fecha,}
}
