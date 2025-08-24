package config

import (
	"flag"
)

func Parse(config *Config) {
	method := flag.String("X", "GET", "method for request")
	header := flag.String("H", "none", "header for request")
	body := flag.String("B", "", "request body")
	timeout := flag.Int64("t", 10, "request timeout (second)")
	protocol := flag.String("p", "http", "request protocol")
	address := flag.String("a", "none", "request address")
	export := flag.String("o", "stdout", "request export (filename)")

	test_request := flag.Bool("test_request", false, "test function 1")
	test_export := flag.Bool("test_export", false, "test function 2")
	flag.Parse()

	*config = Config{
		NET: &Net{
			Address: *address,
			Method:  *method,
			Timeout: *timeout,
			Header:  *header,
			Body:    *body,
		},
		Test: &Test{
			Request: *test_request,
			Export:  *test_export,
		},
		Protocol: *protocol,
		Export:   *export,
		result:   "none",
	}
}

func NewResult(config *Config, result string) {
	config.result = result
}

func GetResult(config *Config) string {
	return config.result
}
