package cmd

import (
	"github.com/a3e/a3e/pkg/config"
	"github.com/spf13/cobra"
)

// Skeleton returns a basic cobra.Command for use all throughout the codebase
func Skeleton(use, short string) *cobra.Command {
	ret := &cobra.Command{
		Use:   use,
		Short: short,
		//SilenceUsage:  true,
		//SilenceErrors: true,
	}

	configFileName := ret.Flags().StringP("filename", "f", "a3e.hcl", "The name of the config file.")
	ret.PreRunE = func(cmd *cobra.Command, args []string) error {
		return config.Parse(*configFileName)
	}
	return ret
}
