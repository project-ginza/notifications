package request

import (
	"github.com/gin-gonic/gin"
	logging "github.com/project-ginza/notifications/log"
	"net/http"
)

var code string = "code"
var message string = "message"
var loggingCh = logging.LoggingCh

func Home(c *gin.Context) {
	loggingCh.PushDebugMessageToChannel("home")
	c.JSON(http.StatusOK, gin.H{
		code: http.StatusOK,
	})
}
