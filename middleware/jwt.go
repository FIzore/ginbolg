package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour) //设置过期时间
	SetClaims := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims) //创建一个新的token对象
	token, err := reqClaim.SignedString(JwtKey)                      //使用密钥进行签名
	if err != nil {
		return "", errmsg.ERROR_TOKEN_WRONG
	}
	return token, errmsg.SUCCESS
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := settoken.Claims.(*MyClaims); code && settoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR_TOKEN_WRONG
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
		}
		checkToken := strings.SplitN(tokenHeader, "", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.Abort()
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
		}
		c.JSON(200, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
