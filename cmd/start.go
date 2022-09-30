package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mariner-group/marinerd/internal/pkg/bootstrap"
)

var (
	configFilePath = ""

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start running",
		Long:  "start running",
		Run: func(c *cobra.Command, args []string) {
			bootstrap.Start(configFilePath)
		},
	}
)

// init parse command-line
func init() {
	startCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "config.yaml", "config file path")
}
