// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Antony Thomas"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shows": {
            "get": {
                "description": "Get all shows currently stored in list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shows"
                ],
                "summary": "List all shows",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "list"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new show to the list",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "shows"
                ],
                "summary": "Add new show",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "The List API",
	Description:      "API for curating a movie and TV show watch list.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
