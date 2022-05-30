---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_circuit Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/core-functionality/circuits/#circuits_1:
  A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.
  Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.
---

# netbox_circuit (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/core-functionality/circuits/#circuits_1):

> A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.
> 
> Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cid` (String)
- `provider_id` (Number)
- `status` (String)
- `type_id` (Number)

### Optional

- `tenant_id` (Number)

### Read-Only

- `id` (String) The ID of this resource.

