---
layout: "anypoint"
page_title: "Anypoint: anypoint_business_group"
sidebar_current: "docs-anypoint-resource-business-group"
description: |-
  Provides an Anypoint Business Group resource.
---

# anypoint\_business_group

Provides resource to allow you to manage Anypoint Business
Groups within your organization. 

## Example Usage

```hcl
# Create a new Business Group
resource "anypoint_business_group" "default" {
  name       = "Terraform Example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Business Group

## Attributes Reference

The following attributes are exported:

* `id` - The unique ID of the business group
* `name` - The name of the business group
* `client_id` - The Client ID of the business group

## Import

Business Groups can not yet be imported
