---
layout: "anypoint"
page_title: "Provider: Anypoint"
sidebar_current: "docs-do-index"
description: |-
  The Anypoint provider is used to interact with the resources supported by the Anypoint Platform APIs. The provider needs to be configured with the proper credentials before it can be used.
---

# Anypoint Provider

The Anypoint provider is used to interact with the resources supported by the Anypoint Platform APIs. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Set the variable values in *.tfvars file
# or using -var="foo_username=..." CLI option
variable "foo_username" {}
variable "foo_password" {}

# Configure the Anypoint Provider
provider "anypoint" {
  username = "${var.foo_username}"
  password = "${var.foo_password}"
}

# Create a business group
resource "anypoint_business_group" "alpha" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) This is the anypoint username. This can also be specified
  with the ANYPOINT_USERNAME shell environment variable.
* `password` - (Required) This is the anypoint password. This can also be specified
  with the ANYPOINT_PASSWORD shell environment variable.
