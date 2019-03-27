package aci

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func deployURL(subID, rg, containerGroup string) string {
	return fmt.Sprintf(
		"https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerInstance/containerGroups/%s?api-version=2018-10-01",
		subID,
		rg,
		containerGroup,
	)
}

// Deploy deploys a container group to the specified subscription (subID) and
// resource group (subID & rg) with the given name containerGroup.
func (c *Client) Deploy(subID, rg, containerGroup string) error {
	cl := gorequest.New().
		Put(deployURL(subID, rg, containerGroup))
	cl = auth(cl, c.token)
	_, _, errs := cl.End()
	if errs != nil {
		return aggregateErrs(errs)
	}
	return nil
}
