package cmd

import (
	"github.com/a3e/a3e/pkg/cfg"
	"github.com/spf13/cobra"
)

// Skeleton returns a basic cobra.Command for use all throughout the codebase
func Skeleton(use, short string) *cobra.Command {
	ret := &cobra.Command{
		Use:           use,
		Short:         short,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	configFileName := "a3e.hcl"
	ret.PersistentFlags().StringVar(
		&configFileName,
		"config",
		configFileName,
		"The name of the config file",
	)
	ret.PreRunE = func(cmd *cobra.Command, args []string) error {
		return cfg.Parse(configFileName)
	}
	return ret
}
