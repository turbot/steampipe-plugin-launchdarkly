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

### Get the environment details of a project

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

### List the projects that have default client side availaility using mobile key

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

### List the flag defaults of a project

```sql
select
  id as project_id,
  key as project_key,
  name as project_name,
  flag_defaults ->> 'temporary' as temporary_flag,
  flag_defaults -> 'booleanDefaults' ->> 'trueDisplayName' as true_display_name,
  flag_defaults -> 'booleanDefaults' ->> 'falseDisplayName' as false_display_name,
  flag_defaults -> 'booleanDefaults' ->> 'onVariation' as on_variation,
  flag_defaults -> 'booleanDefaults' ->> 'offVariation' as off_variation,
  flag_defaults -> 'defaultClientSideAvailability' ->> 'usingEnvironmentId' as client_side_availability_using_environment_id,
  flag_defaults -> 'defaultClientSideAvailability' ->> 'usingMobileKey' as client_side_availability_using_mobile_key
from
  launchdarkly_project;
```