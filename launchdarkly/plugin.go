package launchdarkly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (launchdarkly) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-launchdarkly",
		DefaultTransform: transform.FromCamel(),
		// DefaultIgnoreConfig: &plugin.IgnoreConfig{
		// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		// },
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"launchdarkly_access_token": tablelaunchdarklyAccessToken(ctx),
			"launchdarkly_audit_log":		 tablelaunchdarklyAuditLog(ctx),
      "launchdarkly_project":      tablelaunchdarklyProject(ctx),
			"launchdarkly_environment":  tablelaunchdarklyEnvironment(ctx),
		},
	}
	return p
}
