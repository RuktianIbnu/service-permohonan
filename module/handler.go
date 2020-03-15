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

func (idb *InDB) TestMapping(c *gin.Context) {
	var (
		result gin.H
	)
	token := c.Request.Header.Get("Authorization")
	key := c.Request.Header.Get("key")
	tokenString := token   
	claims := jwt.MapClaims{}
	MapClaims, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		result = gin.H {
			"status": "error " + err.Error(),
			"pesan": "not authorized",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	} else {
		result = gin.H{
			"detail": MapClaims.Claims,
		}
		c.JSON(http.StatusOK, result)
		c.Abort()
	}
	// 	var a string
	// 	for _, val := range claims {
	// 		a = val.(string)
			
	// 	}
	// 	result = gin.H{
	// 		"x": a,
	// 	}
	// 	c.JSON(http.StatusOK, result)
	// 	c.Abort()
	// }
}