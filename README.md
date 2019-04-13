# a3e

Simple container deployment

# Sample Configuration

Configuration is done with [starlark](https://github.com/google/starlark-go/).

Here's an example config file:

```python
app = app("athens")
app.locations(["eastus"])
app.cloud_info([
    {
        "type": "azure",
        "subscription_id": "5ea9ae04-3601-468a-ba84-cb7e82ae1e48",
        "resource-group": "my-resource-group"
    }
])

athens = app.container(image="gomods/athens:v0.3.1", ports=[3000])
athens.env("IN_THE_CLEAR", default="SOMETHING")
athens.env("SECRET_VAR", secret=true)
```
