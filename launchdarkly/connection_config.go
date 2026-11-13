package launchdarkly

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type launchdarklyConfig struct {
	AccessToken *string `hcl:"access_token"`
}

func ConfigInstance() interface{} {
	return &launchdarklyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) launchdarklyConfig {
	if connection == nil || connection.GetConfig() == nil {
		return launchdarklyConfig{}
	}
	config, _ := connection.GetConfig().(launchdarklyConfig)
	return config
}
