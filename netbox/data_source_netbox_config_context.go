package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxConfigContext() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxConfigContextRead,
		Description: `:meta:subcategory:Extras:`,
		Schema: map[string]*schema.Schema{
			"config_context_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"data_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_file": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_context_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_context_regions": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_site_groups": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_sites": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_locations": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_device_types": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_roles": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_platforms": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_cluster_types": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_cluster_groups": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_clusters": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_tenant_groups": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"config_context_tenant": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"data_source_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxConfigContextRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := extras.NewExtrasConfigContexts
}
