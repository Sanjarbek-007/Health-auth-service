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
        "/change-password": {
            "post": {
                "description": "Changes user password",
                "tags": [
                    "auth"
                ],
                "summary": "Changes password",
                "parameters": [
                    {
                        "description": "Change password data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password changed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/forgot-password": {
            "post": {
                "description": "Sends password reset code to user's email",
                "tags": [
                    "auth"
                ],
                "summary": "Forgot password",
                "parameters": [
                    {
                        "description": "User email",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ForgotPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset code sent successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs user in",
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tokens"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Logouts user",
                "tags": [
                    "auth"
                ],
                "summary": "Logouts user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged out successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid user id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Refreshes refresh token",
                "tags": [
                    "auth"
                ],
                "summary": "Refreshes token",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tokens"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user",
                "tags": [
                    "auth"
                ],
                "summary": "Registers user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResp"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reset-password": {
            "post": {
                "description": "Resets user password",
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "Reset password data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ResetPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error while processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ChangePasswordReq": {
            "type": "object",
            "properties": {
                "current_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "models.ForgotPasswordReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.LoginReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.RefreshToken": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.RegisterReq": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
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
                }
            }
        },
        "models.RegisterResp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "models.ResetPasswordReq": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "models.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Authorazation",
	Description:      "Authorazation API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
