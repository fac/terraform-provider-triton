package main

import (
	"github.com/fac/terraform-provider-triton/triton"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: triton.Provider})
}
