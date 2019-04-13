package deploy

import (
	"github.com/a3e/a3e/cmd"
	"github.com/a3e/a3e/pkg/config"
	"github.com/spf13/cobra"
)

type deployCmd struct {
}

func (d deployCmd) runE(*cobra.Command, []string) error {
	// cl := aci.NewClient("TODO")
	_, err := config.Decode("a3e.py")
	return err
	// return human.Check(
	// 	cl.Deploy(
	// 		config.Cfg.ID.SubscriptionID,
	// 		config.Cfg.ID.ResourceGroup,
	// 		config.Cfg.Name,
	// 	),
	// 	"Deployment Complete!",
	// )
}

// Root returns the command for the root of the deploy command tree
func Root() *cobra.Command {
	d := deployCmd{}
	ret := cmd.Skeleton("deploy", "Deploy your containers")
	ret.RunE = d.runE
	return ret.Command
}
