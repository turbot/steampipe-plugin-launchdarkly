![image](https://hub.steampipe.io/images/plugins/turbot/launchdarkly-social-graphic.png)

# LaunchDarkly Plugin for Steampipe

Use SQL to query models, completions and more from LaunchDarkly.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/launchdarkly)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/launchdarkly/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-launchdarkly/issues)

## Quick start

### Install

Download and install the latest LaunchDarkly plugin:

```bash
steampipe plugin install launchdarkly
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/launchdarkly#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/launchdarkly#configuration).

Configure your account details in `~/.steampipe/config/launchdarkly.spc`:

```hcl
connection "launchdarkly" {
  plugin = "launchdarkly"

  # Authentication information
  access_token = "api-dd8ce121-cd11-401c-be02-322b7362111d"
}
```

Or through environment variables:

```sh
LAUNCHDARKLY_ACCESS_TOKEN=api-dd8ce121-cd11-401c-be02-322b7362111d
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-launchdarkly.git
cd steampipe-plugin-launchdarkly
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/launchdarkly.spc
```

Try it!

```
steampipe query
> .inspect launchdarkly
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-launchdarkly/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [LaunchDarkly Plugin](https://github.com/turbot/steampipe-plugin-launchdarkly/labels/help%20wanted)
