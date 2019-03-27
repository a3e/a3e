package aci

import (
	"fmt"
	"net/http"

	"github.com/a3e/a3e/pkg/log/human"
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
	armTpl := newARMTpl("TODOName", nil) // TODO: containers!
	cl := gorequest.New().
		Put(deployURL(subID, rg, containerGroup)).
		SendStruct(armTpl)
	cl = auth(cl, c.token)
	human.Debug("sending the request")
	resp, _, errs := cl.End()
	if errs != nil {
		return aggregateErrs(errs)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf(
			"The request to deploy to Azure failed (got response code %d)",
			resp.StatusCode,
		)
	}
	return nil
}
