package error

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goFramTest/services/log"
	"net/http"
)

var (
	logger = log.Log
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err!=nil {
				url:=c.Request.URL.Path
				method := c.Request.Method
				logger.WithFields(logrus.Fields{
					"url":url,
					"method":method,
					"error":err,
				}).Error()
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
