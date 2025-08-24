package exports

import (
	"fmt"
	"os"

	"github.com/gox7/request/internal/config"
)

func Export(cfg *config.Config) {
	if cfg.Export == "stdout" {
		fmt.Println(config.GetResult(cfg))
	} else {
		err := os.WriteFile(cfg.Export, []byte(config.GetResult(cfg)), 0644)
		if err != nil {
			fmt.Printf("Error export result on file(%s): %s\n", cfg.Export, err.Error())
		}
	}
}
