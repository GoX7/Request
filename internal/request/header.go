package request

import (
	"fmt"
	"os"
	"strings"

	"github.com/gox7/request/internal/config"
)

func Execute(cfg *config.Config) {
	if cfg.NET.Address == "none" {
		fmt.Println("Usage: -a [address]")
		os.Exit(1)
	}

	if strings.ToLower(cfg.Protocol) == "http" {
		if strings.ToLower(cfg.NET.Method) == "get" || strings.ToLower(cfg.NET.Method) == "post" {
			data := doRequest(cfg)
			if data != "" {
				config.NewResult(cfg, string(data))
				return
			}
		} else {
			fmt.Printf("Method \"%s\" is invalid. All methods: GET, POST\n", cfg.NET.Method)
			fmt.Println("Usage: -X [method]")
			os.Exit(1)
		}
	}

	fmt.Printf("Protocol \"%s\" is invalid. All protocol: HTTP\n", cfg.Protocol)
	fmt.Println("Usage: -p [protocol]")
	os.Exit(1)
}
