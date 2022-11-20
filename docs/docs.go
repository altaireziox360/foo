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
        "/api/v1/health": {
            "get": {
                "description": "Responds with the server status as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Checks the status of the server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utility.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Send a dummy post request to test the status of the server",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Checks the status of the server",
                "parameters": [
                    {
                        "description": "Ping JSON",
                        "name": "ping",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Ping"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utility.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Ping": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "utility.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "description": "for errors that occur even if request is successful"
                },
                "extra": {},
                "message": {
                    "type": "string"
                },
                "name": {
                    "description": "name of the error",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/api/v1/",
	Schemes:          []string{"http"},
	Title:            "Minergram",
	Description:      "A picture mining service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
