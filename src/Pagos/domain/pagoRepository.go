//File: domain/pagoRepository.go

package domain

type PagoRepository interface {
	Save(monto int32, pago int32, cambio int32, fecha string) error
	GetAll() ([]map[string]interface{}, error)
	Update(id int, monto int32, pago int32, cambio int32, fecha string) error
	Delete(id int) error
}
