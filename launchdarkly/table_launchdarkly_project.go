package launchdarkly

import (
	"context"

	ldapi "github.com/launchdarkly/api-client-go/v13"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_project",
		Description: "Fetch a list of all projects.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "expand", Require: plugin.Optional},
			},
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key"),
			Hydrate:    getProject,
		},
		Columns: []*plugin.Column{
			{
				Name:        "links",
				Description: "The location and content type of related resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id",
				Description: "The unique identifier of this project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The key of this project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "include_in_snippet_by_default",
				Description: "A boolean value that indicates whether or not flags created in this project are made available to the client-side JavaScript SDK by default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "default_client_side_availability",
				Description: "A set of boolean values which represent the client side availability.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "name",
				Description: "A friendly name for the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tags",
				Description: "A list of tags for the project.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "environments",
				Description: "Details of the environment associated to the project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "filter",
				Description: "A comma-separated list of filters.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("filter"),
			},
			{
				Name:        "expand",
				Description: "A comma-separated list of properties that can reveal additional information in the response.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("expand"),
			},
			{
				Name:        "flag_defaults",
				Description: "Details of the flag defaults for a specific project.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
				Hydrate:     getFlagDefaultsForProject,
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

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_project.listProjects", "connection_error", err)
		return nil, err
	}

	params := client.ProjectsApi.GetProjects(ctx)

	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)

		maxProjects := int64(20)

		if limit < int32(maxProjects) {
			params = params.Limit(int64(limit))
		}
	}

	if d.EqualsQuals["filter"].GetStringValue() != "" {
		params = params.Filter(d.EqualsQualString("filter"))
	}

	if d.EqualsQuals["expand"].GetStringValue() != "" {
		params = params.Filter(d.EqualsQualString("expand"))
	}

	count := 0

	for {
		projects, _, err := params.Execute()
		if err != nil {
			plugin.Logger(ctx).Error("launchdarkly_project.listProjects", "api_error", err)
			return nil, err
		}

		for _, project := range projects.Items {
			d.StreamListItem(ctx, project)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count += len(projects.Items)
		if count >= int(projects.GetTotalCount()) {
			break
		}
		params.Offset(int64(count))
	}

	return nil, nil

}

func getProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	key := d.EqualsQualString("key")

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_project.getProject", "connection_error", err)
		return nil, err
	}

	project, _, err := client.ProjectsApi.GetProject(ctx, key).Execute()
	if err != nil {
		logger.Error("launchdarkly_project.getProject", "api_error", err)
		return nil, err
	}

	return *project, nil
}

func getFlagDefaultsForProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	key := h.Item.(ldapi.Project).Key
	logger.Trace("Key info", key)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_project.getFlagDefaultsForProject", "connection_error", err)
		return nil, err
	}

	flag, _, err := client.ProjectsApi.GetFlagDefaultsByProject(ctx, key).Execute()
	if err != nil {
		logger.Error("launchdarkly_project.getFlagDefaultsForProject", "api_error", err)
		return nil, err
	}

	return flag, nil
}
