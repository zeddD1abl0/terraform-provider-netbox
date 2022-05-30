---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_cluster Data Source - terraform-provider-netbox"
subcategory: ""
description: |-
  
---

# netbox_cluster (Data Source)



## Example Usage

```terraform
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Read-Only

- `cluster_id` (Number)
- `id` (String) The ID of this resource.

