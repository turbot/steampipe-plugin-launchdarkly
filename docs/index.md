---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/launchdarkly.svg"
brand_color: "#191717"
display_name: "LaunchDarkly"
short_name: "launchdarkly"
description: "Steampipe plugin to query projects, teams, metrics, flags and more from LaunchDarkly."
og_description: "Query LaunchDarkly with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/launchdarkly-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# LaunchDarkly + Steampipe

[LaunchDarkly](https://launchdarkly.com) is a feature management platform that enables software teams to build better software faster by safely serving and controlling software features in production.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List feature flags in your LaunchDarkly account:

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

```
+----------------------+----------------------+---------+---------------------------+--------------+
| name                 | key                  | version | creation_date             | kind         |
+----------------------+----------------------+---------+---------------------------+--------------+
| dark-light-mode-node | dark-light-mode-node | 1       | 2023-04-24T10:28:26+05:30 | boolean      |
| first-flag-test      | first-flag-test      | 1       | 2023-04-25T15:36:33+05:30 | multivariate |
| alphabet-flag        | num-flag             | 1       | 2023-04-24T11:20:39+05:30 | multivariate |
| hello-world          | go-world             | 2       | 2023-04-25T16:46:23+05:30 | boolean      |
+----------------------+----------------------+---------+---------------------------+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/launchdarkly/tables)**

## Quick start

### Install

Download and install the latest LaunchDarkly plugin:

```sh
steampipe plugin install launchdarkly
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | LaunchDarkly requires an [Access token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#creating-api-access-tokens) for all requests.                                                                |
| Permissions | Access tokens have the same permissions as the user who creates them, and if the user permissions change, the Access token permissions also change.                                                         |
| Radius      | Each connection represents a single LaunchDarkly Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/launchdarkly.spc`)<br />2. Credentials specified in environment variables, e.g., `LAUNCHDARKLY_ACCESS_TOKEN`. |

### Configuration

Installing the latest launchdarkly plugin will create a config file (`~/.steampipe/config/launchdarkly.spc`) with a single connection named `launchdarkly`:

Configure your account details in `~/.steampipe/config/launchdarkly.spc`:

```hcl
connection "launchdarkly" {
  plugin = "launchdarkly"

  # `access_token`: LaunchDarkly Access Token. (Required)
  # Generate your Access Token per https://docs.launchdarkly.com/home/account-security/api-access-tokens#creating-api-access-tokens
  # This can also be set via the `LAUNCHDARKLY_ACCESS_TOKEN` environment variable.  
  # access_token = "api-dd8ce121-cd11-401c-be02-322b7362111d"
}
```

Alternatively, you can also use the standard LaunchDarkly environment variables to obtain credentials **only if other arguments (`access_token`) are not specified** in the connection:

```sh
export LAUNCHDARKLY_ACCESS_TOKEN=api-dd8ce121-cd11-401c-be02-322b7362111d
```

)