package v1

import (
	"github.com/gin-gonic/gin"
	error2 "goFramTest/middlewares/error"
	"goFramTest/middlewares/logger"
)

func InitRoutes(g *gin.RouterGroup)  {
	g.Use(error2.CatchError())
	g.Use(logger.MiddlewareLogger())
	SetUserRoutes(g)
	SetHelloRoutes(g)
}
