package http_health

import (
	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary get status of the server
// @Description returns ok if the server is up
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /_monitoring/health [get]
func makeGetHealthEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "ok")
	}
}

type Handlers struct{}

func (he Handlers) Register(engine *gin.Engine) {
	engine.GET("/_monitoring/health", makeGetHealthEndpoint())
}
