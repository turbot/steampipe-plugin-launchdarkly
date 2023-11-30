---
title: "Steampipe Table: launchdarkly_account_member - Query LaunchDarkly Account Members using SQL"
description: "Allows users to query LaunchDarkly Account Members, providing insights into members' roles, permissions, and associated metadata."
---

# Table: launchdarkly_account_member - Query LaunchDarkly Account Members using SQL

A LaunchDarkly Account Member is a user who has access to the LaunchDarkly platform. Each member has specific roles and permissions that dictate what they can access and modify within the platform. This includes access to feature flags, user targeting rules, and other resources.

## Table Usage Guide

The `launchdarkly_account_member` table provides insights into the members within LaunchDarkly. As a DevOps engineer, you can explore member-specific details through this table, including roles, permissions, and associated metadata. Use it to uncover information about members, such as their assigned roles, access permissions, and other relevant details.

## Examples

### Basic info
Discover the segments that consist of your account members on LaunchDarkly, including their roles and contact details. This can be useful to understand the distribution of roles and responsibilities within your team.Discover the segments that comprise your LaunchDarkly account members, including their roles and contact details. This can be useful for conducting an audit or understanding team composition.


```sql
select
  id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member;
```

### List the account members created in the last 30 days
Discover the recent additions to your account by identifying members who have been added in the last 30 days. This allows you to stay updated on the newest members and their roles within your organization.Explore which account members were added in the past month. This is useful for tracking recent changes in team composition or user access.


```sql
select
  id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  creation_date >= now() - interval '30' day;
```

### List the acount members with MFA enabled
Explore which account members have the added security of multi-factor authentication (MFA) enabled. This is beneficial for assessing the security measures in place and identifying any potential vulnerabilities.Discover the segments of your team who have enabled multi-factor authentication (MFA) for added security. This can help in identifying areas where security measures are being actively implemented.


```sql
select
  id,
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
Explore which account members have been verified to gain insights into the user base. This can be useful for understanding the proportion of verified users, which can inform decision-making in areas like security and user engagement strategies.Explore the list of verified members within a given account, along with their roles and contact information. This is useful for account management and ensuring all verified members have the appropriate access and roles.


```sql
select
  id,
  first_name || last_name as name,
  role,
  email,
  creation_date
from
  launchdarkly_account_member
where
  verified;
```

### List the custom roles assigned to an account member
Gain insights into the custom roles assigned to each account member, which can help in managing user permissions and access within the system. This is particularly useful in large teams where role-based access control is implemented.Explore which custom roles are assigned to specific individuals within an account. This can help in managing user permissions and access control, ensuring that each member has the appropriate roles for their tasks.


```sql
select
  id,
  first_name || last_name as name,
  email,
  custom_roles
from
  launchdarkly_account_member;
```

### List the default dashboards that the member has chosen to ignore
Explore which default dashboards a member has chosen to ignore to better understand user preferences and tailor the platform experience accordingly.Uncover the details of account members who have chosen to ignore default dashboards. This can be useful in understanding their preferences and improving user experience.


```sql
select
  id,
  first_name || last_name as name,
  email,
  exclude_dashboards
from
  launchdarkly_account_member;
```

### List out the team details of an account member
Gain insights into the team affiliations of account members. This query is particularly useful when you need to understand the distribution of members across different teams within an account.Explore the team details associated with an account member to understand their role and involvement. This is particularly useful in managing user permissions and roles within an organization.


```sql
select
  id,
  first_name || last_name as name,
  t ->> 'key' as team_key,
  t ->> 'name' as team_name
from
  launchdarkly_account_member,
  jsonb_array_elements(teams) as t;
```

### List the account members that have been inactive for more than 30 days
Discover the members of your account who have not been active for over a month. This is useful for understanding user engagement and identifying potential areas for improvement in user retention strategies.Determine the areas in which account members have been inactive for over a month. This can help in identifying users who may need re-engagement efforts or account clean-up.


```sql
select
  id,
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
Explore which actions are permitted to a specific account member. This can be useful in managing access control and ensuring only appropriate privileges are granted.Determine the specific permissions assigned to a member of your account. This is beneficial in managing access control, ensuring each member has the appropriate permissions for their role.


```sql
select
  id,
  first_name || last_name as name,
  p ->> 'actionSet' as action_set,
  p ->> 'actions' as actions,
  p ->> 'resource' as resource
from
  launchdarkly_account_member,
  jsonb_array_elements(permission_grants) as p;
```