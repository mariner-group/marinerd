package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "marinerd",
		Short:        "marinerd",
		Long:         "marinerd",
		SilenceUsage: true,
	}
)

// init
func init() {
	rootCmd.AddCommand(startCmd)
}

// Execute 执行命令行解析
func Execute() {
	rootCmd.Execute()
}
