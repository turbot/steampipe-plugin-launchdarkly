# Table: launchdarkly_account_member

The account members API allows you to invite new members to an account by making a `POST` request to `/api/v2/members`. When you invite a new member to an account, an invitation is sent to the email you provided. Members with "admin" or "owner" roles may create new members, as well as anyone with a "createMember" permission for "member/*".

## Examples

### Basic info

```sql
select
  _id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member;
```

### List the account members created in the last 30 days

```sql
select
  _id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  creation_date >= now() - interval '30' day;
```

### List the members with MFA enabled

```sql
select
  _id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  mfa = 'enabled';
```

### List the verified account members

```sql
select
  _id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  _verified;
```

### List the custom roles assigned to an account member

```sql
select
  _id,
  first_name || last_name as name,
  email,
  customRoles
from
  launchdarkly_account_member;
```

### List the default dashboards that the member has chosen to ignore

```sql
select
  _id,
  first_name || last_name as name,
  email,
  excludedDashboards
from
  launchdarkly_account_member;
```

### List out the team details of an account member

```sql
select
  _id,
  first_name || last_name as name,
  t ->> 'key' as team_key,
  t ->> 'name' as team_name
from
  launchdarkly_account_member,
  jsonb_array_elements(teams) as t;
```

### List the account members that have been inactive for more than 30 days

```sql
select
  _id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  last_seen <= now() - interval '30' day;
```

### List the permissions granted to an account member

```sql
select
  _id,
  first_name || last_name as name,
  p ->> 'actionSet' as action_set,
  p ->> 'actions' as actions,
  p ->> 'resource' as resource
from
  launchdarkly_account_member,
  jsonb_array_elements(permission_grants) as p;
```