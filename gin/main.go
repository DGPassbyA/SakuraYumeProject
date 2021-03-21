package main

import (
	"main/clanDao"
	"main/middlewares"

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
	r.Run("127.0.0.1:8081")
}
