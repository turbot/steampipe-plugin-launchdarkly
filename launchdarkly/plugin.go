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
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"launchdarkly_access_token":   tablelaunchdarklyAccessToken(ctx),
			"launchdarkly_account_member": tablelaunchdarklyAccountMember(ctx),
			"launchdarkly_audit_log":      tablelaunchdarklyAuditLog(ctx),
			"launchdarkly_environment":    tablelaunchdarklyEnvironment(ctx),
			"launchdarkly_feature_flag":   tablelaunchdarklyFeatureFlag(ctx),
			"launchdarkly_project":        tablelaunchdarklyProject(ctx),
			"launchdarkly_team":           tablelaunchdarklyTeam(ctx),
		},
	}
	return p
}
