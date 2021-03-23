package middlewares

import (
	"main/conf"
	clanDao "main/dao"
	token "main/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	}
}
func checkCookies(c *gin.Context, f func(uint)) {
	//Get token from cookies
	tokenString, err := c.Cookie("sakurayume_user_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	//check token
	valid, uid, msg := token.ParseToken(tokenString, conf.SecretKey)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": msg,
		})
		return
	}
	//do func
	f(uid)
}
func GetHistory(c *gin.Context) {
	checkCookies(c, func(uid uint) {
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
	})
}
func AddHistory(c *gin.Context) {
	checkCookies(c, func(uid uint) {
		//uid, alt, round, boss, dmg, flag
		var form AddHistoryForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1,
				"error": err.Error(),
			})
			return
		}
		insertRows := clanDao.InsertHistory(form.ClanName, form.Time, uid, form.Round, form.Boss, form.Dmg, form.Flag)
		if insertRows == true {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "success",
			})
		} else {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "failed",
			})
		}
	})
}
func GetToken(c *gin.Context) {
	token, err := token.CreateToken(conf.SecretKey, string(conf.SecretKey), 123456, false)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "sakurayume_user_token",
		Value:    token,
		Path:     "/",
		MaxAge:   604800,
		Secure:   true,
		HttpOnly: false,
		SameSite: 4,
	})
	c.JSON(200, gin.H{
		"code":  0,
		"token": token,
	})
}

// func ParseToken(c *gin.Context) {
// 	tokenString := c.Query("token")
// 	valid, uid, msg := token.ParseToken(tokenString, conf.SecretKey)
// 	if valid {
// 		c.JSON(200, gin.H{
// 			"result": valid,
// 			"uid":    uid,
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"result": valid,
// 			"err":    msg,
// 		})
// 	}
// }
