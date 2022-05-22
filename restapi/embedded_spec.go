// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Simple bash function.",
    "title": "bash",
    "version": "1.0",
    "x-direktiv-meta": {
      "category": "Unknown",
      "container": "direktiv/bash",
      "issues": "https://github.com/direktiv-apps/bash/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "The bash function is providing bash with pre-installed packages like jq, yq, curl, git and more.",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/bash"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of bash commands.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string",
                        "example": "ls -la"
                      },
                      "continue": {
                        "type": "boolean"
                      },
                      "envs": {
                        "description": "Environment variables set for each command.",
                        "type": "array",
                        "items": {
                          "type": "object",
                          "properties": {
                            "name": {
                              "description": "Name of the variable.",
                              "type": "string"
                            },
                            "value": {
                              "description": "Value of the variable.",
                              "type": "string"
                            }
                          }
                        },
                        "example": [
                          {
                            "name": "MYVALUE",
                            "value": "hello"
                          }
                        ]
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "continue": {
                  "description": "If set to true all commands are getting executed and errors ignored.",
                  "type": "boolean",
                  "default": false,
                  "example": true
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns array of command results. If a command returned JSON it will be included as JSON object.",
            "schema": {
              "type": "object",
              "additionalProperties": false
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Body.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "runtime-envs": "[\n{{- range $index, $element := .Item.Envs }}\n{{- if $index}},{{- end}}\n\"{{ $element.Name }}={{ $element.Value }}\"\n{{- end }}\n]\n",
              "silent": "{{ .Item.Silent }}"
            }
          ]
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Basic"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Return JSON file"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Run Script"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Pipe"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: bash\n    image: direktiv/bash:1.0\n    type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Simple bash function.",
    "title": "bash",
    "version": "1.0",
    "x-direktiv-meta": {
      "category": "Unknown",
      "container": "direktiv/bash",
      "issues": "https://github.com/direktiv-apps/bash/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "The bash function is providing bash with pre-installed packages like jq, yq, curl, git and more.",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/bash"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of bash commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/CommandsItems0"
                  }
                },
                "continue": {
                  "description": "If set to true all commands are getting executed and errors ignored.",
                  "type": "boolean",
                  "default": false,
                  "example": true
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Returns array of command results. If a command returned JSON it will be included as JSON object.",
            "schema": {
              "type": "object",
              "additionalProperties": false
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Body.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "runtime-envs": "[\n{{- range $index, $element := .Item.Envs }}\n{{- if $index}},{{- end}}\n\"{{ $element.Name }}={{ $element.Value }}\"\n{{- end }}\n]\n",
              "silent": "{{ .Item.Silent }}"
            }
          ]
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Basic"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Return JSON file"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Run Script"
          },
          {
            "content": "- id: req\n     type: action\n     action:\n       function: bash",
            "title": "Pipe"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: bash\n    image: direktiv/bash:1.0\n    type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "CommandsItems0": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string",
          "example": "ls -la"
        },
        "continue": {
          "type": "boolean"
        },
        "envs": {
          "description": "Environment variables set for each command.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/CommandsItems0EnvsItems0"
          },
          "example": [
            {
              "name": "MYVALUE",
              "value": "hello"
            }
          ]
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      }
    },
    "CommandsItems0EnvsItems0": {
      "type": "object",
      "properties": {
        "name": {
          "description": "Name of the variable.",
          "type": "string"
        },
        "value": {
          "description": "Value of the variable.",
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
}
