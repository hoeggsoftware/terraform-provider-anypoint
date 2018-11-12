package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hoeggsoftware/terraform-provider-anypoint/anypoint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: anypoint.Provider})
}
