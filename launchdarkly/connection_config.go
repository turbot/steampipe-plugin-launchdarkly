package launchdarkly

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type launchdarklyConfig struct {
	AccessToken      *string `cty:"access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"access_token": {
		Type: schema.TypeString,
	},
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
