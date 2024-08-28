package main

import (
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func main() {
	fmt.Println("Testing Snyk license detection for AGPL-3.0 (Grafana plugin SDK).")

	// 使用 grafana-plugin-sdk-go 库示例
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println("Example response:", response)
}
