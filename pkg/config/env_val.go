package config

import "fmt"

type envVal struct {
	name   string
	def    string
	secret bool
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
func (e envVal) value(envReader func(string) string) (string, error) {
	if e.secret {
		return "", fmt.Errorf("secret envs no supported yet :(")
	}
	fromEnv := envReader(e.name)
	hasDef := e.def != ""
	if fromEnv == "" && !hasDef {
		err := fmt.Errorf("environment variable %s missing, and no default set", e.name)
		return "", err
	}
	return fromEnv, nil
}
