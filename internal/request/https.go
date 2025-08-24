package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gox7/request/internal/config"
)

func doRequest(cfg *config.Config) string {
	request, err := http.NewRequest("GET", cfg.NET.Address, bytes.NewBuffer([]byte(cfg.NET.Body)))
	if err != nil {
		return "Status: ERROR(new request)\nError: " + err.Error()
	}

	request.Header.Set("User-Agent", "request/1.0")
	if cfg.NET.Header != "none" {
		header := strings.Split(cfg.NET.Header, ":")
		request.Header.Set(header[0], header[1])
	}

	client := http.Client{Timeout: time.Second * time.Duration(cfg.NET.Timeout)}
	response, err := client.Do(request)
	if err != nil {
		return "Status: ERROR(request)\nError: " + err.Error()
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "Status: ERROR(read body)\nError: " + err.Error()
	}

	return fmt.Sprintf("Status: OK\nCode: %d\nBody: %s", response.StatusCode, string(body))
}
