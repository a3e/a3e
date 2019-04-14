# a3e

Simple container deployment

# Sample Config

```hcl
name = "athens"
locations = ["eastus"]
containers = [
    {
        image = "gomods/athens:v0.3.1"
        ports = [3000]
        env = [
            // You can specify a default value for an environment variable,
            // but you can always override it with "A3E_<the env var name>".
            //
            // For example, you can override this environment variable with
            // "A3E_IN_THE_CLEAR"
            { name = "IN_THE_CLEAR", default = "SOMETHING" },
            // In this example, there is no default, so you have to set it
            // in your environment. That means if you don't set "A3E_SECRET_VAR"
            // in your environment, the deployment will fail
            { name = "SECRET_VAR" }
        ]
    }
]
```
