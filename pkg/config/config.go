package config

import (
	"io/ioutil"
	"log"

	"github.com/starlight-go/starlight"
)

type Converter interface {
	// k8s.io/client-go/kubernetes/typed/apps/v1
	// Deployments() []apps.Deployment
	// k8s.io/client-go/kubernetes/typed/core/v1
	// Secrets() []corev1.Secret
	// Services() []corev1.Service
}

// Decode decodes the config file at filename into a root config struct.
//
// Returns nil and an error if there was a problem reading or decoding the file
func Decode(filename string) (Converter, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	root := &root{}
	globals := map[string]interface{}{
		"app": root.app,
	}
	if _, err := starlight.Eval(bytes, globals, nil); err != nil {
		return nil, err
	}
	log.Printf("here is everythjing: %#V", *root)
	return nil, nil

	// ret := new(Root)
	// if err := hcl.Decode(ret, string(bytes)); err != nil {
	// 	return nil, err
	// }
	// return ret, nil
	return nil, nil
}
