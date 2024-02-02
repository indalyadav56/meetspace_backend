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
        "/v1/auth/forgot-password": {
            "post": {
                "description": "Forgot password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "forgot-password",
                "parameters": [
                    {
                        "description": "forgot password request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Login user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login-user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login user successfully"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User logout User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "user-logout",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/refresh-token": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "refresh jwt token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refresh token",
                "parameters": [
                    {
                        "description": "refresh token request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "description": "Register User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register-user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Register user successfully"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/send-email": {
            "post": {
                "description": "Send email to user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "send-email",
                "parameters": [
                    {
                        "description": "send email request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SendEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/verify-email": {
            "post": {
                "description": "Verify email otp.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "verify-email",
                "parameters": [
                    {
                        "description": "verify email request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.VerifyEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/chat/messages": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Register User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Message"
                ],
                "summary": "Register User account",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.GetChatMessageRequestBody"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/messages/{room_id}": {
            "get": {
                "description": "GetChatMessageByRoomId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Message"
                ],
                "summary": "GetChatMessageByRoomId",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.GetChatMessageRequestBody"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/room/contact": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "UserLogin User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "UserLogin User account",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/room/groups": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "UserLogin User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Group"
                ],
                "summary": "UserLogin User account",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/rooms": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetChatRooms",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "GetChatRooms",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateChatRoom",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "CreateChatRoom",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "DeleteChatRoom",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "DeleteChatRoom",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/client/users": {
            "get": {
                "description": "GetClientUsers account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client-User"
                ],
                "summary": "GetClientUsers account",
                "responses": {}
            },
            "post": {
                "description": "ClientAddUser account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client-User"
                ],
                "summary": "ClientAddUser account",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ClientAddUser"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/clients": {
            "get": {
                "description": "GetAllClients User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "GetAllClients User account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client's company name",
                        "name": "company_name",
                        "in": "query"
                    },
                    {
                        "description": "GetAllClients login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ClientCreateData"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "UserLogin User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "UserLogin User account",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ClientCreateData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/clients/{id}": {
            "get": {
                "description": "GetClientById User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "GetClientById User account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/check-email": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User create",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get all users",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user-update",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User create",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by ID",
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Client": {
            "type": "object",
            "properties": {
                "client_user_id": {
                    "type": "string"
                },
                "company_domain": {
                    "type": "string"
                },
                "company_name": {
                    "type": "string"
                },
                "company_size": {
                    "type": "integer"
                },
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "types.ClientAddUser": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "created_by": {
                    "$ref": "#/definitions/models.Client"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updated_by": {
                    "$ref": "#/definitions/models.Client"
                }
            }
        },
        "types.ClientCreateData": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.CreateUserData": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "created_by": {
                    "$ref": "#/definitions/models.Client"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_by": {
                    "$ref": "#/definitions/models.Client"
                }
            }
        },
        "types.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email",
                "new_password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "types.GetChatMessageRequestBody": {
            "type": "object",
            "properties": {
                "chat_room_id": {
                    "type": "string"
                },
                "current_user_id": {
                    "type": "string"
                }
            }
        },
        "types.LoginRequest": {
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
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "types.LogoutRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "types.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "types.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "types.SendEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "types.VerifyEmailRequest": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                }
            }
        },
        "utils.Metadata": {
            "type": "object",
            "properties": {
                "api_version": {
                    "type": "string",
                    "default": "v1"
                },
                "server_info": {
                    "$ref": "#/definitions/utils.ServerInfo"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/utils.Metadata"
                },
                "request_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "time_stamp": {
                    "type": "string"
                }
            }
        },
        "utils.ServerInfo": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
