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
        "/api/v1/": {
            "get": {
                "tags": [
                    "/"
                ],
                "summary": "Root URL - for health check",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/auth/login/": {
            "post": {
                "description": "비밀번호는 영문 + 숫자 8-12자리",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "로그인 API",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LogInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LogInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.LogInResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.LogInResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.LogInResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/bookmark/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookmark"
                ],
                "summary": "특정 유저가 저장한 북마크 중 offset부터 limit까지 목록을 반환",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit과 offset은 같이 입력해야 합니다",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit과 offset은 같이 입력해야 합니다",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.GetAllBookmarksResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Bookmark"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllBookmarksResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "새 북마크를 DB에 저장. 북마크 링크는 필수 데이터이고, 그 외는 옵셔널.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookmark"
                ],
                "summary": "특정 유저가 북마크를 DB에 추가하는 API",
                "parameters": [
                    {
                        "description": "body params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddBookmarkRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AddBookmarkResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Bookmark"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.AddBookmarkResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/bookmark/{id}/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookmark"
                ],
                "summary": "특정 유저가 저장한 북마크 중 id가 일치하는 북마크 상세 정보 1개를 반환",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bookmark ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.GetBookmarkByIdResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Bookmark"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GetBookmarkByIdResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "모든 유저 목록 반환",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "특정 id부터 조회할 때 사용",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit 개수만큼 조회할 때 사용",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllUsersWithErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllUsersWithErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "새 유저 추가",
                "parameters": [
                    {
                        "description": "비밀번호는 영문 + 숫자 8-12자리",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.AddUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.AddUserWithErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.AddUserWithErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/check-email": {
            "get": {
                "description": "입력한 이메일을 사용하는 유저가 있다면 Body의 Message로 True 반환, 없다면 False 반환.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "이메일 중복 체크.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "중복 체크 할 이메일 주소 입력. 최대 길이 50자 제한.",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CheckEmailDuplicateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CheckEmailDuplicateWithErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.CheckEmailDuplicateWithErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "UserID를 사용해 유저 1명 정보 읽기",
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
                            "$ref": "#/definitions/models.GetUserByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GetUserByIdWithErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GetUserByIdWithErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.AddBookmarkRequest": {
            "type": "object",
            "required": [
                "link",
                "user_id"
            ],
            "properties": {
                "link": {
                    "type": "string",
                    "example": "https://cheesecat47.github.io"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "models.AddBookmarkResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.AddUserRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "5~50자 길이. 자세한 형식은 [go-playground/validator] 참고\n\n[go-playground/validator]: https://github.com/go-playground/validator",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 5,
                    "example": "cheesecat47@gmail.com"
                },
                "name": {
                    "description": "알파벳, 숫자, 유니코드 문자 사용 가능",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                },
                "password": {
                    "description": "영문 + 숫자 8-12자리",
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 8,
                    "example": "qwerty123"
                }
            }
        },
        "models.AddUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "생성된 유저 정보",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.User"
                        }
                    ]
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.AddUserWithErrorResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Authentication": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "description": "5~50자 길이. 자세한 형식은 [go-playground/validator] 참고\n\n[go-playground/validator]: https://github.com/go-playground/validator",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Bookmark": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.CheckEmailDuplicateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "생성된 인증 정보",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Authentication"
                        }
                    ]
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.CheckEmailDuplicateWithErrorResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetAllBookmarksResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetAllUsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "description": "유저 정보 배열",
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        },
                        "size": {
                            "description": "유저 배열 길이",
                            "type": "integer"
                        }
                    }
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetAllUsersWithErrorResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetBookmarkByIdResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetUserByIdResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "ID가 일치하는 유저 정보",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.User"
                        }
                    ]
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.GetUserByIdWithErrorResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.LogInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 5,
                    "example": ""
                },
                "password": {
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 8,
                    "example": "비밀번호는 영문 + 숫자 8-12자리"
                }
            }
        },
        "models.LogInResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "Authority": {
                    "description": "사용자 권한\n - [AuthorityNone]\n - [AuthorityAdmin]",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "알파벳, 숫자, 유니코드 문자 사용 가능",
                    "type": "string"
                },
                "updatedAt": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
