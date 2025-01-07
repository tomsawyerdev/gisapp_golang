package middlewares

import (
	"fmt"
	"gisapi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	//"strings"
	//"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Middleware logic before request

		// Middleware logic after request
		//fmt.Println("Method:", c.Request.Method)
		fmt.Println("Logger:", c.GetHeader("Accept"))

		c.Next()

	}
}

type Claims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Middleware logic before request

		// Middleware logic after request
		fmt.Println("Is Auth: ----------------------------")
		//fmt.Println("Method:", c.Request.Method)
		//fmt.Println("Accept:", c.GetHeader("Accept"))
		//fmt.Println("Authorization:", c.Request.Header["Authorization"])
		//fmt.Println("Session:", c.Request.Header["Session"])

		headers := c.Request.Header
		//val, ok := c.Request.Header["Authorization"]
		val, ok := headers["Session"]
		if ok {

			//fmt.Printf("Authorization key header is present with value %s\n", val)
			//strings.Split(val[0], " ")

			tokenString := val[0] // strings.Split(val[0], " ")[1]

			//fmt.Println("Token:", tokenString)
			claims := Claims{}

			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("supersecret"), nil
			})

			//fmt.Println("Exp:", claims.ExpiresAt, time.Now().Unix()-claims.ExpiresAt.Unix())
			///fmt.Println("Expired ?:", (time.Now().Unix()-claims.ExpiresAt.Unix()) > 0)

			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "Invalid token"})
				c.Abort()
				return
			}

			//fmt.Println("Claims", claims)
			var userid int

			userid, err = models.GetUserIdFromUUID(claims.UUID)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "Bad LookUp"})
				c.Abort()
				return
			}

			c.Set("userid", userid)
			c.Next()
		} else {
			//fmt.Println("Authorization key header is not present")
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "StatusUnauthorized"})
			c.Abort()
			//c.Next()
			//c.AbortWithStatus(401)

			//return
		}
	}
}

func SetId() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("userid", 1)
		c.Next()
	}
}
