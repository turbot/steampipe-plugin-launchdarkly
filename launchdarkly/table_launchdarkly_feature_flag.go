package launchdarkly

import (
	"context"

	ldapi "github.com/launchdarkly/api-client-go/v13"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyFeatureFlag(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_feature_flag",
		Description: "Fetch a list of all feature flags.",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listFeatureFlags,
			RetryConfig: &plugin.RetryConfig{
				ShouldRetryErrorFunc: shouldRetryError([]string{"429"}),
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project_key", "key"}),
			Hydrate:    getFeatureFlag,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A human-friendly name for the feature flag.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "A unique key used to reference the flag in your code.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "Kind of feature flag.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description of the feature flag.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "Version of the feature flag.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "creation_date",
				Description: "API key to use with mobile SDKs.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "include_in_snippet",
				Description: "TDeprecated, use clientSideAvailability. Whether this flag should be made available to the client-side JavaScript SDK",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "temporary",
				Description: "Whether the flag is a temporary flag.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "maintainer_id",
				Description: "Associated maintainerId for the feature flag.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "maintainer_team_key",
				Description: "The key of the associated team that maintains this feature flag.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "archived",
				Description: "Boolean indicating if the feature flag is archived.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "archived_date",
				Description: "Time when the feature flag has been archived.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("ArchivedDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "project_key",
				Description: "The key of this project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "variation",
				Description: "An array of possible variations for the flag.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "Tags for the feature flag.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "links",
				Description: "The location and content type of related resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "maintainer",
				Description: "Details of the maintainer for the feature flags.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "goal_ids",
				Description: "An array of goal IDs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "experiments",
				Description: "Information about the experiments related to the feature flag.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_properties",
				Description: "Information about the custom properties.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "defaults",
				Description: "The index, from the array of variations for this flag, of the variation to serve by default when targeting is on or off.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "environments",
				Description: "A JSON object containing configuration information for different environments.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "client_side_availability",
				Description: "An array of possible variations for the flag.",
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

type launchdarklyFeatureFlag struct {
	ldapi.FeatureFlag
	ProjectKey string
}

// LIST FUNCTION

func listFeatureFlags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	project := h.Item.(ldapi.Project)
	projectKey := project.Key

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_feature_flag.listFeatureFlags", "connection_error", err)
		return nil, err
	}

	params := client.FeatureFlagsApi.GetFeatureFlags(ctx, projectKey)

	count := 0

	for {
		flags, _, err := params.Execute()
		if err != nil {
			plugin.Logger(ctx).Error("launchdarkly_feature_flag.listFeatureFlags", "api_error", err)
			return nil, err
		}

		for _, flag := range flags.Items {
			d.StreamListItem(ctx, launchdarklyFeatureFlag{flag, projectKey})
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count += len(flags.Items)
		if count >= int(flags.GetTotalCount()) {
			break
		}
		params.Offset(int64(count))
	}
	return nil, nil
}

func getFeatureFlag(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_feature_flag.getFeatureFlag", "connection_error", err)
		return nil, err
	}

	projectKey := d.EqualsQualString("project_key")
	featureFlagKey := d.EqualsQualString("key")

	flag, _, err := client.FeatureFlagsApi.GetFeatureFlag(ctx, projectKey, featureFlagKey).Execute()

	if err != nil {
		logger.Error("launchdarkly_feature_flag.getFeatureFlag", "api_error", err)
		return nil, err
	}

	return launchdarklyFeatureFlag{*flag, projectKey}, nil
}
