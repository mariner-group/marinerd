package bootstrap

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/mariner-group/marinerd/internal/pkg/config"
)

func Start(configFilePath string) {
	cfg, err := config.Load(configFilePath)
	if err != nil {
		fmt.Printf("[ERROR] load config err: %s\n", err)
		return
	}

	c, err := yaml.Marshal(cfg)
	if err != nil {
		fmt.Printf("[ERROR] config yaml marshal fail\n")
		return
	}
	fmt.Printf(string(c))

	// TODO initial logger

	// initial context
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	go cfg.APIServer.Run()
}
