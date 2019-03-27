package aci

import (
	"fmt"
)

type container struct {
	name  string
	image string
	ports []int
}

func newArmTpl(name string, containers []container) *aciARMTpl {
	ret := &aciARMTpl{
		APIVersion: apiVersion,
		Type:       typeVal,
		Name:       name,
		Properties: aciARMTpl_sub23{},
	}
	ret.Properties.Containers = nil

	for _, container := range containers {
		ports := []aciARMTpl_sub7{}
		for _, port := range container.ports {
			ports = append(ports, aciARMTpl{
				Port: port,
				Name: fmt.Sprintf("port_%d", port),
			})
		}
		ret.Properties.Contaienrs = append(
			ret.Properties.Containers,
			aciARMTpl_sub13{
				Name: container.name,
				Properties: aciARMTpl_sub12{
					Image: container.image,
					Ports: ports,
				},
			},
		)
	}
}
