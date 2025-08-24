package util

import (
	"github.com/gox7/request/internal/config"
	"github.com/gox7/request/tests"
)

func NewTest(cfg *config.Config) {
	if cfg.Test.Request {
		tests.TEST_Request()
	}
	if cfg.Test.Export {
		tests.TEST_Export()
	}
}
