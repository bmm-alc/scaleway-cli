<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw account`
Account API
  
- [SSH keys management commands](#ssh-keys-management-commands)
  - [Add an SSH key to your project](#add-an-ssh-key-to-your-project)
  - [Get an SSH key from your project](#get-an-ssh-key-from-your-project)
  - [Initialize SSH key](#initialize-ssh-key)
  - [List all SSH keys of your project](#list-all-ssh-keys-of-your-project)
  - [Remove an SSH key from your project](#remove-an-ssh-key-from-your-project)
  - [Update an SSH key on your project](#update-an-ssh-key-on-your-project)

  
## SSH keys management commands

SSH keys management commands.


### Add an SSH key to your project

Add an SSH key to your project.

**Usage:**

```
scw account ssh-key add [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name |  | The name of the SSH key |
| public-key | Required | SSH public key. Currently ssh-rsa, ssh-dss (DSA), ssh-ed25519 and ecdsa keys with NIST curves are supported |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| organization-id |  | Organization ID to use. If none is passed the default organization ID will be used |


**Examples:**


Add a given ssh key
```
scw account ssh-key add name=foobar public-key="$(cat <path/to/your/public/key>)"
```




### Get an SSH key from your project

Get an SSH key from your project.

**Usage:**

```
scw account ssh-key get <ssh-key-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| ssh-key-id | Required | The ID of the SSH key |



### Initialize SSH key

Initialize SSH key.

**Usage:**

```
scw account ssh-key init
```



### List all SSH keys of your project

List all SSH keys of your project.

**Usage:**

```
scw account ssh-key list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `name_asc`, `name_desc` |  |
| name |  |  |
| project-id |  |  |
| organization-id |  |  |



### Remove an SSH key from your project

Remove an SSH key from your project.

**Usage:**

```
scw account ssh-key remove <ssh-key-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| ssh-key-id | Required |  |


**Examples:**


Remove a given SSH key
```
scw account ssh-key remove 11111111-1111-1111-1111-111111111111
```




### Update an SSH key on your project

Update an SSH key on your project.

**Usage:**

```
scw account ssh-key update <ssh-key-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| ssh-key-id | Required |  |
| name |  | Name of the SSH key |


