package middleware

import (
	"Todoapp/helper"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			response := helper.BuildErrorResponse("failed to process requeest", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		//token, err := ValidateToken(authHeader)
		respar, err := jwt.Parse(authHeader, func(tok *jwt.Token)(interface{}, error){
			if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signin method %v", tok.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if respar.Valid {
			return
			//claims := respar.Claims.(jwt.MapClaims)
			//log.Println("claim[user_id] : ", claims["user_id"])
			//log.Println("claim[exp]     : ", claims["exp"])
			//log.Println("claim[iat]     : ", claims["iat"])
			//log.Println("claim[iss]     : ", claims["iss"])
		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("token is not valid", err.Error(),nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}