## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#19](https://github.com/turbot/steampipe-plugin-launchdarkly/pull/19))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#17](https://github.com/turbot/steampipe-plugin-launchdarkly/pull/17))
- Recompiled plugin with Go version `1.21`. ([#17](https://github.com/turbot/steampipe-plugin-launchdarkly/pull/17))

## v0.0.1 [2023-05-19]

_What's new?_

- New tables added
  - [launchdarkly_access_token](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_access_token)
  - [launchdarkly_account_member](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_account_member)
  - [launchdarkly_audit_log](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_audit_log)
  - [launchdarkly_environment](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_environment)
  - [launchdarkly_feature_flag](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_feature_flag)
  - [launchdarkly_project](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_project)
  - [launchdarkly_team](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables/launchdarkly_team)
