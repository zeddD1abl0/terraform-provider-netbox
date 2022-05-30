---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_circuit_termination Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/core-functionality/circuits/#circuit-terminations:
  The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.
  Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.
---

# netbox_circuit_termination (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/core-functionality/circuits/#circuit-terminations):

> The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.
>
> Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `circuit_id` (Number)
- `site_id` (Number)
- `term_side` (String)

### Optional

- `port_speed` (Number)
- `upstream_speed` (Number)

### Read-Only

- `id` (String) The ID of this resource.

