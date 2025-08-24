package main

import (
	setting "github.com/gox7/request/internal/config"
	"github.com/gox7/request/internal/exports"
	"github.com/gox7/request/internal/request"
	"github.com/gox7/request/internal/util"
)

func main() {
	config := new(setting.Config)
	setting.Parse(config)

	util.NewTest(config)
	request.Execute(config)
	exports.Export(config)
}
