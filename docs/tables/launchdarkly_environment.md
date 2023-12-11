---
title: "Steampipe Table: launchdarkly_environment - Query LaunchDarkly Environments using SQL"
description: "Allows users to query Environments in LaunchDarkly, specifically providing insights into environment-specific details such as keys, names, and color labels."
---

# Table: launchdarkly_environment - Query LaunchDarkly Environments using SQL

An Environment in LaunchDarkly is a workspace for teams to manage feature flag rollouts. Each environment is isolated from others, allowing teams to manage their feature flags separately. Environments are typically correlated with a phase of the deployment pipeline, such as development, staging, or production.

## Table Usage Guide

The `launchdarkly_environment` table provides insights into environments within LaunchDarkly. As a DevOps engineer, explore environment-specific details through this table, including keys, names, and color labels. Utilize it to uncover information about environments, such as their current state, associated tags, and the configuration of their default TTLs.

## Examples

### Basic info
Explore which settings are associated with different environments in LaunchDarkly to understand how the system is configured and to aid in troubleshooting or optimization efforts.Explore the basic information related to various environments in LaunchDarkly to understand their unique identifiers, associated projects, and color coding. This can help in managing and organizing different environments effectively.


```sql+postgres
select
  name,
  id,
  key,
  color,
  project_key
from
  launchdarkly_environment;
```

```sql+sqlite
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
Discover the segments that showcase the correlation between different environments and projects. This can assist in understanding how various projects are distributed across different environments, aiding in efficient project management.Discover the segments that detail the association between project and environment in LaunchDarkly. This is useful for understanding the structure and organization of different environments within specific projects.


```sql+postgres
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

```sql+sqlite
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
Explore which environments are operating in secure mode for enhanced safety and data protection. This is useful for assessing the security posture of your systems.Discover the segments that are operating in secure mode to ensure enhanced security and privacy. This is particularly beneficial in identifying and managing the environments that prioritize data protection.


```sql+postgres
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  secure_mode;
```

```sql+sqlite
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  secure_mode = 1;
```

### List the environments that have tracking enabled for new flags
Discover the segments where tracking is enabled for new flags, providing a useful resource for monitoring feature flag usage and understanding user interactions.Discover the segments where tracking is enabled for new flags. This is useful to ensure that you are collecting data from all the relevant environments for your feature flags.


```sql+postgres
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  default_track_events;
```

```sql+sqlite
select
  name as environment_name,
  id as environment_id,
  key as environment_key
from
  launchdarkly_environment
where
  default_track_events = 1;
```

### List the approval settings for environments
Determine the areas in which approval settings for various environments are applied, focusing on the requirements and permissions for changes and approvals. This is useful for understanding and managing the approval process within your project's environments.Gain insights into the approval process for different environments within a project. This query can be used to understand the approval requirements, such as minimum approvals needed and whether users can review their own requests or apply declined changes, providing a clear overview of control and permissions.

```sql+postgres
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

```sql+sqlite
select
  name as environment_name,
  id as environment_id,
  project_key,
  json_extract(approval_settings, '$.required') as approval_required,
  json_extract(approval_settings, '$.bypassApprovalsForPendingChanges') as approval_for_pending_changes,
  json_extract(approval_settings, '$.minNumApprovals') as minimum_approvals,
  json_extract(approval_settings, '$.canReviewOwnRequest') as review_own_request,
  json_extract(approval_settings, '$.canApplyDeclinedChanges') as apply_declined_changes,
  json_extract(approval_settings, '$.serviceKind') as service_kind,
  json_extract(approval_settings, '$.serviceConfig') as service_configuration,
  json_extract(approval_settings, '$.requiredApprovalTags') as required_approval_for_flags_with_tags
from
  launchdarkly_environment;
```