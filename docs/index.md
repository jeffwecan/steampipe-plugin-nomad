---
organization: jeffwecan
category: ["public cloud"]
icon_url: "/images/plugins/turbot/nomad.svg"
brand_color: "#00BC7F"
display_name: "Nomad"
short_name: "nomad"
description: "Steampipe plugin for Nomad clusters."
---

# Nomad

The Nomad plugin is used to interact HashiCorp Nomad resources.

## Installation

To download and install the latest nomad plugin:

```bash
$ steampipe plugin install nomad
Installing plugin nomad...
$
```

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files.  Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

### Scope

Each Nomad connection is scoped to a single Nomad cluster, with a single set of credentials.  You may configure multiple Nomad connections if desired, with each connecting to a different cluster.

### Configuration Arguments

The Nomad plugin allows you set static credentials with the `secret_id` argument.

```hcl
connection "nomad_cluster_prod" {
  plugin    = "nomad"
  address   = "https://some-nomad-cluster"
  region    = "global"
  secret_id = "<ACL token / secret ID>"
}
```
