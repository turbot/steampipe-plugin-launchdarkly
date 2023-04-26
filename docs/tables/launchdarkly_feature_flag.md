# Table: launchdarkly_feature_flag

Thisg table provides information about feature flags used by your application. Feature flags are used to control percentage rollouts, target specific contexts, or toggle off a feature programmatically. By querying this table, you can view the representation of a feature flag and perform various tasks related to feature management.

## Examples

### Basic info

```sql
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

```sql
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

### List out the archived feature flags

```sql
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

### List out all the temporary feature flags

```sql
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

### List out the maintainer details of a feature flag

```sql
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