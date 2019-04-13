package config

import (
	corev1 "k8s.io/api/core/v1"
	apps "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type Converter interface {
	// k8s.io/client-go/kubernetes/typed/apps/v1
	Deployments() []apps.Deployment
	// k8s.io/client-go/kubernetes/typed/core/v1
	Secrets() []corev1.Secret
	Services() []corev1.Service
}

// Decode decodes the config file at filename into a root config struct.
//
// Returns nil and an error if there was a problem reading or decoding the file
func Decode(filename string) (Converter, error) {
	// bytes, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	return nil, err
	// }
	// ret := new(Root)
	// if err := hcl.Decode(ret, string(bytes)); err != nil {
	// 	return nil, err
	// }
	// return ret, nil
	return nil, nil
}
