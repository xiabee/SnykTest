package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func main() {
	fmt.Println("Testing Snyk license detection for AGPL-3.0 and vulnerabilities.")

	// 使用 grafana-plugin-sdk-go 库示例
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println("Example response:", response)

	// 使用 gin 库示例
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	fmt.Println("Gin server is set up but not running in this example.")

	// 使用 jwt-go 库示例
	token := createJWTToken()
	fmt.Println("Generated JWT token:", token)
}

// createJWTToken 使用 github.com/dgrijalva/jwt-go 库创建一个 JWT token
func createJWTToken() string {
	// 创建一个新的 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// 使用一个示例密钥对 token 进行签名
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Error creating JWT token:", err)
		return ""
	}

	return tokenString
}
