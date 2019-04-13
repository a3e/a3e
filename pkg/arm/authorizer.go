package arm

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	homedir "github.com/mitchellh/go-homedir"
)

// NewAuthorizer returns the correct authorizer for the given settings and/or based on the value
// of the AZURE_AUTH_METHOD environment variable, which may be one of:
// clientcredentials, clientcertificate, usernamepassword, msi, or cli (default).
func NewAuthorizer() (autorest.Authorizer, error) {
	settings := struct {
		authMethod          string
		tenantID            string
		clientID            string
		clientSecret        string
		certificatePath     string
		certificatePassword string
		username            string
		password            string
		envName             string
		resource            string
		environment         azure.Environment
	}{
		authMethod:          os.Getenv("AZURE_AUTH_METHOD"),
		tenantID:            os.Getenv("AZURE_TENANT_ID"),
		clientID:            os.Getenv("AZURE_CLIENT_ID"),
		clientSecret:        os.Getenv("AZURE_CLIENT_SECRET"),
		certificatePath:     os.Getenv("AZURE_CERTIFICATE_PATH"),
		certificatePassword: os.Getenv("AZURE_CERTIFICATE_PASSWORD"),
		username:            os.Getenv("AZURE_USERNAME"),
		password:            os.Getenv("AZURE_PASSWORD"),
		envName:             os.Getenv("AZURE_ENVIRONMENT"),
		resource:            os.Getenv("AZURE_AD_RESOURCE"),
	}

	settings.environment = azure.PublicCloud
	if settings.envName != "" {
		val, err := azure.EnvironmentFromName(settings.envName)
		if err != nil {
			return nil, err
		}
		settings.environment = val
	}

	if settings.resource == "" {
		settings.resource = strings.TrimSuffix(settings.environment.KeyVaultEndpoint, "/")
	}

	// 1. Client credentials
	if (settings.clientSecret != "") || settings.authMethod == "clientcredentials" {
		config := auth.NewClientCredentialsConfig(settings.clientID, settings.clientSecret, settings.tenantID)
		config.AADEndpoint = settings.environment.ActiveDirectoryEndpoint
		config.Resource = settings.resource
		return config.Authorizer()
	}

	// 2. Client Certificate
	if (settings.certificatePath != "") || settings.authMethod == "clientcertificate" {
		config := auth.NewClientCertificateConfig(settings.certificatePath, settings.certificatePassword, settings.clientID, settings.tenantID)
		config.AADEndpoint = settings.environment.ActiveDirectoryEndpoint
		config.Resource = settings.resource
		return config.Authorizer()
	}

	// 3. Username Password
	if (settings.username != "" && settings.password != "") || settings.authMethod == "usernamepassword" {
		config := auth.NewUsernamePasswordConfig(settings.username, settings.password, settings.clientID, settings.tenantID)
		config.AADEndpoint = settings.environment.ActiveDirectoryEndpoint
		config.Resource = settings.resource
		return config.Authorizer()
	}

	// 4. MSI
	if settings.authMethod == "msi" {
		config := auth.NewMSIConfig()
		config.Resource = settings.resource
		config.ClientID = settings.clientID
		return config.Authorizer()
	}

	// TODO: decide how to handle prompt on stdout, caching, etc.
	// 5. Device Code
	if settings.authMethod == "devicecode" {
		// TODO: This will be required on every execution. Consider caching.
		config := auth.NewDeviceFlowConfig(settings.clientID, settings.tenantID)
		return config.Authorizer()
	}

	// 5. CLI
	return auth.NewAuthorizerFromCLIWithResource(settings.resource)

}

func NewDeviceCodeAuthorizerWithCache() (autorest.Authorizer, error) {
	// using only the settings relevent to this authorizer
	settings := struct {
		tenantID string
		clientID string
	}{
		tenantID: os.Getenv("AZURE_TENANT_ID"),
		clientID: os.Getenv("AZURE_CLIENT_ID"),
	}

	// TODO: remove hard-coded cache location
	dir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	tokenCachePath := filepath.Join(dir, ".a3e")

	// load token from cache if available
	token, err := adal.LoadToken(tokenCachePath)
	config := auth.NewDeviceFlowConfig(settings.clientID, settings.tenantID)

	if err != nil {
		// authenticate via devicecode flow
		config := auth.NewDeviceFlowConfig(settings.clientID, settings.tenantID)
		spt, err := config.ServicePrincipalToken()
		if err != nil {
			return nil, err
		}
		tmp := spt.Token()
		token = &tmp
		err = adal.SaveToken(tokenCachePath, 0600, *token)
		if err != nil {
			return nil, err
		}
	} else {
		oauthConfig, err := adal.NewOAuthConfig(config.AADEndpoint, config.TenantID)
		if err != nil {
			return nil, err
		}
		spt, err := adal.NewServicePrincipalTokenFromManualToken(*oauthConfig, settings.clientID, azure.PublicCloud.ActiveDirectoryEndpoint, *token, func(t adal.Token) error {
			log.Printf("adal.NewServicePrincipalTokenFromManualToken: refreshed")
			return nil
		})
		if err != nil {
			return nil, err
		}
		if err := spt.EnsureFresh(); err != nil {
			return nil, err
		}
		tmp := spt.Token()
		err = adal.SaveToken(tokenCachePath, 0600, tmp)
		if err != nil {
			return nil, err
		}
		token = &tmp
	}
	return autorest.Authorizer(autorest.NewBearerAuthorizer(token)), nil

}
