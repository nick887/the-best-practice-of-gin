package v1

import (
	"github.com/gin-gonic/gin"
	"goFramTest/controllers/v1/biz/hello"
)

func SetHelloRoutes(g *gin.RouterGroup)  {
	g.GET("/hello", hello.Hello)
}
