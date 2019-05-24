// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-17 16:24:42.2766447 +0700 +07 m=+0.069960801

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is rest api demo",
        "title": "REST API TO DO LIST",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "quanpham",
            "url": "https://www.facebook.com/quanphamptit",
            "email": "phamducquanptithcm@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:3001",
    "basePath": "/api/v1",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "body: username, password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.SwaggerLoginBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"user\": {\"id\": 1,\"created_at\": \"2019-05-15T16:44:54+07:00\",\"username\": \"ebisu\",\"password\": \"$2a$10$UEC866tSde0fZTyRQBTOvOHv9T4qUMgUBhREXkeGWG5s2gHXXWCzm\"}, \"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ\"}}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list": {
            "post": {
                "security": [
                    {
                        "OAuth2Application": []
                    }
                ],
                "description": "Get list todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get list todo",
                "parameters": [
                    {
                        "description": "List todos",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.ListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"OK\", \"data\": { \"total\": 1,\"todos\": [{ \"id\": 1,\"title\": \"learning golang\",\"status\": 0,\"created_at\": \"2019-05-15 16:57:29\",\"updated_at\": \"2019-05-15 16:57:29\"}] }}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.SwaggerListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/todo": {
            "post": {
                "security": [
                    {
                        "OAuth2Application": []
                    }
                ],
                "description": "Create a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "description": "Create a new todo",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\": {\"todo\": {\"id\": 1,\"created_at\": \"2019-05-15T16:57:29.4549295+07:00\",\"status\": 0,\"user_id\": 1,\"title\": \"learning golang\"} } }",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/todo/{id}": {
            "put": {
                "security": [
                    {
                        "OAuth2Application": []
                    }
                ],
                "description": "Update todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Update todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id todo update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data update",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"OK\", \"data\": \"{\"todo\": {\"title\": \"build rest api todo list demo\", \"status\": 0} }\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2Application": []
                    }
                ],
                "description": "Delete a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id todo delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 0, \"message\": \"OK\", \"data\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"user\": { \"id\": 2,\"created_at\": \"2019-05-15T16:54:05.8135074+07:00\",\"username\": \"quanpham\",\"password\": \"$2a$10$Kb4PV4dc.jaUnnjTviWMB.DY0JaHdKBRSEegYrWUZhgXHJQ.tkuu6\"}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.TodoInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.TodoModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "todo.CreateRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "todo.DataResponse": {
            "type": "object",
            "properties": {
                "todos": {
                    "type": "object",
                    "$ref": "#/definitions/models.TodoModel"
                }
            }
        },
        "todo.ListRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "todo.SwaggerListResponse": {
            "type": "object",
            "properties": {
                "todos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TodoInfo"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "todo.UpdateRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
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
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerLoginBody": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}