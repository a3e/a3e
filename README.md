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
            { name = "IN_THE_CLEAR", val = "SOMETHING" },
            { name = "SECRET_VAR", from-env = "SECRET_LOCALLY" }
        ]
    }
]
```
