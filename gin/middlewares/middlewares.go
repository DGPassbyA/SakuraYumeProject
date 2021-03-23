package middlewares

import (
	"log"
	"main/conf"
	clanDao "main/dao"
	token "main/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
func GetHistory(c *gin.Context) {
	var form GetHistoryPayload
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	var history = clanDao.GetHistoryByTime(form.ClanName, form.Time)
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
}
func AddHistory(c *gin.Context) {
	//uid, alt, round, boss, dmg, flag
	var form AddHistoryForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	//check the token and get uid
	//....
	//remember to change `required` of uid in the struct
	c.JSON(200, gin.H{
		"result": clanDao.InsertHistory(form.ClanName, form.Time, form.Uid, form.Round, form.Boss, form.Dmg, form.Flag),
	})
}
func GetToken(c *gin.Context) {
	token, err := token.CreateToken(conf.SecretKey, string(conf.SecretKey), 123456, false)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}
func ParseToken(c *gin.Context) {
	tokenString := c.Query("token")
	valid, msg := token.ParseToken(tokenString, conf.SecretKey)
	if valid {
		c.JSON(200, gin.H{
			"result": valid,
			"uid":    msg,
		})
	} else {
		c.JSON(200, gin.H{
			"result": valid,
			"err":    msg,
		})
	}
}
