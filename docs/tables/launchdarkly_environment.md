# Table: launchdarkly_environment

Environments allow you to maintain separate rollout rules in different contexts, from local development to QA, staging, and production. With the LaunchDarkly Environments API, you can programmatically create, delete, and update environments.

## Examples

### Basic info

```sql
select
  name,
  id,
  key,
  color,
  project_key
from
  launchdarkly_environment;
```

### List the environment details of a project

```sql
select
  e.name as environment_name,
  e.id as environment_id,
  e.key as environment_key,
  p.name as project_name,
  p.id as project_id
from
  launchdarkly_environment as e
  left join launchdarkly_project as p on p.key = e.project_key;
```

### List the environments that run in secure mode

```sql
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  secure_mode;
```

### List the environments that have tracking enabled for new flags

```sql
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  default_track_events;
```

### List the approval settings for environments

```sql
select
  name as environment_name,
  id as environment_id,
  project_key,
  approval_settings ->> 'required' as approval_required,
  approval_settings ->> 'bypassApprovalsForPendingChanges' as approval_for_pending_changes,
  approval_settings ->> 'minNumApprovals' as minimum_approvals,
  approval_settings ->> 'canReviewOwnRequest' as review_own_request,
  approval_settings ->> 'canApplyDeclinedChanges' as apply_declined_changes,
  approval_settings ->> 'serviceKind' as service_kind,
  approval_settings ->> 'serviceConfig' as service_configuration,
  approval_settings ->> 'requiredApprovalTags' as required_approval_for_flags_with_tags
from
  launchdarkly_environment;
```
