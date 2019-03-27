package aci

import (
	"strconv"
)

type container struct {
	name  string
	image string
	ports []int
}

func newARMTpl(name string, containers []container) *aciARMTpl {
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
			ports = append(ports, aciARMTpl_sub7{
				Port: strconv.Itoa(port),
			})
		}
		ret.Properties.Containers = append(
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
	return ret
}
