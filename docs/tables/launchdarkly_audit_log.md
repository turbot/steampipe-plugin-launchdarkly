# Table: launchdarkly_audit_log

The audit log contains a record of all the changes made to any resource in the system. You can filter the audit log by timestamps, or use a custom policy to select which entries to receive.

## Examples

### Basic info

```sql
select
  name,
  id,
  account_id,
  audit_log_title,
  kind,
  date
from
  launchdarkly_audit_log;
```

### List the most recent actions for the account

```sql
select
  date,
  audit_log_title,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName') as actor_display_name,
  title_verb
from
  launchdarkly_audit_log
order by
  date desc
limit
  10;
```

### List the most common actors

```sql
select
  account_id,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName') as actor_display_name,
  count(*)
from
  launchdarkly_audit_log
group by
  account_id,
  actor_display_name
order by
  count desc;
```

### List out the most common actions

```sql
select
  title_verb,
  count(*)
from
  launchdarkly_audit_log
group by
  title_verb
order by
  count desc;
```

### Find all project creation events

```sql
select
  date,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName') as actor_display_name,
  audit_log_title,
  title_verb
from
  launchdarkly_audit_log,
  jsonb_array_elements(accesses) as a
where
  a ->> 'action' = 'createProject'
order by
  date desc;
```