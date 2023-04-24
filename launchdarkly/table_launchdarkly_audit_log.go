package launchdarkly

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tablelaunchdarklyAuditLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "launchdarkly_audit_log",
		Description: "Fetch a list of all access tokens.",
		List: &plugin.ListConfig{
			Hydrate: listAuditLogs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate: getAuditLog,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource this audit log entry refers to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The ID of the audit log entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "short_description",
				Description: "Shorter version of the change recorded in the audit log entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_id",
				Description: "A unique identifier of the member of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description of the change recorded in the audit log entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date",
				Description: "Date of the audit log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Date").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "accesses",
				Description: "Details on the actions performed and resources acted on in this audit log entry.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "kind",
				Description: "Type of resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comment",
				Description: "Optional comment for the audit log entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "subject",
				Description: "Optional comment for the audit log entry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "member",
				Description: "Summary of the member like email, first name, last name etc.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "token",
				Description: "Access token data representation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "app",
				Description: "Authorized app data representation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
			},
			{
				Name:        "title_verb",
				Description: "The action and resource recorded in this audit log entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "audit_log_title",
				Description: "A description of what occurred, in the 'format member' 'titleVerb' 'target'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Title"),
			},
			{
				Name:        "target",
				Description: "Target resource representation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),	
			},
			{
				Name:        "parent",
				Description: "Parent resource representation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
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

func listAuditLogs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_audit_log.listAuditLogs", "connection_error", err)
		return nil, err
	}

	// Add support for optional quals

	auditLogs, _, err := client.AuditLogApi.GetAuditLogEntries(ctx).Execute()
	if err != nil {
		logger.Error("launchdarkly_audit_log.listAuditLogs", "api_error", err)
		return nil, err
	}
	for _, item := range auditLogs.Items {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func getAuditLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("launchdarkly_audit_log.getAuditLog", "connection_error", err)
		return nil, err
	}

	auditLog, _, err := client.AuditLogApi.GetAuditLogEntry(ctx, id).Execute()
	if err != nil {
		logger.Error("launchdarkly_audit_log.getAuditLog", "api_error", err)
		return nil, err
	}

	return auditLog, nil
}

