package config

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// Root is the root of a config file
type Root struct {
	Name       string      `json:"name"`
	Locations  []string    `json:"locations"`
	Containers []Container `json:"containers"`
}

// Container is a single container in a config file
type Container struct {
	Image string            `json:"image"`
	Ports []int             `json:"ports"`
	Env   map[string]EnvVal `json:"env"`
}

// EnvVal is a single environment variable in a config file
type EnvVal struct {
	Name    string  `json:"name"`
	FromEnv *string `json:"from-env,omitempty"`
	Val     *string `json:"val,omitempty"`
}

// Value gets the value of the environment variable. It first checks
// whether the value is set and returns it if so. Otherwise, it checks the
// environment variable specified in FromEnv.
//
// For easy testing purposes, this function uses envReader to look up
// environment variables. If you want to look up environment variables
// on the underlying host, simply pass 'os.Getenv' to this function.
//
// If either is set to the empty string, this returns the empty string
//
// If Val is not set and FromEnv is set to an environment variable that
// doesn't exist (including the empty string), this returns the empty string
//
// If neither FromEnv nor Val is set, this returns the empty string
func (e EnvVal) Value(envReader func(string) string) string {
	if e.Val != nil {
		return *e.Val
	}
	if e.FromEnv != nil {
		return envReader(*e.FromEnv)
	}
	return ""
}

// Decode decodes the config file at filename into a root config struct.
//
// Returns nil and an error if there was a problem reading or decoding the file
func Decode(filename string) (*Root, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ret := new(Root)
	if err := hcl.Decode(ret, string(bytes)); err != nil {
		return nil, err
	}
	return ret, nil
}
