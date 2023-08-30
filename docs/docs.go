// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "jobor",
            "url": "https://github.com/v-mars/jobor"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/jobor/audit-log": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/env": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/jobor/env/:id": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/jobor/envs": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/gen-token": {
            "post": {
                "responses": {}
            }
        },
        "/api/v1/jobor/log": {
            "get": {
                "description": "jobor log get",
                "tags": [
                    "jobor log"
                ],
                "summary": "jobor log get summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/jobor/migrate": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/state-code": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/task": {
            "get": {
                "description": "jobor task get",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task get summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "jobor task post",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task post summary",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.PostTaskReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/jobor/task/{id}": {
            "get": {
                "description": "jobor task get by id",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task get by id summary",
                "responses": {}
            },
            "put": {
                "description": "jobor task put",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task put summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.PutTaskReq"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "jobor task run",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task run summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "jobor task delete",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task delete summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/jobor/tasks": {
            "get": {
                "description": "jobor task all get",
                "tags": [
                    "jobor task"
                ],
                "summary": "jobor task all get summary",
                "responses": {}
            }
        },
        "/api/v1/jobor/user-self": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/user-switch/:user_id": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/jobor/worker": {
            "get": {
                "description": "jobor worker get",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker get summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "jobor worker post",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker post summary",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/worker.PostWorkerReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/jobor/worker/{id}": {
            "get": {
                "description": "jobor worker get by id",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker get by id summary",
                "responses": {}
            },
            "put": {
                "description": "jobor worker put",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker put summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/worker.PutWorkerReq"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "jobor worker delete",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker delete summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/jobor/workers": {
            "get": {
                "description": "jobor worker all get",
                "tags": [
                    "jobor worker"
                ],
                "summary": "jobor worker all get summary",
                "responses": {}
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "user login summary",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mw.LoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/logout": {
            "post": {
                "description": "user logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "user logout summary",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mw.LogoutReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/oidc/callback": {
            "get": {
                "description": "oidc callback",
                "tags": [
                    "login"
                ],
                "summary": "oidc callback summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/oidc/redirect": {
            "get": {
                "description": "oidc redirect login",
                "tags": [
                    "login"
                ],
                "summary": "oidc redirect login summary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "next",
                        "name": "next",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/sys/api": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/sys/api-auto-update": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/sys/api/:id": {
            "put": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/sys/apis": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/sys/role": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/sys/role-tree": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/sys/role/:id": {
            "put": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/sys/roles": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/sys/user": {
            "get": {
                "responses": {}
            },
            "post": {
                "responses": {}
            }
        },
        "/api/v1/sys/user-sync": {
            "get": {
                "responses": {}
            }
        },
        "/api/v1/sys/user/:id": {
            "get": {
                "responses": {}
            },
            "put": {
                "responses": {}
            },
            "delete": {
                "responses": {}
            }
        },
        "/api/v1/sys/users": {
            "get": {
                "responses": {}
            }
        },
        "/routes": {
            "get": {
                "responses": {}
            }
        }
    },
    "definitions": {
        "mw.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "mw.LogoutReq": {
            "type": "object",
            "properties": {
                "id_token_hint": {
                    "type": "string"
                },
                "post_logout_redirect_uri": {
                    "type": "string"
                }
            }
        },
        "task.Api": {
            "type": "object",
            "properties": {
                "auth_data": {
                    "$ref": "#/definitions/task.AuthData"
                },
                "auth_method": {
                    "type": "string"
                },
                "body": {
                    "type": "string"
                },
                "content_type": {
                    "type": "string"
                },
                "form_data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.KvFiled"
                    }
                },
                "forms": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "header": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "header_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.KvFiled"
                    }
                },
                "method": {
                    "type": "string"
                },
                "payload": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "www_form_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.KvFiled"
                    }
                }
            }
        },
        "task.AuthData": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "task.Dingding": {
            "type": "object",
            "properties": {
                "enable": {
                    "type": "boolean"
                },
                "webhooks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.Webhooks"
                    }
                }
            }
        },
        "task.Email": {
            "type": "object",
            "properties": {
                "enable": {
                    "type": "boolean"
                },
                "receivers": {
                    "description": "多个邮箱地址以逗号分割",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "task.KvFiled": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "task.Lark": {
            "type": "object",
            "properties": {
                "enable": {
                    "type": "boolean"
                },
                "webhooks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.Webhooks"
                    }
                }
            }
        },
        "task.Notify": {
            "type": "object",
            "properties": {
                "dingding": {
                    "$ref": "#/definitions/task.Dingding"
                },
                "email": {
                    "$ref": "#/definitions/task.Email"
                },
                "lark": {
                    "$ref": "#/definitions/task.Lark"
                },
                "webhook": {
                    "$ref": "#/definitions/task.Webhook"
                }
            }
        },
        "task.PostTaskReq": {
            "type": "object",
            "properties": {
                "alarm_policy": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/task.TaskData"
                },
                "deleted": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "expect_code": {
                    "type": "integer"
                },
                "expect_content": {
                    "type": "string"
                },
                "expr": {
                    "type": "string"
                },
                "lang": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "next": {
                    "type": "string"
                },
                "notify": {
                    "$ref": "#/definitions/task.Notify"
                },
                "prev": {
                    "type": "string"
                },
                "retry": {
                    "type": "integer"
                },
                "route_policy": {
                    "type": "integer"
                },
                "routing_key": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "timeout": {
                    "type": "integer"
                },
                "updater": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "task.PutTaskReq": {
            "type": "object",
            "properties": {
                "alarm_policy": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/task.TaskData"
                },
                "deleted": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "expect_code": {
                    "type": "integer"
                },
                "expect_content": {
                    "type": "string"
                },
                "expr": {
                    "type": "string"
                },
                "lang": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "next": {
                    "type": "string"
                },
                "notify": {
                    "$ref": "#/definitions/task.Notify"
                },
                "prev": {
                    "type": "string"
                },
                "retry": {
                    "type": "integer"
                },
                "route_policy": {
                    "type": "integer"
                },
                "routing_key": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "timeout": {
                    "type": "integer"
                },
                "updater": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "task.TaskData": {
            "type": "object",
            "properties": {
                "api": {
                    "$ref": "#/definitions/task.Api"
                },
                "content": {
                    "type": "string"
                }
            }
        },
        "task.Webhook": {
            "type": "object",
            "properties": {
                "enable": {
                    "type": "boolean"
                },
                "urls": {
                    "description": "多个api url以逗号分割",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "task.Webhooks": {
            "type": "object",
            "properties": {
                "secret": {
                    "type": "string"
                },
                "webhook_url": {
                    "type": "string"
                }
            }
        },
        "worker.PostWorkerReq": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "lease_update": {
                    "type": "integer"
                },
                "routing_key": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "worker.PutWorkerReq": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "lease_update": {
                    "type": "integer"
                },
                "routing_key": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Jobor 定时任务 API",
	Description:      "支持host,server,network等\nAuthorization Bearer token\nheader:  x-enc-data = yes",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
