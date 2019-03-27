package aci

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func auth(cl *gorequest.SuperAgent, token string) *gorequest.SuperAgent {
	return cl.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}
