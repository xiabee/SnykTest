package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func main() {
	fmt.Println("Testing Snyk license detection for vulnerabilities.")

	// 使用 grafana-plugin-sdk-go 库示例
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println("Example response:", response)

	// 使用 gin 库示例
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World from Gin!")
	})

	// 使用 uuid 生成一个示例 UUID
	newUUID := uuid.New()
	fmt.Println("Generated UUID:", newUUID)

	// 启动 gin HTTP 服务器
	go func() {
		if err := r.Run(":8080"); err != nil {
			fmt.Println("Failed to start Gin server:", err)
		}
	}()

	// 保持主进程运行一段时间（模拟实际使用）
	select {}
}
