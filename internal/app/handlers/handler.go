package handlers

import (
	"gapp1/internal/app/services"
	"gapp1/pkg/response"

	"github.com/gin-gonic/gin"
)

type Ping struct {
	pingSerivce *services.PingService
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (p *Ping) PingHandler(c *gin.Context) {
	m, _ := p.pingSerivce.Ping()
	response.Success(c, 0, m, nil)
}
