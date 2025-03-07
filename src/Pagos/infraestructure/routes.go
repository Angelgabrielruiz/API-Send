package infraestructure

import (
	"Send/src/Pagos/application"
	"Send/src/Pagos/infraestructure/adapters"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	
)

func SetupRoutesPago(r *gin.Engine, rabbitConn *amqp.Connection) {
	// Crear el repositorio RabbitMQ
	rabbitRepo := application.NewRabbitRepository(rabbitConn)

	// Crear el repositorio de MySQL
	ps := adapters.NewMySQLPago()

	// Crear los casos de uso
	createPagoController := NewCreatePagoController(*application.NewCreatePago(ps, rabbitRepo))
	getPagoController := NewGetPagoController(*application.NewGetPago(ps))
	updatePagoController := NewUpdatePagoController(*application.NewUpdatePago(ps))
	deletePagoController := NewDeletePagoController(*application.NewDeletePago(ps))

	// Configurar las rutas
	r.POST("/pagos", createPagoController.Execute)
	r.GET("/pagos", getPagoController.Execute)
	r.PUT("/pagos/:id", updatePagoController.Execute)
	r.DELETE("/pagos/:id", deletePagoController.Execute)
}
