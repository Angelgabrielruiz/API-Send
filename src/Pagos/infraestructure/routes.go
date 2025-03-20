package infraestructure

import (
    "Send/src/Pagos/application"
    "Send/src/Pagos/infraestructure/adapters"

    "github.com/gin-gonic/gin"
    "github.com/streadway/amqp"
)

func SetupRoutesPago(r *gin.Engine, rabbitConn *amqp.Connection) {


    rabbitBroker := adapters.NewRabbitMQBroker(rabbitConn)


    ps := adapters.NewMySQLPago()


    sendPagoMessageUseCase := application.NewSendPagoMessageUseCase(rabbitBroker)

    createPagoUseCase := application.NewCreatePago(ps, sendPagoMessageUseCase)

    createPagoController := NewCreatePagoController(*createPagoUseCase)
    getPagoController := NewGetPagoController(*application.NewGetPago(ps))
    updatePagoController := NewUpdatePagoController(*application.NewUpdatePago(ps))
    deletePagoController := NewDeletePagoController(*application.NewDeletePago(ps))

    r.POST("/pagos", createPagoController.Execute)
    r.GET("/pagos", getPagoController.Execute)
    r.PUT("/pagos/:id", updatePagoController.Execute)
    r.DELETE("/pagos/:id", deletePagoController.Execute)
}
