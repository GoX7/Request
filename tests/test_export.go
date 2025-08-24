package tests

import (
	"fmt"
	"os"

	"github.com/gox7/request/internal/config"
	"github.com/gox7/request/internal/exports"
)

func TEST_Export() {
	fmt.Println("Start export test...")
	cfg := &config.Config{Export: "test_export_request.txt"}
	config.NewResult(cfg, "test result")

	fmt.Println("Export...")
	exports.Export(cfg)

	result, err := os.ReadFile("test_export_request.txt")
	if err != nil {
		fmt.Println("Error read file: " + err.Error())
		os.Exit(1)
	}

	if result != nil {
		fmt.Println("Export test: OK")
		os.Remove("test_export_request.txt")
	} else {
		fmt.Println("Export test: ERROR")
		os.Remove("test_export_request.txt")
	}
}
