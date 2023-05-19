package launchdarkly

import (
	"context"

	ldapi "github.com/launchdarkly/api-client-go/v13"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_environment",
		Description: "Fetch a list of all environments.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listEnvironments,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_key", "key"}),
			Hydrate:    getEnvironment,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "ID of the environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "A project-unique key for the new environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "A human-friendly name for the new environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_key",
				Description: "API key to use with client-side SDKs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mobile_key",
				Description: "API key to use with mobile SDKs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "color",
				Description: "The color used to indicate this environment in the UI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default_ttl",
				Description: "The default time (in minutes) that the PHP SDK can cache feature flag rules locally.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "secure_mode",
				Description: "Ensures that one end user of the client-side SDK cannot inspect the variations for another end user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "default_track_events",
				Description: "Enables tracking detailed information for new flags by default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "require_comments",
				Description: "Whether members who modify flags and segments through the LaunchDarkly user interface are required to add a comment.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "confirm_changes",
				Description: "Whether members who modify flags and segments through the LaunchDarkly user interface are required to confirm those changes.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "project_key",
				Description: "The key of this project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "links",
				Description: "Links to other resources within the API. Includes the URL and content type of those resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags for this environment.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "approval_settings",
				Description: "Describe the approval settings of an environment",
				Type:        proto.ColumnType_JSON,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

type launchdarklyProjectEnvironment struct {
	ldapi.Environment
	ProjectKey string
}

// LIST FUNCTION

func listEnvironments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	project := h.Item.(ldapi.Project)
	projectKey := project.Key

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_environment.listEnvironments", "connection_error", err)
		return nil, err
	}

	params := client.EnvironmentsApi.GetEnvironmentsByProject(ctx, projectKey)

	count := 0

	for {
		environments, _, err := params.Execute()
		if err != nil {
			plugin.Logger(ctx).Error("launchdarkly_environment.listEnvironments", "api_error", err)
			return nil, err
		}

		for _, environment := range environments.Items {
			d.StreamListItem(ctx, launchdarklyProjectEnvironment{environment, projectKey})
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count += len(environments.Items)
		if count >= int(environments.GetTotalCount()) {
			break
		}
		params.Offset(int64(count))
	}
	return nil, nil
}

func getEnvironment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_environment.getEnvironment", "connection_error", err)
		return nil, err
	}

	projectKey := d.EqualsQualString("project_key")
	environmentKey := d.EqualsQualString("key")

	environment, _, err := client.EnvironmentsApi.GetEnvironment(ctx, projectKey, environmentKey).Execute()

	if err != nil {
		logger.Error("launchdarkly_environment.getEnvironment", "api_error", err)
		return nil, err
	}

	return launchdarklyProjectEnvironment{*environment, projectKey}, nil
}
