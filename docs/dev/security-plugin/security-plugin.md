# Security plugin support for ODFE-cli

This document describes the command embedded into `odfe-cli` that supports `odfe-security-plugin`.

Following are some examples of using such command:

```sh
# shows help message for security plugin
odfe-cli security --help

# retrieve all internal users hosted in security plugin
# print to standard output in default format
odfe-cli security get user

# retrieve all internal users hosted in security plugin,
# print to standard out in json format
odfe-cli security get user --json

# retrieve admin users details and print to standard out in yaml format
odfe-cli security get user admin --yaml

# create a role in security plugin from file my-role.json
# print query result to standard output in json format
odfe-cli security create role my-role -f my-role.json --json

# create a role in security plugin from file my-role.json
odfe-cli security patch role-mapping my-role -f my-role-patch.json

# delete a tenant in security plugin
odfe-cli security delete tenant my-tenant

# delete multiple tenant in security plugin
odfe-cli security delete tenant bad-tenant-0 bad-tenant-1 bad-tenant-2

# change password
odfe-cli account change-password -f pwd.json

# print configuration in security plugin to standard out in json format
odfe-cli config --json

# update configuration in security plugin
odfe-cli config update -f config.json

# update configuration in security plugin
odfe-cli config patch -f config-patch.json

# flush cache
odfe-cli security flush-cache
```

---

## Commands

### CRUD commands

This table lists all `crud` operations to resources hosted in `odfe-security-plugin`

Available resources are:
  - `action-group` (alias: `ac`)
  - `user` (alias: `u`)
  - `role` (alias: `r`)
  - `role-mapping` (alias: `rmp`)
  - `tenant` (alias: `tnt`) 

Command | Arguments | Description
--------|-----------|-----------
`get` | `<resource> [name]` | Get resources in security plugin. If resource name is not specified, get all available resources.
`create` | `<resource> <name>` | Create resource in security plugin.
`patch` | `<resource> [name]` | Patch resources in security plugin. If resource name is not specified, patch all available resources.
`delete` | `<resource> <names>` | Delete resource in security plugin.

#### Flags

CRUD commands support following flags

Flag | Alias | Description
-----|-------|-----------
`--file` | `-f` | Input file that is used as query body, used by `create` and `patch`
`--json` | \- | Output query response to `json` format
`--yaml` | \- | Output query response to `yaml` format

---

### Other commands

Command | Arguments | Description
--------|-----------|-----------
`account` | `[change-password]` | Show current account information or change password. If no file input provided, this command requires user to manually input current and new password
`config` | `[update]` or `[patch]` | Update configuration file in `odfe-security-plugin`
`flush-cache` | \- | Execute query on `_opendistro/_security/api/cache` with `DELETE` method
`health-check` | \- | Execute query on `_opendistro/_security/health` with `GET` method

#### Flags

CRUD commands support following flags

Flag | Alias | Description
-----|-------|-----------
`--file` | `-f` | Input file that is used as query body, used by `config` and `account change-password`
`--json` | \- | Output query response to `json` format
`--yaml` | \- | Output query response to `yaml` format
