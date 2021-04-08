package main

import (
	"main/conf"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CorsTLS())
	r.GET("/api/getHistory", middlewares.GetHistory)
	r.POST("/api/parseToken", middlewares.ParseToken)
	r.POST("/api/addHistory", middlewares.AddHistory)
	r.StaticFile("/api/boss.json", "./conf/boss.json")
	r.RunTLS(conf.Domain, "./conf/key/ssl.pem", "./conf/key/ssl.key")
}
