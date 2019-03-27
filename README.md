# a3e

Simple container deployment

# Sample Config

```hcl
name = "athens"
locations = ["eastus"]
auth = {
    subscription-id = "5ea9ae04-3601-468a-ba84-cb7e82ae1e48"
    resource-group = "my-resource-group"
}
containers = [
    {
        image = "gomods/athens:v0.3.1"
        ports = [3000]
        env = [
            { name = "IN_THE_CLEAR", val = "SOMETHING" },
            { name = "SECRET_VAR", from-env = "SECRET_LOCALLY" }
        ]
    }
]
```
