package main

import (
	"fmt"
	"os"

	"github.com/a3e/a3e/cmd"
	"github.com/a3e/a3e/cmd/alpha"
	"github.com/a3e/a3e/cmd/beta"
	"github.com/a3e/a3e/cmd/deploy"
	"github.com/a3e/a3e/pkg/log/human"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Root sets up the command for the top level a3e probram,
// adds all child commands, and does additional configuration.
// It only needs to be called once and then run by the main function
func Root() *cobra.Command {
	root := cmd.Skeleton("a3e", "Simple container deployment")
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	root.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.cmd.yaml)",
	)

	root.PersistentFlags().BoolVar(
		&human.IsDebugging,
		"debug",
		false,
		"Turn on debug logging",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	root.AddCommand(alpha.Root())
	root.AddCommand(beta.Root())
	root.AddCommand(deploy.Root())
	// RootCmd.AddCommand(build.Command())
	return root.Command
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
