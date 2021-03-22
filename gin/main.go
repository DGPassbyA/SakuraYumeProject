package main

import (
	"log"
	"main/clanDao"
	"main/middlewares"
	token "main/util"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/getHistory", func(c *gin.Context) {
		clanName := c.Query("clanname")
		time := c.Query("time")
		var history = clanDao.GetHistoryByTime(clanName, time)
		if history == nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "error",
			})
		} else {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "success",
				"data":    history,
			})
		}
	})
	r.GET("/getToken", func(c *gin.Context) {
		SecretKey := []byte("TestKey")
		token, err := token.CreateToken(SecretKey, "SakuraYume", 123456, false)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"token": token,
		})
	})
	r.GET("/parseToken", func(c *gin.Context) {
		SecretKey := []byte("TestKey")
		tokenString := c.Query("token")
		valid, err := token.ParseToken(tokenString, SecretKey)
		c.JSON(200, gin.H{
			"result": valid,
			"err":    err,
		})
	})
	r.Run("127.0.0.1:8081")
}
