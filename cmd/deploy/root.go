package deploy

import (
	"github.com/a3e/a3e/pkg/aci"
	"github.com/a3e/a3e/pkg/err"
	"github.com/spf13/cobra"
)

// Root returns the command for the root of the deploy command tree
func Root() *cobra.Command {
	return &cobra.Command{
		Use: "deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			cl := aci.NewClient("TODO")
			return err.Check(
				cl.Deploy("TODOSub", "TODORG", "TODOContainerGroup"),
				"Deployment Complete!",
			)
		},
	}
}
