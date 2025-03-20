package infraestructure

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "Send/src/Pagos/application"
)

type CreatePagoController struct {
    createPagoUseCase application.CreatePago
}

func NewCreatePagoController(useCase application.CreatePago) *CreatePagoController {
    return &CreatePagoController{
        createPagoUseCase: useCase,
    }
}

func (c *CreatePagoController) Execute(ctx *gin.Context) {
    var req struct {
        Monto  int32  `json:"monto"`
        Pago   int32  `json:"pago"`
        Cambio int32  `json:"cambio"`
        Fecha  string `json:"fecha"`
    }

    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := c.createPagoUseCase.Execute(req.Monto, req.Pago, req.Cambio, req.Fecha)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Pago creado exitosamente"})
}
