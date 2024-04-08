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
            "url": "https://t.me/Zatrasz",
            "email": "zatrasz@ya.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars": {
            "get": {
                "description": "Получить список автомобилей с возможностью фильтрации и пагинации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получение данных с фильтрацией по всем полям и пагинацией",
                "operationId": "get-cars-and-pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Фильтр по названию автомобиля",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы для пагинации",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество элементов на странице для пагинации",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список автомобилей",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Car"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Добавить в систему несколько автомобилей, используя их номера регистрации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Добавление нескольких автомобилей",
                "operationId": "add-cars",
                "parameters": [
                    {
                        "description": "Массив номеров регистрации автомобилей",
                        "name": "request.regNums",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Автомобили успешно добавлены",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении автомобилей",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Изменить данные автомобиля по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Изменение одного или нескольких полей по идентификатору",
                "operationId": "update-cars-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Данные автомобиля для обновления",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Данные автомобиля успешно обновлены",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Не удалось проанализировать запрос JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при обновлении данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "delete": {
                "description": "Удалить автомобиль по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Удаление автомобиля по идентификатору",
                "operationId": "delete-cars-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Автомобиль успешно удален",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный идентификатор автомобиля",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при удалении автомобиля",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Car": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.People"
                },
                "ownerId": {
                    "type": "integer"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.People": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4141",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "ТЗ Go - апрель.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}