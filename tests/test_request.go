package tests

import (
	"fmt"

	"github.com/gox7/request/internal/config"
	"github.com/gox7/request/internal/request"
)

func TEST_Request() {
	fmt.Println("Start request test...")
	cfg := &config.Config{
		NET: &config.Net{
			Method:  "GET",
			Address: "https://ifconfig.me/ip",
			Header:  "none",
			Timeout: 15,
		},
		Protocol: "http",
	}

	fmt.Println("Request...")
	request.Execute(cfg)

	if config.GetResult(cfg) != "" {
		fmt.Println("Request test: OK")
	} else {
		fmt.Println("Request test: ERROR")
	}
}
