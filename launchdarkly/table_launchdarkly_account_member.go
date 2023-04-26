package launchdarkly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyAccountMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_account_member",
		Description: "Fetch a list of all account members.",
		List: &plugin.ListConfig{
			Hydrate: listAccountMembers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAccountMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The member's ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "first_name",
				Description: "First name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_name",
				Description: "Last Name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "The member's built-in role. If the member has no custom roles, this role will be in effect.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The member's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "verified",
				Description: "Whether the member's email address has been verified.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "pending_invite",
				Description: "Whether the member has a pending invitation.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "pending_email",
				Description: "The member's email address before it has been verified, for accounts where email verification is required.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "custom_roles",
				Description: "The set of custom roles (as keys) assigned to the member.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "mfa",
				Description: "Whether multi-factor authentication is enabled for this member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "exclude_dashboards",
				Description: "Default dashboards that the member has chosen to ignore.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "last_seen",
				Description: "Last seen timestamp of the member.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastSeen").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "teams",
				Description: "Details on the teams this member is assigned to.",
				Type:        proto.ColumnType_JSON,

			},
			{
				Name:        "permission_grants",
				Description: "A list of permission grants. Permission grants allow a member to have access to a specific action, without having to create or update a custom role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "creation_date",
				Description: "Time when the member was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationDate").Transform(transform.UnixMsToTimestamp),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		},
	}
}

//// LIST FUNCTION

func listAccountMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_account_member.listAccountMembers", "connection_error", err)
		return nil, err
	}

	params := client.AccountMembersApi.GetMembers(ctx)

	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)

		maxTeams := int64(20)

		if limit < int32(maxTeams) {
			params = params.Limit(int64(limit))
		}
	}

	count := 0

	for {
		accountMembers, _, err := params.Execute()
		if err != nil {
			logger.Error("launchdarkly_account_member.listAccountMembers", "api_error", err)
			return nil, err
		}

		for _, item := range accountMembers.Items {
			d.StreamListItem(ctx, item)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count += len(accountMembers.Items)
		if count >= int(*accountMembers.TotalCount) {
			break
		}
		params.Offset(int64(count))
	}
	return nil, nil
}

func getAccountMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	memberId := d.EqualsQualString("id")

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_account_member.getAccountMember", "connection_error", err)
		return nil, err
	}

	member, _, err := client.AccountMembersApi.GetMember(ctx, memberId).Execute()
	if err != nil {
		logger.Error("launchdarkly_account_member.getAccountMember", "api_error", err)
		return nil, err
	}

	return member, nil
}
