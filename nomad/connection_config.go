package nomad

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type nomadConfig struct {
	Address      *string  `cty:"address"`
}


var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &nomadConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) nomadConfig {
	if connection == nil || connection.Config == nil {
		return nomadConfig{}
	}
	config, _ := connection.Config.(nomadConfig)
	return config
}
