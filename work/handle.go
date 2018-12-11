package work

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	log.Info("222222222222222")
	c.JSON(http.StatusOK, gin.H{
		"code": 111,
		"msg":  "hello word",
		"data": 222,
	})
}
