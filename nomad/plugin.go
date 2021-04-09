/*
Package nomad implements a steampipe plugin for nomad.

This plugin provides data that Steampipe uses to present foreign
tables that represent Amazon Nomad resources.
*/
package nomad

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	// "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-nomad"

// Plugin creates this (nomad) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		// DefaultTransform: transform.FromCamel(),
		// DefaultGetConfig: &plugin.GetConfig{
		// 	ShouldIgnoreError: isNotFoundError([]string{"ResourceNotFoundException", "NoSuchEntity"}),
		// },
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"nomad_node": tableNomadNode(ctx),
			"nomad_job":  tableNomadJob(ctx),
		},
	}

	return p
}
