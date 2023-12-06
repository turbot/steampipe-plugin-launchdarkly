---
title: "Steampipe Table: launchdarkly_feature_flag - Query LaunchDarkly Feature Flags using SQL"
description: "Allows users to query LaunchDarkly Feature Flags, specifically the key, state, and description, providing insights into feature management and rollout status."
---

# Table: launchdarkly_feature_flag - Query LaunchDarkly Feature Flags using SQL

LaunchDarkly is a feature management platform that allows software teams to control feature rollouts, test in production, and perform progressive delivery. It provides a centralized way to manage and control the visibility and rollout of features in your applications. LaunchDarkly helps you mitigate the risk of software releases and improve the speed and reliability of your delivery pipeline.

## Table Usage Guide

The `launchdarkly_feature_flag` table provides insights into feature flags within LaunchDarkly. As a software engineer or product manager, explore feature-specific details through this table, including the key, state, and description. Utilize it to uncover information about features, such as their rollout status, enabling or disabling features, and understanding the impact of feature changes.

## Examples

### Basic info
Explore the characteristics of different feature flags, including their names, keys, versions, creation dates, and types. This information can be useful to understand the various aspects of your feature flags and to track their history over time.Explore the fundamental details of your feature flags, such as their names, versions, and creation dates, to gain a comprehensive overview and better manage your feature releases. This can be particularly useful in scenarios where you need to assess the overall status of your feature flags or identify any issues.

```sql+postgres
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag;
```

```sql+sqlite
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag;
```

### List the feature flags that have been created in the last 30 days
Discover the recently created feature flags within the past month. This is useful for keeping track of new features or changes that have been introduced to your application or service.Discover the segments that have been introduced in the past month. This can be useful for understanding recent changes or additions to your product's functionality.


```sql+postgres
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  creation_date >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  creation_date >= datetime('now', '-30 day');
```

### List out the archived feature flags
Explore which feature flags have been archived to manage your application's features more effectively. This could be useful in understanding the evolution of your application's features and make informed decisions about future development.Explore which feature flags have been archived in LaunchDarkly to better manage your flag inventory and understand the evolution of your feature rollout strategies.

```sql+postgres
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  archived;
```

```sql+sqlite
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  archived = 1;
```

### List out all the temporary feature flags
Discover the segments that utilize temporary feature flags, allowing for a better understanding of how these are used in your application and potentially highlighting areas for optimization or risk. This can be especially useful in managing and tracking the use of feature flags for testing or temporary features.Explore which feature flags are temporary in your LaunchDarkly setup. This can help manage and clean up temporary flags that are no longer needed, improving system efficiency.

```sql+postgres
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  temporary;
```

```sql+sqlite
select
  name,
  key,
  version,
  creation_date,
  kind
from
  launchdarkly_feature_flag
where
  temporary = 1;
```

### List out the maintainer details of a feature flag
Discover the individuals responsible for maintaining specific feature flags. This is beneficial for understanding who to contact regarding changes or issues with a particular feature flag.Explore which feature flags are being managed by whom, to identify the individuals responsible for specific functions and their roles. This is useful for understanding accountability and communication paths within a team managing feature flags.

```sql+postgres
select
  name,
  key,
  version,
  maintainer ->> '_id' as maintainer_id,
  maintainer ->> 'email' as maintainer_email_id,
  (maintainer ->> 'firstName') || ' ' || (maintainer ->> 'lastName')as maintainer_name,
  maintainer ->> 'role' as maintainer_role
from
  launchdarkly_feature_flag;
```

```sql+sqlite
select
  name,
  key,
  version,
  json_extract(maintainer, '$._id') as maintainer_id,
  json_extract(maintainer, '$.email') as maintainer_email_id,
  (json_extract(maintainer, '$.firstName') || ' ' || json_extract(maintainer, '$.lastName')) as maintainer_name,
  json_extract(maintainer, '$.role') as maintainer_role
from
  launchdarkly_feature_flag;
```