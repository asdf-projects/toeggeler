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
        "/authenticate": {
            "post": {
                "description": "Authenticate (Passwords for all available users is \"1234\")",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authenticate"
                ],
                "summary": "Authenticate by username and password",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthenResponse"
                        }
                    }
                }
            }
        },
        "/games": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Submit a game by listing every event.\nAvailable events: \u003cul\u003e\u003cli\u003eGAME_START\u003c/li\u003e\u003cli\u003eGOAL\u003c/li\u003e\u003cli\u003eOWN_GOAL\u003c/li\u003e\u003cli\u003eFOETELI\u003c/li\u003e\u003cli\u003eGAME_END\u003c/li\u003e\u003c/ul\u003e",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Games"
                ],
                "summary": "Submit a game",
                "parameters": [
                    {
                        "description": "Game events",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.GameEvent"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Game"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/stats": {
            "get": {
                "description": "Get statistics for all available users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Get statistics for all available users (FAKE DATA)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.Stats"
                            }
                        }
                    }
                }
            }
        },
        "/stats/{id}": {
            "get": {
                "description": "Get statistics for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stats"
                ],
                "summary": "Get statistics for a user (FAKE DATA, available IDs: 1, 2, 3, 4)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Stats"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get a list of all available users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a list of all available users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a user by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete an existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AuthenRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.AuthenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateUserRequest": {
            "type": "object",
            "required": [
                "mail",
                "password",
                "username"
            ],
            "properties": {
                "mail": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.Game": {
            "type": "object",
            "properties": {
                "gameEnd": {
                    "type": "integer"
                },
                "gameId": {
                    "type": "string"
                },
                "gameStart": {
                    "type": "integer"
                },
                "score": {
                    "$ref": "#/definitions/controllers.Score"
                },
                "team1": {
                    "$ref": "#/definitions/controllers.Team"
                },
                "team2": {
                    "$ref": "#/definitions/controllers.Team"
                }
            }
        },
        "controllers.Score": {
            "type": "object",
            "properties": {
                "team1": {
                    "type": "integer"
                },
                "team2": {
                    "type": "integer"
                }
            }
        },
        "controllers.Stats": {
            "type": "object",
            "properties": {
                "foetelis": {
                    "type": "integer"
                },
                "goals": {
                    "type": "integer"
                },
                "losses": {
                    "type": "integer"
                },
                "ownGoals": {
                    "type": "integer"
                },
                "playerId": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "wins": {
                    "type": "integer"
                }
            }
        },
        "controllers.Team": {
            "type": "object",
            "properties": {
                "defense": {
                    "type": "integer"
                },
                "offense": {
                    "type": "integer"
                }
            }
        },
        "controllers.UpdateUserRequest": {
            "type": "object",
            "required": [
                "mail"
            ],
            "properties": {
                "mail": {
                    "type": "string"
                }
            }
        },
        "models.GameEvent": {
            "description": "Game event description",
            "type": "object",
            "properties": {
                "event": {
                    "description": "Event type",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.GameEventType"
                        }
                    ]
                },
                "gameId": {
                    "description": "ID of the game",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the event",
                    "type": "integer"
                },
                "player": {
                    "description": "Required when is is \"GOAL\", \"OWN_GOAL\" or \"FOETELI\"",
                    "type": "integer"
                },
                "team1": {
                    "description": "Required when event is \"GAME_START\"",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Team"
                        }
                    ]
                },
                "team2": {
                    "description": "Required when event is \"GAME_START\"",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Team"
                        }
                    ]
                },
                "timestamp": {
                    "description": "Unix timestamp",
                    "type": "integer"
                }
            }
        },
        "models.GameEventType": {
            "type": "string",
            "enum": [
                "GAME_START",
                "GAME_END",
                "GOAL",
                "OWN_GOAL",
                "FOETELI"
            ],
            "x-enum-varnames": [
                "GAME_START",
                "GAME_END",
                "GOAL",
                "OWN_GOAL",
                "FOETELI"
            ]
        },
        "models.Team": {
            "description": "A team consists of an offensive and defensive player",
            "type": "object",
            "properties": {
                "defense": {
                    "description": "ID of the player on defense",
                    "type": "integer"
                },
                "offense": {
                    "description": "ID of the player on offense",
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "description": "User information",
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the user",
                    "type": "integer"
                },
                "mail": {
                    "description": "Email address",
                    "type": "string"
                },
                "username": {
                    "description": "Username",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Enter the token with the ` + "`" + `Bearer: ` + "`" + ` prefix, e.g. \"Bearer abcde12345\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Toeggeler Server API",
	Description:      "Api specification",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}