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
        "/api/accounts": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create a user account with given balance",
                "operationId": "account-create",
                "parameters": [
                    {
                        "description": "amount of money on the new account",
                        "name": "amount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateAccount"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/accounts/deposit/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Deposit funds to an account and add en entry to the transfers table",
                "operationId": "account-deposit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "amount of money",
                        "name": "amount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DepositFunds"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/accounts/withdraw/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Withdraw funds from an account and add en entry to the transfers table",
                "operationId": "account-withdraw",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "amount of money",
                        "name": "amount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WithdrawFunds"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/accounts/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get a user account information by its ID",
                "operationId": "account-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account info",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.CreateAccount": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "models.DepositFunds": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "models.WithdrawFunds": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "REST API Service",
	Description:      "Swagger API for REST API Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
