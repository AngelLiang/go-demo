// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/loginWithJSON": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录认证"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "string",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/main.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.LoginUser": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "Password字段表示登录密码\nbind:\"required\"表示这个字段是必需的，需要在绑定时提供值\njson:\"password\"表示在JSON序列化/反序列化时使用的键名",
                    "type": "string"
                },
                "username": {
                    "description": "Username字段表示登录用户名\nbind:\"required\"表示这个字段是必需的，需要在绑定时提供值\njson:\"username\"表示在JSON序列化/反序列化时使用的键名",
                    "type": "string"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
