---
title: "Steampipe Table: launchdarkly_project - Query LaunchDarkly Projects using SQL"
description: "Allows users to query LaunchDarkly Projects, specifically returning details about each project, including its name, key, and environment details."
---

# Table: launchdarkly_project - Query LaunchDarkly Projects using SQL

A LaunchDarkly Project is a fundamental organizational unit within the LaunchDarkly system. It is used to group related feature flags together. Each project contains environments that control the targeting rules for each feature flag.

## Table Usage Guide

The `launchdarkly_project` table provides insights into projects within LaunchDarkly. As a software engineer, explore project-specific details through this table, including its name, key, and associated environment details. Utilize it to understand the organization and control of feature flags within each project.

## Examples

### Basic info
Explore the basic information of your projects in LaunchDarkly to understand their unique identifiers and keys. This can aid in project management and tracking.Explore the basic information of projects in LaunchDarkly to understand their unique identifiers and associated keys. This can help in managing and organizing projects effectively.


```sql
select
  name,
  id,
  key
from
  launchdarkly_project;
```

### Get the environment details of a project
Explore the environmental aspects of a specific project to gain insights into its unique identifiers and associated details. This can be particularly useful for project management, to ensure the correct environment is being used for each project.Explore the environment details associated with a particular project. This is useful in understanding the project's settings and configurations, especially when managing multiple projects or environments.


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
Explore projects that have been configured for default availability on the client-side via a mobile key. This allows for the identification of projects that are accessible on mobile devices, facilitating mobile-based operations and activities.Explore which projects are configured for default client-side availability via a mobile key. This can be useful to identify projects that are optimized for mobile access, enhancing user experience and engagement.


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
This query can be used to gain insights into the default settings of a project in LaunchDarkly, such as the temporary flag, the true and false display names, and the client-side availability. This can be particularly useful for project managers or developers who want to understand the project's default configurations and make informed decisions about potential changes.Explore the default settings of a project's flags to understand their configuration and usage. This can help in managing and optimizing the project's feature flags.


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