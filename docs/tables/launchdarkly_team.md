---
title: "Steampipe Table: launchdarkly_team - Query LaunchDarkly Teams using SQL"
description: "Allows users to query Teams in LaunchDarkly, providing access to team information and member details."
---

# Table: launchdarkly_team - Query LaunchDarkly Teams using SQL

LaunchDarkly is a feature management platform that empowers all teams to safely deliver and control software through feature flags. Teams in LaunchDarkly represent a collection of users, which can be used to manage feature flag targeting and permissions. It provides a centralized way to manage team members and their associated roles and permissions within the LaunchDarkly platform.

## Table Usage Guide

The `launchdarkly_team` table provides insights into Teams within LaunchDarkly. As an administrator or team lead, explore team-specific details through this table, including team keys, names, and associated member details. Utilize it to manage and monitor team permissions, roles, and feature flag targeting in the LaunchDarkly platform.

## Examples

### Basic info
Explore which elements within your team have been recently modified on LaunchDarkly. This can help keep track of changes over time and maintain an updated version history.Explore the details of your team's activity on LaunchDarkly by identifying when a particular feature was created, last modified, and its current version. This can help in understanding the progression and changes in the team's feature development over time.


```sql+postgres
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team;
```

```sql+sqlite
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team;
```

### List teams that have been created in the last 30 days
Discover the teams that have been recently formed within the past month. This can be beneficial for understanding the growth and expansion of your organization's teams.Discover the teams that have been recently formed within the last month. This is useful for tracking new team formation and understanding recent changes in your organization's structure.


```sql+postgres
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  creation_date >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  creation_date >= datetime('now', '-30 day');
```

### List the maintainer details of a team
Determine the number of maintainers for a team and their respective details, such as their ID, name, role, and email. This is useful to understand the team's structure and the individuals responsible for its maintenance.This query is useful for gaining insights into the number and details of maintainers associated with a particular team. It can help in understanding the team's structure and the roles of different maintainers, which is essential for effective team management.


```sql+postgres
select
  name,
  description,
  key,
  maintainers -> 'totalCount' as maintainer_count,
  i ->> '_id' as maintainer_id,
  (i ->> 'firstName') || ' ' || (i ->> 'lastName')as maintainer_name,
  i ->> 'role' as maintainer_role,
  i ->> 'email' as maintainer_email
from
  launchdarkly_team,
  jsonb_array_elements(maintainers -> 'items') as i;
```

```sql+sqlite
select
  name,
  description,
  key,
  json_extract(maintainers, '$.totalCount') as maintainer_count,
  json_extract(i.value, '$._id') as maintainer_id,
  (json_extract(i.value, '$.firstName') || ' ' || json_extract(i.value, '$.lastName')) as maintainer_name,
  json_extract(i.value, '$.role') as maintainer_role,
  json_extract(i.value, '$.email') as maintainer_email
from
  launchdarkly_team,
  json_each(json_extract(maintainers, '$.items')) as i;
```

### List the teams that haven't been modified in the last 30 days
Explore which teams in LaunchDarkly have remained unchanged over the past month. This is useful for identifying potential areas of stagnation or lack of activity within your organization.Discover the teams that have remained unchanged over the past month. This can help identify areas of stability or inactivity within the organization.


```sql+postgres
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  last_modified <= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  last_modified <= datetime('now', '-30 day');
```

### List the teams that have been synced with an external identity provider
Explore the teams that have been synchronized with an external identity provider to understand the history and current status of integration. This is useful for managing and auditing your team's identity management practices.Explore which teams have been synchronized with an external identity provider. This is useful for understanding the integration status of your teams and ensuring data consistency across platforms.


```sql+postgres
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  idp_synced;
```

```sql+sqlite
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team
where
  idp_synced = 1;
```

### List the project details associated to a team
Discover the segments that link various project details to a specific team. This can be particularly useful in assessing the team's workload and understanding the scope of their projects.Explore which projects are linked to a specific team. This can help in efficiently managing resources by understanding the distribution of projects across different teams.


```sql+postgres
select
  name as team_name,
  key as team_key,
  projects ->> 'totalCount' as projects_count,
  i ->> 'key' as project_key,
  i ->> 'name' as project_name
from
  launchdarkly_team,
  jsonb_array_elements(projects -> 'items') as i;
```

```sql+sqlite
select
  name as team_name,
  key as team_key,
  json_extract(projects, '$.totalCount') as projects_count,
  json_extract(i.value, '$.key') as project_key,
  json_extract(i.value, '$.name') as project_name
from
  launchdarkly_team,
  json_each(projects, '$.items') as i;
```