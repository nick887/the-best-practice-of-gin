package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	v1 "goFramTest/routers/v1"
	"goFramTest/services"
	"goFramTest/services/log"
	"net/http"
)

var (router *gin.Engine
	logger = log.Log)

func init() {
	router = gin.New()
	router.NoRoute(noRouteHandler())
	version1 := router.Group("/v1")
	v1.InitRoutes(version1)
	services.InitService()
}

func main() {
	logger.Info("Server Running on Port : ", 9090)
	router.Run(":9090")
}

func noRouteHandler() gin.HandlerFunc{
	return  func(c *gin.Context) {
		url:=c.Request.URL.Path
		ip:=c.Request.Host
		method:=c.Request.Method
		logger.WithFields(logrus.Fields{
			"url":url,
			"ip":ip,
			"method":method,
		}).Warn()
		c.JSON(http.StatusOK,"error url")
	}
}
