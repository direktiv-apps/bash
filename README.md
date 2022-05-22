
# bash 1.0

Simple bash function.

---
- #### Image: direktiv/bash 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/bash/issues
- #### URL: https://github.com/direktiv-apps/bash
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About bash

The bash function is providing bash with pre-installed packages like jq, yq, curl, git and more.

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: bash
    image: direktiv/bash:1.0
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
          - command: env
            envs:
            - name: HELLO
   ```
   #### Return JSON file
   ```yaml
   - id: getter 
      type: action
      action:
        function: bash
        input: 
          commands:
          - command: 'bash -c "echo { \\\"hello\\\": \\\"world\\\" } > data.json"'
          - command: cat data.json
   ```
   #### Run Script
   ```yaml
   - id: getter 
      type: action
      action:
        function: bash
        files:
        - key: run.sh
          value: |
            #!/bin/bash
            echo Hello World
          scope: inline
        input: 
          commands:
          - command: ./run.sh
   ```
   #### Pipe
   ```yaml
   - id: req
    type: action
    action:
      function: bash
   ```

### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Returns array of command results. If a command returned JSON it will be included as JSON object.

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of bash commands. |  |
| continue | boolean| `bool` |  | | If set to true all commands are getting executed and errors ignored. | `true` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `ls -la` |
| continue | boolean| `bool` |  | |  |  |
| envs | [][PostParamsBodyCommandsItemsEnvsItems](#post-params-body-commands-items-envs-items)| `[]*PostParamsBodyCommandsItemsEnvsItems` |  | | Environment variables set for each command. | `[{"name":"MYVALUE","value":"hello"}]` |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |


#### <span id="post-params-body-commands-items-envs-items"></span> postParamsBodyCommandsItemsEnvsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | Name of the variable. |  |
| value | string| `string` |  | | Value of the variable. |  |

 
