package launchdarkly

import (
	"context"

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
			Hydrate: listProjects,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate: getProject,
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
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
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

	var params = client.ProjectsApi.GetProjects(ctx)
	offsetparam := d.EqualsQuals["offset"].GetInt64Value()
	logger.Trace("Apple, checking output of offset", offsetparam)
	if d.EqualsQuals["offset"].GetInt64Value() != 0 {
		params.Offset(d.EqualsQuals["offset"].GetInt64Value())
	}

	if d.EqualsQuals["limit"].GetInt64Value() != 0 {
		params.Limit(d.EqualsQuals["limit"].GetInt64Value())
	} else {
		params.Limit(20)
	}

	if d.EqualsQuals["filter"].GetStringValue() != "" {
		params.Filter(d.EqualsQuals["filter"].GetStringValue())
	}

	if d.EqualsQuals["sort"].GetStringValue() != "" {
		params.Sort(d.EqualsQuals["sort"].GetStringValue())
	}

	if d.EqualsQuals["expand"].GetStringValue() != "" {
		params.Expand(d.EqualsQuals["expand"].GetStringValue())
	}

	count := 0

	for {
		projects, _, err := params.Execute()
		checkTotal := projects.GetTotalCount()
		logger.Trace("Check total count error", checkTotal)
		if err != nil {
			plugin.Logger(ctx).Error("listProject", "api_error", err)
			return nil, err
		}

		for _, item := range projects.Items {
			d.StreamListItem(ctx, item)

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
	id := d.EqualsQualString("id")
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_project.getProject", "connection_error", err)
		return nil, err
	}

	token, _, err := client.ProjectsApi.GetProject(ctx, id).Execute()
	if err != nil {
		logger.Error("launchdarkly_project.getProject", "api_error", err)
		return nil, err
	}

	return token, nil
}