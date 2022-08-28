
# bash 1.0

Ubuntu bash environment

---
- #### Categories: build, misc
- #### Image: direktiv.azurecr.io/functions/bash 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/bash/issues
- #### URL: https://github.com/direktiv-apps/bash
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About bash

This function provides a Ubuntu 22.04 environment. The following packages are installed:

- git 
- sed 
- wget 
- grep 
- curl 
- make 
- jq 
- openssh-client

The function can run multiple commands but if multiple commands have to be chained together or use pipes it requires a 'bash -c' command, e.g.  `bash -c 'cd dir && cat file.txt > newfile'`

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: bash
  image: direktiv.azurecr.io/functions/bash:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: bash
  type: action
  action:
    function: bash
    input: 
      commands:
      - command: ls -la
```
   #### Using ad-hoc files
```yaml
- id: bash
  type: action
  action:
    function: bash
    input: 
      files:
      - name: hello.sh
        data: |-
          #!/bin/bash
          echo "Hello World"
        mode: '0755'
      commands:
      - command: ./hello.sh
```
   #### Return JSON file
```yaml
- id: bash 
  type: action
  action:
    function: bash
    input: 
      commands:
      - command: 'bash -c "echo { \\\"hello\\\": \\\"world\\\" } > data.json"'
      - command: cat data.json
```
   #### Pipe
```yaml
- id: bash  
  type: action
  action:
    function: bash
    input: 
      commands:
      - command: bash -c 'echo "File Data" > myfile.txt'
      - command: cat myfile.txt
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "",
    "success": true
  },
  {
    "result": {
      "hello": "world"
    },
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bash | [][PostOKBodyBashItems](#post-o-k-body-bash-items)| `[]*PostOKBodyBashItems` |  | |  |  |


#### <span id="post-o-k-body-bash-items"></span> postOKBodyBashItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `ls -la` |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| envs | [][PostParamsBodyCommandsItemsEnvsItems](#post-params-body-commands-items-envs-items)| `[]*PostParamsBodyCommandsItemsEnvsItems` |  | | Environment variables set for each command. | `[{"name":"MYVALUE","value":"hello"}]` |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |


#### <span id="post-params-body-commands-items-envs-items"></span> postParamsBodyCommandsItemsEnvsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | Name of the variable. |  |
| value | string| `string` |  | | Value of the variable. |  |

 
