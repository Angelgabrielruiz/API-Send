package infraestructure

import (
	"Send/src/Pagos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPagoController struct {
	useCase application.GetPago
}

func NewGetPagoController(useCase application.GetPago) *GetPagoController {
	return &GetPagoController{useCase: useCase}
}

func (gp_c *GetPagoController) Execute(c *gin.Context) {
	pagos, err := gp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los pagos"})
		return
	}

	c.JSON(http.StatusOK, pagos)
}
