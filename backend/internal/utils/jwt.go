package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// 在這邊實作 JWT 的產生邏輯
// 能夠接受任何 struct 並將其轉換成 JWT
func GenerateJWT(obj interface{}) (string, error) {
	secretKey := viper.GetString("jwt.secret")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	val := reflect.ValueOf(obj)
	typeOfObj := val.Type()

	for i := 0; i < val.NumField(); i++ {
		claims[typeOfObj.Field(i).Name] = val.Field(i).Interface()
	}

	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 在這邊實作 JWT 的解析邏輯
// 能夠接受 JWT 並將其內容轉換成 struct
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := viper.GetString("jwt.secret")
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
