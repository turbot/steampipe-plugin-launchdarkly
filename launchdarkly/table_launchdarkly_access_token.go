package launchdarkly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyAccessToken(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_access_token",
		Description: "Fetch a list of all access tokens.",
		List: &plugin.ListConfig{
			Hydrate: listAccessTokens,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAccessToken,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the access token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier of the access token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner_id",
				Description: "A unique identifier of the owner of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "member_id",
				Description: "A unique identifier of the member of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A description for the access token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_date",
				Description: "Creation date of the access token.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_modified",
				Description: "Last modified date of the access token.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModified").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "role",
				Description: "Built-in role for the token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "token",
				Description: "The token value. When creating or resetting, contains the entire token value. Otherwise, contains the last four characters.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_token",
				Description: "Whether this is a service token or a personal token.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "default_api_version",
				Description: "The default API version for this token.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "last_used",
				Description: "Date and time when the access token was last used.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastUsed").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "member",
				Description: "Summary of the member like email, first name, last name etc.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_role_ids",
				Description: "A list of custom role IDs to use as access limits for the access token.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "inline_role",
				Description: "An array of policy statements, with three attributes: effect, resources, actions. May be used in place of a built-in or custom role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "links",
				Description: "The location and content type of related resources.",
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

//// LIST FUNCTION

func listAccessTokens(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_access_token.listAccessTokens", "connection_error", err)
		return nil, err
	}

	tokens, _, err := client.AccessTokensApi.GetTokens(ctx).Execute()
	if err != nil {
		logger.Error("launchdarkly_access_token.listAccessTokens", "api_error", err)
		return nil, err
	}
	for _, item := range tokens.Items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func getAccessToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_access_token.getAccessToken", "connection_error", err)
		return nil, err
	}

	token, _, err := client.AccessTokensApi.GetToken(ctx, id).Execute()
	if err != nil {
		logger.Error("launchdarkly_access_token.listAccessTokens", "api_error", err)
		return nil, err
	}

	return token, nil
}
