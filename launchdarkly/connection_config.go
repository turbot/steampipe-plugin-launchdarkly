package launchdarkly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type launchdarklyConfig struct {
	AccessToken *string `hcl:"access_token"`
}

func ConfigInstance() interface{} {
	return &launchdarklyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) launchdarklyConfig {
	if connection == nil || connection.Config == nil {
		return launchdarklyConfig{}
	}
	config, _ := connection.Config.(launchdarklyConfig)
	return config
}
