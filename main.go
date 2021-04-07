package main

import (
	"github.com/jeffwecan/steampipe-plugin-nomad/nomad"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: nomad.Plugin})
}
