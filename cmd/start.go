package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mariner-group/marinerd/internal/pkg/apiserver"
)

var (
	cfgFile   string
	apiServer apiserver.APIServer

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start running",
		Long:  "start running",
		Run: func(c *cobra.Command, args []string) {
			if err := apiServer.Initialize(); err != nil {
				fmt.Println("Initialize error:", err)
				os.Exit(1)
			}
			if err := apiServer.Run(); err != nil {
				fmt.Println("Run error:", err)
				os.Exit(1)
			}
		},
	}
)

// init parse command-line
func init() {
	cobra.OnInitialize(initConfig)
	startCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./conf/config.yaml", "config file path")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./conf")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	fmt.Println("Load config from ", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
