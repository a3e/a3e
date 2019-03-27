package main

import (
	"github.com/a3e/a3e/aci"
	"github.com/spf13/cobra"
)

func deploy() *cobra.Command {
	return &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			cl := aci.NewClient("TODO")
			return cl.Deploy("TODOSub", "TODORG", "TODOContainerGroup")
		},
	}

}
