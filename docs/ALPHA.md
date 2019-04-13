# ALPHA

```bash
TMP=$(az account show)
export AZURE_SUBSCRIPTION_ID=$(echo $TMP | jq -r .id)
export AZURE_TENANT_ID=$(echo $TMP | jq -r .tenantId)

go run . alpha login
go run . alpha resources-list
go run . alpha resources-deploy-subscription
go run . alpha resources-deploy-group-parameters
go run . alpha resources-deploy-group-empty
```
