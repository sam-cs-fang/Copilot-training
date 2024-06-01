package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 處理請求
		c.Next()

		// 計算響應時間
		latency := time.Since(t)

		// 輸出日誌，格式：[HTTP方法] 路由 請求耗時
		log.Printf("[%s] %s %s\n", c.Request.Method, c.Request.URL.Path, latency)
	}
}

// 實作一個請求的驗證機制，驗證以下內容
// 1. 請求頭中是否包含 Authorization
// 2. Authorization 是否為 Bearer token
// 3. JWT token 是否有效
// 4. 驗證成功，則將 userId 設置到 Context 中
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth := c.GetHeader("Authorization")
		// if auth == "" {
		// 	c.JSON(401, gin.H{"message": "Authorization header is required"})
		// 	c.Abort()
		// 	return
		// }

		// bearerToken := strings.Split(auth, " ")
		// if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		// 	c.JSON(401, gin.H{"message": "Authorization header must be Bearer token"})
		// 	c.Abort()
		// 	return
		// }

		// tokenString := bearerToken[1]
		// parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	jwtSecret := viper.GetString("postgres.username")
		// 	hmacSampleSecret := []byte(jwtSecret)
		// 	return hmacSampleSecret, nil
		// })

		// if err != nil {
		// 	c.JSON(401, gin.H{
		// 		"message": "Invalid token",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// claims, ok := parsedToken.Claims.(jwt.MapClaims)
		// if ok && parsedToken.Valid {
		// 	c.Set("userId", claims["userId"])
		// } else {
		// 	c.JSON(401, gin.H{
		// 		"message": "Invalid token",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		c.Set("userId", 1)

		// 處理請求
		c.Next()
	}
}
