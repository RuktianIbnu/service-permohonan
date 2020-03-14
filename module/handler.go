package module

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
)

func Auth(c *gin.Context) {
	var result gin.H
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){

		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		result = gin.H{
			"status": "success",
			"pesan": "token verified",
		}
		c.JSON(http.StatusOK, result)
		c.Abort()
	} else {
		result = gin.H {
			"status": "error " + err.Error(),
			"pesan": "not authorized",
			"token": token,
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func CekAuth(c *gin.Context) bool{
	var status bool
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		status = true
	} else {
		status = false
	}
	return status
}