package rc

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const DefaultFileName = ".a3e.toml"

// Params are the parameters that go in the rc file
type Params struct {
	ClientID string `json:"CLIENT_ID" envconfig:"CLIENT_ID" required:"true" default:"04b07795-8ddb-461a-bbee-02f9e1bf7b46"`
	TenantID string `json:"TENANT_ID" envconfig:"TENANT_ID" required:"true"`
	SubID    string `json:"SUBSCRIPTION_ID" envconfig:"SUBSCRIPTION_ID" required:"true"`
	ResGroup string `json:"RESOURCE_GROUP" envconfig:"RESOURCE_GROUP" required:"true:"`
}

// FetchFromFile fetches parameters from the given private params file. Returns
// nil and a suitable error if the file was not found or there was an
// error reading it. Pass the empty string for fname if you want to use
// the default (DefaultFileName)
func FetchFromFile(fname string) (*Params, error) {
	if fname == "" {
		fname = DefaultFileName
	}
	fileBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret := new(Params)
	if _, err := toml.Decode(string(fileBytes), ret); err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

func FetchFromEnv() (*Params, error) {
	ret := new(Params)
	if err := envconfig.Process("A3E", ret); err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}
