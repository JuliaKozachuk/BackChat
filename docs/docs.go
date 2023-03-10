// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/example/SignUp": {
            "post": {
                "description": "register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Сreate a new account"
                ],
                "summary": "writes the user to the database",
                "parameters": [
                    {
                        "description": "user",
                        "name": "get",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpInput"
                        }
                    },
                    "400": {
                        "description": "email is not unique"
                    },
                    "500": {
                        "description": "user registration failed"
                    }
                }
            }
        },
        "/example/SingIn": {
            "post": {
                "description": "User login to the system by mail, password, verification code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": " login user",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "jwt",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/example/userID": {
            "get": {
                "description": "getting all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/migrations.Users"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "verification_code": {
                    "type": "string"
                }
            }
        },
        "controllers.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Username          string ` + "`" + `swaggerignore:\"true\" json:\"username\" ` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "migrations.Users": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "id_user": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9888",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BackChat Api",
	Description:      "This is a  server BackChat.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
