package beta

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	var cmdRoot = &cobra.Command{
		Use:   "beta",
		Short: "Beta commands",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmdRoot.AddCommand(cmdSimple, cmdEcho())
	return cmdRoot
}

// cmdEcho is a sample command which is used as a stub for more advanced usage.
func cmdEcho() *cobra.Command {
	ret := &cobra.Command{
		Use:   "echo",
		Short: "A sample command which is used as a stub for more advanced usage.",
	}

	value := ret.Flags().StringP("value", "v", "", "Value to echo")
	ret.MarkFlagRequired("value")

	ret.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Printf("echo: %s\n", *value)
		return nil
	}
	return ret
}

// cmdSimple is a sample command which shows very simple usage.
var cmdSimple = &cobra.Command{
	Use:   "hello",
	Short: "A sample command which shows very simple usage.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello, world.\n")
	},
}
