package main

import (
	"main/conf"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/getHistory", middlewares.GetHistory)
	r.GET("/getToken", middlewares.GetToken)
	r.POST("/parseToken", middlewares.ParseToken)
	r.POST("/addHistory", middlewares.AddHistory)
	r.Run(conf.Domain)
}
