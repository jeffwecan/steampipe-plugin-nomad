package nomad

import (
	"context"
	"fmt"

	"github.com/hashicorp/nomad/api"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// ACMService returns the service connection for AWS ACM service
func NomadClient(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {

	// have we already created and cached the service?
	// serviceCacheKey := fmt.Sprintf("nomad-%s", region)
	// if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
	// 	return cachedData.(*api.Client), nil
	// }

	conf := api.DefaultConfig()
	conf.Address = *GetConfig(d.Connection).Address

	// conf.Region = d.Get("region").(string)
	// conf.SecretID = d.Get("secret_id").(string)

	// // HTTP basic auth configuration.
	// httpAuth := d.Get("http_auth").(string)
	// if httpAuth != "" {
	// 	var username, password string
	// 	if strings.Contains(httpAuth, ":") {
	// 		split := strings.SplitN(httpAuth, ":", 2)
	// 		username = split[0]
	// 		password = split[1]
	// 	} else {
	// 		username = httpAuth
	// 	}
	// 	conf.HttpAuth = &api.HttpBasicAuth{Username: username, Password: password}
	// }

	// // TLS configuration items.
	// conf.TLSConfig.CACert = d.Get("ca_file").(string)
	// conf.TLSConfig.ClientCert = d.Get("cert_file").(string)
	// conf.TLSConfig.ClientKey = d.Get("key_file").(string)
	// conf.TLSConfig.CACertPEM = []byte(d.Get("ca_pem").(string))
	// conf.TLSConfig.ClientCertPEM = []byte(d.Get("cert_pem").(string))
	// conf.TLSConfig.ClientKeyPEM = []byte(d.Get("key_pem").(string))

	// // Get the vault token from the conf, VAULT_TOKEN
	// // or ~/.vault-token (in that order)
	// var err error
	// vaultToken := d.Get("vault_token").(string)
	// if vaultToken == "" {
	// 	vaultToken, err = getToken()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to configure Nomad API: %s", err)
	}

	return client, nil
}
