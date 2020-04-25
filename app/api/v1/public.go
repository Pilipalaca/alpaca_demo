package v1

import (
	"alpaca_demo/app/api"
	"github.com/gin-gonic/gin"
	"os"
)

/**
 * Ping
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func Ping(c *gin.Context) {
	runStatus := os.Getenv("RUN_STATUS")
	if runStatus != "" {
		api.OutputFull(c, runStatus, nil)
	}
	api.OutputFull(c, "pong", nil)
	return
}
