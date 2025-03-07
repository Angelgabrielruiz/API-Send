package infraestructure

import (
	"Send/src/Pagos/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreatePagoController struct {
	useCase application.CreatePago
}

func NewCreatePagoController(useCase application.CreatePago) *CreatePagoController {
	return &CreatePagoController{useCase: useCase}
}

func (cp_c *CreatePagoController) Execute(c *gin.Context) {

	var requestBody struct {
		Monto  int32  `json:"monto"`
		Pago   int32  `json:"pago"`
		Cambio int32  `json:"cambio"`
		Fecha  string `json:"fecha"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := cp_c.useCase.Execute(requestBody.Monto, requestBody.Pago, requestBody.Cambio, requestBody.Fecha); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el pago"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pago creado exitosamente"})
}
