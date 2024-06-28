// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while Deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getall": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "GetAll page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetAll User",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetAll Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllUsers"
                        }
                    },
                    "401": {
                        "description": "Error while GetAll",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getbyid/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "GetById page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetById User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetById Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.User"
                        }
                    },
                    "401": {
                        "description": "Error while GetByIdd",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "LoginUser page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "/LoginUser",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "LoginUser Successful",
                        "schema": {
                            "$ref": "#/definitions/genprotos.User"
                        }
                    },
                    "401": {
                        "description": "Error while LoginUserd",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/registr": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "Create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update",
                        "name": "Update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Error while created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.GetAllUsers": {
            "type": "object",
            "properties": {
                "Users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.User"
                    }
                }
            }
        },
        "genprotos.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authourization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "Voting service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
