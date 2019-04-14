package alpha

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/a3e/a3e/pkg/arm"
	"github.com/a3e/a3e/pkg/privcfg"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	var cmdRoot = &cobra.Command{
		Use:   "alpha",
		Short: "Alpha commands",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	cmdRoot.AddCommand(cmdResourcesDeploySubscription, cmdResourcesDeployGroupParameters, cmdResourcesDeployGroupEmpty, cmdListResources, cmdLogin)
	return cmdRoot
}

var cmdResourcesDeploySubscription = &cobra.Command{
	Use:   "resources-deploy-subscription",
	Short: "A subscription level ARM deployment that deploys multiple resource groups.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		armTemplate := `
		{
			"$schema": "https://schema.management.azure.com/schemas/2018-05-01/subscriptionDeploymentTemplate.json#",
			"contentVersion": "1.0.0.1",
			"parameters": {
				"rgNamePrefix": {
					"type": "string",
					"defaultValue": "190300-test-"
				},
				"rgLocation": {
					"type": "string",
					"defaultValue": "eastus"
				},
				"instanceCount": {
					"type": "int",
					"defaultValue": 2
				}
			},
			"variables": {},
			"resources": [
				{
					"type": "Microsoft.Resources/resourceGroups",
					"apiVersion": "2018-05-01",
					"location": "[parameters('rgLocation')]",
					"name": "[concat(parameters('rgNamePrefix'), copyIndex())]",
					"copy": {
						"name": "rgCopy",
						"count": "[parameters('instanceCount')]"
					},
					"properties": {}
				}
			],
			"outputs": {}
		}`

		template1 := map[string]interface{}{}
		if err := json.Unmarshal([]byte(armTemplate), &template1); err != nil {
			log.Fatal(err)
		}

		result, err := deployTemplateSubscription("deploy-1", "eastus", template1, nil)
		if err != nil {
			log.Fatal(err)
		}
		_ = result

	},
}

var cmdResourcesDeployGroupParameters = &cobra.Command{
	Use:   "resources-deploy-group-parameters",
	Short: "A group level ARM deployment that deploys an ARM template with parameters.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		armTemplate := `
		{
			"$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
			"contentVersion": "1.0.0.0",
			"parameters": {
				"location": {
					"type": "String"
				},
				"storageAccountName": {
					"type": "String"
				},
				"accountType": {
					"type": "String"
				},
				"kind": {
					"type": "String"
				},
				"accessTier": {
					"type": "String"
				},
				"supportsHttpsTrafficOnly": {
					"type": "Bool"
				}
			},
			"variables": {},
			"resources": [
				{
					"type": "Microsoft.Storage/storageAccounts",
					"sku": {
						"name": "[parameters('accountType')]"
					},
					"kind": "[parameters('kind')]",
					"name": "[parameters('storageAccountName')]",
					"apiVersion": "2018-07-01",
					"location": "[parameters('location')]",
					"properties": {
						"accessTier": "[parameters('accessTier')]",
						"supportsHttpsTrafficOnly": "[parameters('supportsHttpsTrafficOnly')]"
					},
					"dependsOn": []
				}
			],
			"outputs": {}
		}`

		armParameters := `
		{
			"location": {
				"value": "eastus"
			},
			"storageAccountName": {
				"value": "test190300"
			},
			"accountType": {
				"value": "Standard_RAGRS"
			},
			"kind": {
				"value": "StorageV2"
			},
			"accessTier": {
				"value": "Hot"
			},
			"supportsHttpsTrafficOnly": {
				"value": true
			}
		}
		`

		parameters1 := map[string]interface{}{}
		if err := json.Unmarshal([]byte(armParameters), &parameters1); err != nil {
			log.Fatal(err)
		}

		template1 := map[string]interface{}{}
		if err := json.Unmarshal([]byte(armTemplate), &template1); err != nil {
			log.Fatal(err)
		}

		result, err := deployTemplateGroup("deploy-1", "190300-test-0", template1, parameters1)
		if err != nil {
			log.Fatal(err)
		}
		_ = result

	},
}

var cmdResourcesDeployGroupEmpty = &cobra.Command{
	Use:   "resources-deploy-group-empty",
	Short: "A group level ARM deployment that deploys an empty template to remove all resources in the group.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		armTemplate := `
		{
			"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
			"contentVersion": "1.0.0.0",
			"resources": [

			]
		}`

		template1 := map[string]interface{}{}
		if err := json.Unmarshal([]byte(armTemplate), &template1); err != nil {
			log.Fatal(err)
		}

		result, err := deployTemplateGroup("deploy-1", "190300-test-0", template1, nil)
		if err != nil {
			log.Fatal(err)
		}
		_ = result
	},
}

var cmdListResources = &cobra.Command{
	Use:   "resources-list",
	Short: "List all resources in the subscription.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listResources()
	},
}

var cmdLogin = &cobra.Command{
	Use:   "login",
	Short: "Login to Azure using the device code login flow.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func deployTemplateSubscription(deploymentName, location string, template, parameters interface{}) (res resources.DeploymentExtended, err error) {

	config := NewConfig()
	client := resources.NewDeploymentsClient(config.SubscriptionID)
	client.Authorizer = config.Authorizer

	ctx := context.Background()
	deployment := resources.Deployment{
		Location: &location,
		Properties: &resources.DeploymentProperties{
			Template:   template,
			Parameters: parameters,
			Mode:       resources.Incremental,
		},
	}
	future, err := client.CreateOrUpdateAtSubscriptionScope(ctx, deploymentName, deployment)
	if err != nil {
		return res, fmt.Errorf("cannot create deployment: %v", err)
	}
	err = future.WaitForCompletion(ctx, client.Client)
	if err != nil {
		return res, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}
	return future.Result(client)

}

func deployTemplateGroup(deploymentName, resourceGroup string, template, parameters interface{}) (res resources.DeploymentExtended, err error) {

	config := NewConfig()
	client := resources.NewDeploymentsClient(config.SubscriptionID)
	client.Authorizer = config.Authorizer

	ctx := context.Background()
	deployment := resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template:   template,
			Parameters: parameters,
			Mode:       resources.Complete,
		},
	}
	future, err := client.CreateOrUpdate(ctx, resourceGroup, deploymentName, deployment)
	if err != nil {
		return res, fmt.Errorf("cannot create deployment: %v", err)
	}
	err = future.WaitForCompletion(ctx, client.Client)
	if err != nil {
		return res, fmt.Errorf("cannot get the create deployment future response: %v", err)
	}
	return future.Result(client)

}

func listResources() {

	config := NewConfig()
	client := resources.NewClient(config.SubscriptionID)
	client.Authorizer = config.Authorizer

	// list all resources
	for result, err := client.ListComplete(context.Background(), "", "", nil); result.NotDone(); err = result.Next() {
		if err != nil {
			log.Fatal(err)
		}
		b, _ := json.Marshal(result.Value())

		fmt.Printf("%s\n", b)
	}

}

func login() {
	config := NewConfig()
	client := resources.NewClient(config.SubscriptionID)
	client.Authorizer = config.Authorizer
}

type Config struct {
	*privcfg.Params
	Authorizer autorest.Authorizer
}

func NewConfig() *Config {
	config := new(Config)
	paramsFromFile, fileErr := privcfg.FetchFromFile("")
	paramsFromEnv, envErr := privcfg.FetchFromEnv()
	if fileErr != nil && envErr != nil {
		log.Fatal("AZURE_APPLICATION_ID, AZURE_TENANT_ID and AZURE_SUBSCRIPTION_ID are required.")
	} else if envErr == nil { // env comes first
		config.Params = paramsFromEnv
	} else if fileErr == nil {
		config.Params = paramsFromFile
	}

	tmp, err := arm.NewDeviceCodeAuthorizerWithCache()
	if err != nil {
		log.Fatal(err)
	}
	config.Authorizer = tmp
	return config
}
