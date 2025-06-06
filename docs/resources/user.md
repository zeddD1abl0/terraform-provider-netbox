---
# generated by https://github.com/fbreckle/terraform-plugin-docs
page_title: "netbox_user Resource - terraform-provider-netbox"
subcategory: "Authentication"
description: |-
  This resource is used to manage users.
---

# netbox_user (Resource)

This resource is used to manage users.

## Example Usage

```terraform
resource "netbox_user" "test" {
  username = "johndoe"
  password = "Abcdefghijkl1"
  active   = true
  staff    = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `password` (String, Sensitive)
- `username` (String)

### Optional

- `active` (Boolean) Defaults to `true`.
- `email` (String) Defaults to `""`.
- `first_name` (String) Defaults to `""`.
- `group_ids` (Set of Number)
- `last_name` (String) Defaults to `""`.
- `staff` (Boolean) Defaults to `false`.

### Read-Only

- `id` (String) The ID of this resource.


