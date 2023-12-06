---
title: "Steampipe Table: launchdarkly_access_token - Query LaunchDarkly Access Tokens using SQL"
description: "Allows users to query LaunchDarkly Access Tokens, specifically to retrieve and inspect the access token details, providing insights into token usage and activities."
---

# Table: launchdarkly_access_token - Query LaunchDarkly Access Tokens using SQL

LaunchDarkly is a feature management platform that enables teams to safely deliver and control software through feature flags. Access Tokens in LaunchDarkly are used to authenticate API calls. They can be scoped to provide either global or project-specific access.

## Table Usage Guide

The `launchdarkly_access_token` table provides insights into access tokens within LaunchDarkly's feature management platform. As a developer or security analyst, explore token-specific details through this table, including scopes, projects, and associated metadata. Utilize it to uncover information about tokens, such as those with global access, the projects associated with each token, and the verification of token activities.

## Examples

### Basic info
Explore the creation and last usage dates, along with the owner and role details, of access tokens to help manage their lifecycle and maintain security.Discover the segments that have been recently used in the LaunchDarkly platform. This allows for an assessment of user activity and role assignment, which can be helpful in understanding platform usage patterns and managing access control.


```sql+postgres
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

```sql+sqlite
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
Explore which access tokens are associated with specific members and when they were created. This is useful for auditing purposes, enabling you to track access and identify any potential security risks.Explore the relationship between access tokens and their associated member information, including their unique identifiers and roles. This can be useful for understanding user activity and permissions, as well as tracking the creation dates of these tokens for security or auditing purposes.


```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  json_extract(member, '$._id') as member_id,
  json_extract(member, '$.email') as member_email_id,
  (json_extract(member, '$.firstName') || ' ' || json_extract(member, '$.lastName')) as member_name,
  json_extract(member, '$.role') as member_role,
  creation_date
from
  launchdarkly_access_token;
```

### List the access tokens that have been created in the last 30 days
Discover the access tokens that have been recently created to understand any potential security risks or unusual activity. This can be useful in maintaining the security of your system by identifying any unauthorized or unexpected tokens.Discover the access tokens that were generated in the past month. This can help you monitor recent activity and manage access control effectively.


```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  creation_date,
  owner_id,
  role
from
  launchdarkly_access_token
where
  creation_date >= datetime('now', '-30 day');
```

### List the access tokens which haven't been used in the last 30 days
Explore which access tokens have been inactive for the past 30 days. This can help in identifying unused or potentially expired tokens, aiding in system clean-up and security measures.Explore which access tokens have remained inactive for the past 30 days. This can be useful for identifying potential security risks or cleaning up unused resources.


```sql+postgres
select
  id,
  name,
  last_used
from
  launchdarkly_access_token
where
  last_used <= now() - interval '30' day;
```

```sql+sqlite
select
  id,
  name,
  last_used
from
  launchdarkly_access_token
where
  last_used <= datetime('now', '-30 day');
```

### Access key count by member name
Determine the number of access keys associated with each member in your LaunchDarkly system. This can help manage and monitor user permissions and security within your organization.Analyze the number of access keys associated with each member to understand their level of system access. This could be useful for auditing purposes or to identify potential security risks.


```sql+postgres
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

```sql+sqlite
select
  json_extract(member, '$._id') as member_id,
  json_extract(member, '$.firstName') || ' ' || json_extract(member, '$.lastName') as member_name,
  count (id) as access_key_count
from
  launchdarkly_access_token
group by
  member,
  json_extract(member, '$._id');
```

### Get the details of access tokens with read-only permission
Explore which access tokens have read-only permissions to understand their usage and ownership details. This can help in maintaining security by ensuring that unauthorized changes aren't being made.Explore which access tokens have been assigned read-only permissions to understand their usage and ownership. This could be useful for auditing purposes, ensuring that only appropriate users have read-only access and identifying any potential security risks.


```sql+postgres
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

```sql+sqlite
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