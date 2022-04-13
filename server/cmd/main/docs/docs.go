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
        "/apartments": {
            "get": {
                "description": "Get details of all apartments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apartments"
                ],
                "summary": "Get details of all apartments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Apartment"
                            }
                        }
                    }
                }
            }
        },
        "/complaints": {
            "get": {
                "description": "Get details of all complaints",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "complaints"
                ],
                "summary": "Get details of all complaints",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Complaint"
                            }
                        }
                    }
                }
            }
        },
        "/leases": {
            "get": {
                "description": "Get details of all leases",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leases"
                ],
                "summary": "Get details of all leases",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Lease"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new lease",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leases"
                ],
                "summary": "Create lease",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Lease"
                            }
                        }
                    }
                }
            }
        },
        "/leases/{leaseId}": {
            "get": {
                "description": "Get details of leases",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leases"
                ],
                "summary": "Get details lease",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "leaseId",
                        "name": "leaseId",
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
                                "$ref": "#/definitions/models.Lease"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update the existing lease",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leases"
                ],
                "summary": "Update lease",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Lease"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the existing lease",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leases"
                ],
                "summary": "Delete lease",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Lease"
                            }
                        }
                    }
                }
            }
        },
        "/listings": {
            "get": {
                "description": "Get details of all listings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listings"
                ],
                "summary": "Get details of all listings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Listing"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new listing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listings"
                ],
                "summary": "Create listings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Listing"
                            }
                        }
                    }
                }
            }
        },
        "/listings/{listingId}": {
            "get": {
                "description": "Get details of the listing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listings"
                ],
                "summary": "Get details of listing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "listingId",
                        "name": "leaseId",
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
                                "$ref": "#/definitions/models.Listing"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update the existing listing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listings"
                ],
                "summary": "Update listing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Listing"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the existing lease",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listings"
                ],
                "summary": "Delete listings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Listing"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Create a new token with the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Create a new token",
                "parameters": [
                    {
                        "description": "Create token",
                        "name": "creds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Payload"
                        }
                    }
                }
            }
        },
        "/societies": {
            "get": {
                "description": "Get details of all societies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "societies"
                ],
                "summary": "Get details of all societies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Society"
                            }
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get details of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get details of all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token header",
                        "name": "token",
                        "in": "header",
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
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Credentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.Payload": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
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
        "models.Apartment": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "amenities": {
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
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Complaint": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Lease": {
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
                "leaseEndDate": {
                    "type": "string"
                },
                "leaseStartDate": {
                    "type": "string"
                },
                "listingId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Listing": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "houseType": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isleased": {
                    "type": "boolean"
                },
                "listingImg": {
                    "type": "string"
                },
                "listingType": {
                    "type": "string"
                },
                "rent": {
                    "type": "integer"
                },
                "societyId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Society": {
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
                "societyAddress": {
                    "type": "string"
                },
                "societyAmenities": {
                    "type": "string"
                },
                "societyCity": {
                    "type": "string"
                },
                "societyImg": {
                    "type": "string"
                },
                "societyName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
