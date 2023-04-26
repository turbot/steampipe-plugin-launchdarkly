# Table: launchdarkly_team

A team is a group of members in your LaunchDarkly account. A team can have maintainers who are able to add and remove team members. It also can have custom roles assigned to it that allows shared access to those roles for all team members.

## Examples

### Basic info

```sql
select
  name,
  key,
  creation_date,
  last_modified,
  version
from
  launchdarkly_team;
```

### List the teams that have been created in the last 30 days

```sql
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

### List the maintainer details of a team

```sql
select
  maintainers -> 'totalCount' as maintainer_count,
  i ->> '_id' as maintainer_id,
  (i ->> 'firstName') || ' ' || (i ->> 'lastName')as maintainer_name,
  i ->> 'role' as maintainer_role,
  i ->> 'email' as maintainer_email
from
  launchdarkly_team,
  jsonb_array_elements(maintainers -> 'items') as i;
```

### List the teams that haven't been modified in the last 30 days

```sql
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

### List the teams that has been synced with an external identity provider

```sql
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

### List the project details associated to a team

```sql
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