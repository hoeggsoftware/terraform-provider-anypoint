package anypoint

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ANYPOINT_USERNAME", nil),
				Description: "Anypoint platform user name",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ANYPOINT_PASSWORD", nil),
				Description: "Anypoint platform password",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"anypoint_business_group": resourceAnypointBusinessGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return config.Client()
}
