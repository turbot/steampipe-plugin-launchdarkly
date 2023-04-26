# Table: launchdarkly_access_token

Access tokens are long-term credentials for an user. The personal access token or service token are used to authenticate with LaunchDarkly.

## Examples

### Basic info

```sql
select
  name,
  id,
  creation_date,
  owner_id,
  role,
  last_used
from
  launchdarkly_access_token;
```

### List of access tokens with their member information and date of creation

```sql
select
  name,
  id,
  member ->> '_id' as member_id,
  member ->> 'email' as member_email_id,
 (member ->> 'firstName') || ' ' || (member ->> 'lastName')as member_name,
  member ->> 'role' as member_role,
  creation_date
from
  launchdarkly_access_token;
```

### List the access tokens that have been created in the last 30 days

```sql
select
  name,
  id,
  creation_date,
  owner_id,
  role
from
  launchdarkly_access_token
where
  creation_date >= now() - interval '30' day;
```

### List of access tokens which haven't been used in the last 30 days

```sql
select
  id,
  name,
  last_used
from
  launchdarkly_access_token
where
  last_used <= now() - interval '30' day;
```

### Access key count by member name

```sql
select
  (member ->> '_id') as member_id,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName')as member_name,
  count (id) as access_key_count
from
  launchdarkly_access_token
group by
  member,
  (member ->> '_id');
```

### Get the details of an access token with read-only permission

```sql
select
  name,
  id,
  creation_date,
  owner_id,
  role,
  last_used
from
  launchdarkly_access_token
where
  role = 'reader';
```