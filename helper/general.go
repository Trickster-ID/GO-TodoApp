package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func Ifelse(param1 interface{}, param2 interface{}) interface{} {
	if param1 == 0 || param1 == "" {
		return param2
	}
	return param1
}

func Getjwtdata(token string) interface{} {
	respar, _ := jwt.Parse(token, func(tok *jwt.Token)(interface{}, error){
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", tok.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	
	if respar.Valid {
		claims := respar.Claims.(jwt.MapClaims)
		userid := fmt.Sprintf("%v", claims["user_id"])
		log.Println("active userid : ", userid)
		id, errconv := strconv.Atoi(userid)
		if errconv != nil{
			return 0
		}
		return id
	}else{
		return 0
	}
}
