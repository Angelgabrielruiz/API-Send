package infraestructure

import (
    "github.com/gin-gonic/gin"
    "Send/src/Pagos/application"
    "Send/src/Pagos/infraestructure/adapters" 
)

func SetupRoutesPago(r *gin.Engine) {
    ps := adapters.NewMySQLPago() 

    createPagoController := NewCreatePagoController(*application.NewCreatePago(ps))
    getPagoController := NewGetPagoController(*application.NewGetPago(ps))
    updatePagoController := NewUpdatePagoController(*application.NewUpdatePago(ps))
    deletePagoController := NewDeletePagoController(*application.NewDeletePago(ps))

    r.POST("/pagos", createPagoController.Execute)
    r.GET("/pagos", getPagoController.Execute)
    r.PUT("/pagos/:id", updatePagoController.Execute)
    r.DELETE("/pagos/:id", deletePagoController.Execute)
}