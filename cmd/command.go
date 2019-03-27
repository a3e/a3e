package cmd

import "github.com/spf13/cobra"

// Skeleton returns a basic cobra.Command for use all throughout the codebase
func Skeleton(use, short string) *cobra.Command {
	return &cobra.Command{
		Use:           use,
		Short:         short,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
}
