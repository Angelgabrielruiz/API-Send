package infraestructure

import (
	"Send/src/Pagos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePagoController struct {
	useCase application.UpdatePago
}

func NewUpdatePagoController(useCase application.UpdatePago) *UpdatePagoController {
	return &UpdatePagoController{useCase: useCase}
}

func (up_c *UpdatePagoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var input struct {
		Monto  int32  `json:"monto"`
		Pago   int32  `json:"pago"`
		Cambio int32  `json:"cambio"`
		Fecha  string `json:"fecha"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := up_c.useCase.Execute(id, input.Monto, input.Pago, input.Cambio, input.Fecha); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}
