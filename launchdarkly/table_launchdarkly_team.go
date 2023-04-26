package launchdarkly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyTeam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_team",
		Description: "Fetch a list of all teams.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "filter", Require: plugin.Optional},
				{Name: "expand", Require: plugin.Optional},
			},
			Hydrate: listTeams,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key"),
			Hydrate: getTeam,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A human-friendly name for the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The team key.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A description for the access token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "access",
				Description: "Defines the access levels designated to the team members.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "creation_date",
				Description: "Creation date of the team.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_modified",
				Description: "Last modified date and team.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModified").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "links",
				Description: "The location and content type of related resources.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "Custom roles assigned to the team.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "idp_synced",
				Description: "Whether the team has been synced with an external identity provider (IdP). Team sync is available to customers on an Enterprise plan.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "members",
				Description: "Team member details.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "projects",
				Description: "Project details associated with the team.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "maintainers",
				Description: "Team maintainer details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "version",
				Description: "The team version.",
				Type:        proto.ColumnType_INT,
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

func listTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_team.listTeams", "connection_error", err)
		return nil, err
	}

	params := client.TeamsApi.GetTeams(ctx)

	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)

		maxTeams := int64(20)

		if limit < int32(maxTeams) {
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
		teams, _, err := params.Execute()
		if err != nil {
			logger.Error("launchdarkly_team.listTeams", "api_error", err)
			return nil, err
		}

		for _, item := range teams.Items {
			d.StreamListItem(ctx, item)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count += len(teams.Items)
		if count >= int(*teams.TotalCount) {
			break
		}
		params.Offset(int64(count))
	}
	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	key := d.EqualsQualString("key")

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_team.getTeam", "connection_error", err)
		return nil, err
	}

	team, _, err := client.TeamsApi.GetTeam(ctx, key).Execute()
	if err != nil {
		logger.Error("launchdarkly_team.getTeam", "api_error", err)
		return nil, err
	}

	return team, nil
}
