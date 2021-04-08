package nomad

import (
	"context"
	"fmt"

	"github.com/hashicorp/nomad/api"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// ACMService returns the service connection for AWS ACM service
func NomadClient(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {
	logger := plugin.Logger(ctx)

	clientConfig := api.DefaultConfig()

	nomadConfig := GetConfig(d.Connection)
	if &nomadConfig != nil {
		if nomadConfig.Address != nil {
			logger.Trace("NomadClient:: using provided address / conf.Address:", nomadConfig.Address)
			clientConfig.Address = *nomadConfig.Address
		}
		if nomadConfig.Region != nil {
			logger.Trace("NomadClient:: using provided region:", nomadConfig.Region)
			clientConfig.Region = *nomadConfig.Region
		}
		if nomadConfig.SecretID != nil {
			logger.Trace("NomadClient:: using provided secret ID")
			clientConfig.SecretID = *nomadConfig.SecretID
		}
	}

	client, err := api.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to configure Nomad API: %s", err)
	}

	return client, nil
}
