# Table: launchdarkly_project

Projects allow you to manage multiple different software projects under one LaunchDarkly account. Each project has its own unique set of environments and feature flags.

## Examples

### Basic info

```sql
select
  name,
  id,
  key
from
  launchdarkly_project;
```

### List the project details in an environment

```sql
select
  p.name as project_name,
  p.id as project_id,
  p.key as project_key,
  e.name as environment_name,
  e.id as environment_id
from
  launchdarkly_project as p
  left join launchdarkly_environment as e on p.key = e.project_key;
```

### List the projects that have default client side availaility using mobile key enabled

```sql
select
  name as project_name,
  id as project_id,
  key as project_key
from
  launchdarkly_project
where
  default_client_side_availability ->> 'usingMobileKey' = 'true';
```

