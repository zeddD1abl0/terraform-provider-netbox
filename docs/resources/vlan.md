---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_vlan Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  
---

# netbox_vlan (Resource)



## Example Usage

```terraform
resource "netbox_vlan" "example1" {
  name = "VLAN 1"
  vid = 1777
  tags = []
}

# Assume netbox_tenant, netbox_site, and netbox_tag resources exist
resource "netbox_vlan" "example2" {
  name = "VLAN 2"
  vid = 1778
  status = "reserved"
  description = "Reserved example VLAN"
  tenant_id = netbox_tenant.ex.id
  site_id = netbox_site.ex.id
  tags = [netbox_tag.ex.name]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `tags` (Set of String)
- `vid` (Number)

### Optional

- `description` (String)
- `role_id` (Number)
- `site_id` (Number)
- `status` (String)
- `tenant_id` (Number)

### Read-Only

- `id` (String) The ID of this resource.

