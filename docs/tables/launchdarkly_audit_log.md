---
title: "Steampipe Table: launchdarkly_audit_log - Query LaunchDarkly Audit Logs using SQL"
description: "Allows users to query LaunchDarkly Audit Logs, specifically the detailed history of changes made to any resources and who made those changes, providing insights into resource management and potential security risks."
---

# Table: launchdarkly_audit_log - Query LaunchDarkly Audit Logs using SQL

LaunchDarkly Audit Logs are a feature within the LaunchDarkly service that track and log every change made to any resources within the platform. It provides detailed historical data about who made changes, what changes were made, and when those changes occurred. This feature is crucial for maintaining security, accountability, and understanding the evolution of resources over time.

## Table Usage Guide

The `launchdarkly_audit_log` table provides insights into the detailed history of changes made to any resources within the LaunchDarkly service. As a System Administrator or Security Specialist, explore change-specific details through this table, including who made the changes, what changes were made, and the timestamp of those changes. Utilize it to monitor resource management, identify potential security risks, and maintain accountability for changes made within the platform.

## Examples

### Basic info
Explore the audit logs to gain insights into various actions performed within your account, such as who made changes and when, to maintain security and accountability.Explore which changes have been made in your LaunchDarkly settings by identifying instances where the audit log has been updated. This allows you to monitor and review the configuration changes for better understanding and control over your feature flag management.


```sql+postgres
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

```sql+sqlite
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
Assess the most recent activities within an account to understand the actions taken and by whom. This can provide valuable insights into user behavior and activity trends.Discover the latest activities related to your account, such as who performed them and when. This is useful to monitor account activity and track recent changes.


```sql+postgres
select
  date,
  audit_log_title,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName') as actor_display_name,
  name,
  title_verb
from
  launchdarkly_audit_log
order by
  date desc
limit
  10;
```

```sql+sqlite
select
  date,
  audit_log_title,
  (json_extract(member, '$.firstName')) || ' ' || (json_extract(member, '$.lastName')) as actor_display_name,
  name,
  title_verb
from
  launchdarkly_audit_log
order by
  date desc
limit
  10;
```

### List the most common actors
Discover the most frequently occurring individuals in your audit logs to identify who is making the most changes. This can be useful for understanding user behavior, ensuring compliance, and identifying potential security risks.Gain insights into the most frequently appearing actors within a specific account. This is useful for identifying the most active members and understanding user behavior within the account.


```sql+postgres
select
  account_id,
  name,
  (member ->> 'firstName') || ' ' || (member ->> 'lastName') as actor_display_name,
  count(*)
from
  launchdarkly_audit_log
group by
  account_id,
  actor_display_name,
  name
order by
  count desc;
```

```sql+sqlite
select
  account_id,
  name,
  (json_extract(member, '$.firstName')) || ' ' || (json_extract(member, '$.lastName')) as actor_display_name,
  count(*)
from
  launchdarkly_audit_log
group by
  account_id,
  actor_display_name,
  name
order by
  count(*) desc;
```

### List out the most common actions
Gain insights into the most frequently occurring actions within your audit log to better understand user behavior and system usage.Explore which actions are most commonly performed in your system. This insight can be used to understand user behavior and optimize system design for better efficiency.


```sql+postgres
select
  title_verb,
  name,
  count(*)
from
  launchdarkly_audit_log
group by
  title_verb,
  name
order by
  count desc;
```

```sql+sqlite
select
  title_verb,
  name,
  count(*)
from
  launchdarkly_audit_log
group by
  title_verb,
  name
order by
  count(*) desc;
```

### Find all project creation events
Discover the instances where new projects were initiated. This query is useful to track the creation of new projects, providing insights into who initiated them and when, which can aid in project management and accountability.Determine the instances where new projects were initiated. This allows you to track project creation activities, providing valuable insights into team productivity and workflow patterns.


```sql+postgres
select
  launchdarkly_audit_log.name,
  launchdarkly_audit_log.id,
  launchdarkly_audit_log.date,
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

```sql+sqlite
select
  launchdarkly_audit_log.name,
  launchdarkly_audit_log.id,
  launchdarkly_audit_log.date,
  (json_extract(member, '$.firstName')) || ' ' || (json_extract(member, '$.lastName')) as actor_display_name,
  audit_log_title,
  title_verb
from
  launchdarkly_audit_log,
  json_each(accesses) as a
where
  json_extract(a.value, '$.action') = 'createProject'
order by
  date desc;
```

### List events that occurred over the last five minutes
Explore recent activity by identifying events that have transpired in the last five minutes. This is particularly useful for real-time monitoring and immediate response to changes or anomalies.Explore recent activities by listing events that happened in the last five minutes. This can help in real-time monitoring and quick response to changes or issues.


```sql+postgres
select
  audit_log_title,
  name,
  id,
  date
from
  launchdarkly_audit_log
where
  name = 'audit-log-name'
  and date >= now() - interval '5 minutes';
```

```sql+sqlite
select
  audit_log_title,
  name,
  id,
  date
from
  launchdarkly_audit_log
where
  name = 'audit-log-name'
  and date >= datetime('now', '-5 minutes');
```

### List ordered events that occurred between five to ten minutes ago
Explore events that happened within a specific time frame in the recent past. This is particularly useful for tracking changes and identifying any anomalies or unexpected activity within that period.Explore the sequence of events that took place within a specific timeframe in the past. This is particularly useful for tracking changes or incidents that occurred between five to ten minutes ago, allowing for timely response and action.


```sql+postgres
select
  name,
  audit_log_title,
  id,
  date
from
  launchdarkly_audit_log
where
  name = 'audit-log-name'
  and date between (now() - interval '10 minutes') and (now() - interval '5 minutes')
order by
  date asc;
```

```sql+sqlite
select
  name,
  audit_log_title,
  id,
  date
from
  launchdarkly_audit_log
where
  name = 'audit-log-name'
  and date between (datetime('now', '-10 minutes')) and (datetime('now', '-5 minutes'))
order by
  date asc;
```