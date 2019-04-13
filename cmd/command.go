package cmd

import (
	"github.com/spf13/cobra"
)

const defaultConfigFileName = "a3e.py"

type SkeletonCmd struct {
	*cobra.Command
	ConfigFileName string
}

// Skeleton returns a basic cobra.Command for use all throughout the codebase
func Skeleton(use, short string) *SkeletonCmd {
	ret := &SkeletonCmd{
		ConfigFileName: defaultConfigFileName,
	}

	cmd := &cobra.Command{
		Use:           use,
		Short:         short,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().StringVar(
		&ret.ConfigFileName,
		"configfile",
		defaultConfigFileName,
		"The filename of the a3e config file",
	)
	ret.Command = cmd

	return ret
}
