package main

import (
	"fmt"

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
}
