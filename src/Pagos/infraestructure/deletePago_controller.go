package infraestructure

import (
	"Send/src/Pagos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePagoController struct {
	useCase application.DeletePago
}

func NewDeletePagoController(useCase application.DeletePago) *DeletePagoController {
	return &DeletePagoController{useCase: useCase}
}

func (ds_c *DeletePagoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pago ID"})
		return
	}

	ds_c.useCase.Execute(id)
	c.JSON(http.StatusOK, gin.H{"message": "Pago deleted successfully"})
}
