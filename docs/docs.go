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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Get all events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get all events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetEventsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Create an event",
                "parameters": [
                    {
                        "description": "Event object",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.CreateResponse"
                        }
                    }
                }
            }
        },
        "/events/popular": {
            "get": {
                "description": "Get popular events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get popular events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Top K",
                        "name": "topK",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetEventsResponse"
                        }
                    }
                }
            }
        },
        "/events/{id}": {
            "get": {
                "description": "Get an event by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get an event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Update an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Event",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Delete an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/events/{id}/join": {
            "post": {
                "description": "Join an event",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Join an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/locations": {
            "get": {
                "description": "Get locations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Get locations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Location level",
                        "name": "level",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Location"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Create a new location",
                "parameters": [
                    {
                        "description": "Create Location",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateLocationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payload.CreateResponse"
                        }
                    }
                }
            }
        },
        "/locations/{id}": {
            "put": {
                "description": "Update a location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Update a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Location",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateLocationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Location"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a location",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "locations"
                ],
                "summary": "Delete a location",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Location ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/reviews/{event_id}": {
            "get": {
                "description": "Get reviews of an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Get reviews of an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GetReviewsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a review",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Create a review",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Review",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateReviewRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payload.CreateResponse"
                        }
                    }
                }
            }
        },
        "/reviews/{id}": {
            "delete": {
                "description": "Delete a review",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Delete a review",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Review ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
            },
            "put": {
                "description": "Update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by id",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Event": {
            "type": "object",
            "properties": {
                "age_limit": {
                    "type": "integer"
                },
                "capacity": {
                    "type": "integer"
                },
                "categories": {
                    "description": "relational fields",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dress_code": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organizer": {
                    "description": "non-mandatory fields",
                    "type": "string"
                },
                "participants": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "price": {
                    "type": "number"
                },
                "start_time": {
                    "type": "string"
                },
                "total_likes": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Level": {
            "type": "string",
            "enum": [
                "country",
                "state",
                "city"
            ],
            "x-enum-varnames": [
                "Country",
                "State",
                "City"
            ]
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "level": {
                    "$ref": "#/definitions/models.Level"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Review": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "event_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "joined_event": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "payload.CreateEventRequest": {
            "type": "object",
            "properties": {
                "age_limit": {
                    "type": "integer"
                },
                "capacity": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "dress_code": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organizer": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "payload.CreateLocationRequest": {
            "type": "object",
            "required": [
                "level",
                "name"
            ],
            "properties": {
                "level": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "payload.CreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "payload.CreateReviewRequest": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "payload.GeneralResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "payload.GetEventsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Event"
                    }
                }
            }
        },
        "payload.GetReviewsResponse": {
            "type": "object",
            "properties": {
                "average_rating": {
                    "type": "number"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Review"
                    }
                }
            }
        },
        "payload.LoginRequest": {
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
        "payload.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "payload.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
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
        "payload.UpdateEventRequest": {
            "type": "object",
            "properties": {
                "age_limit": {
                    "type": "integer"
                },
                "capacity": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "dress_code": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organizer": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateLocationRequest": {
            "type": "object",
            "properties": {
                "level": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "payload.UpdateUserRequest": {
            "type": "object",
            "properties": {
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
	Host:             "localhost:9999",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Match-Event Backend API",
	Description:      "A matchmaking service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
