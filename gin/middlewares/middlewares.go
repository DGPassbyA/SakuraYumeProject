package middlewares

import (
	"fmt"
	"main/conf"
	clanDao "main/dao"
	token "main/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func CorsTLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			return
		}
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
		c.JSON(200, gin.H{
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
			c.JSON(200, gin.H{
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

func ParseToken(c *gin.Context) {
	tokenString := c.PostForm("token")
	fmt.Println(tokenString)
	valid, _, msg := token.ParseToken(tokenString, conf.SecretKey)
	if valid {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "sakurayume_user_token",
			Value:    tokenString,
			Path:     "/",
			MaxAge:   604800,
			Secure:   true,
			HttpOnly: false,
			SameSite: 4,
		})
		c.JSON(200, gin.H{
			"code":   0,
			"result": valid,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 1,
			"err":  msg,
		})
	}
}
