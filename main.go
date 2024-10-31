package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/urfave/negroni"
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

	// 使用 jwt-go 库生成 JWT token
	token := createJWTToken()
	fmt.Println("Generated JWT token:", token)

	// 使用 negroni 中间件示例
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Negroni!"))
	}))
	http.ListenAndServe(":3000", n)
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
