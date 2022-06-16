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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login and get Authorization Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "send request email, password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.User"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Route Path for Get New Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Refresh Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/outpatient": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient with New Medic Record Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "GetAllPatient",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.OutPatientResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Add New Medic Record for Patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "New Medic Record",
                "parameters": [
                    {
                        "description": "new medic record",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminMedicRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/outpatient/doctor": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Process Medic Record by Doctor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "Process Doctor",
                "parameters": [
                    {
                        "description": "process medic record by doctor",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DoctorMedicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/outpatient/nurse": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Process Medic Record by Doctor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OutPatient"
                ],
                "summary": "Process Nurse",
                "parameters": [
                    {
                        "description": "process medic record by nurse",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NurseMedicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "GetAllPatient",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Patient"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "CreatePatient",
                "parameters": [
                    {
                        "description": "patient details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch Patient Data By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "GetPatientByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "patient id",
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
                                    "$ref": "#/definitions/response.MessageData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/models.Patient"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "DeletePatient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "patient id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/patient/:id/update": {
            "put": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Fetch All Patient Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patient"
                ],
                "summary": "UpdatePatient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "patient id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "patient details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.MessageOnly"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/refresh_token": {
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Route Path for Get New Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pass access token here",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "417": {
                        "description": "Expectation Failed",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.MedicRecord": {
            "type": "object",
            "properties": {
                "bloodTension": {
                    "description": "SerialNumber uint",
                    "type": "integer"
                },
                "bodyTemperature": {
                    "type": "integer"
                },
                "complaint": {
                    "type": "string"
                },
                "diagnose": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "medicalStaff": {
                    "$ref": "#/definitions/models.MedicalStaff"
                },
                "medicalStaffID": {
                    "type": "integer"
                },
                "patient": {
                    "$ref": "#/definitions/models.Patient"
                },
                "patientID": {
                    "type": "integer"
                },
                "prescription": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "models.MedicalFacility": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.MedicalStaff": {
            "type": "object",
            "properties": {
                "facility": {
                    "$ref": "#/definitions/models.MedicalFacility"
                },
                "facility_id": {
                    "type": "integer"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Patient": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "blood_type": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medic_record": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MedicRecord"
                    }
                },
                "resident_registration": {
                    "type": "string"
                }
            }
        },
        "models.Role": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "roles": {
                    "$ref": "#/definitions/models.Role"
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "request.AdminMedicRecord": {
            "type": "object",
            "properties": {
                "complaint": {
                    "type": "string"
                },
                "date_check": {
                    "type": "string"
                },
                "medical_facility_id": {
                    "type": "integer"
                },
                "medical_staff_id": {
                    "type": "integer"
                },
                "patient_id": {
                    "type": "integer"
                },
                "session_id": {
                    "type": "integer"
                }
            }
        },
        "request.DoctorMedicRequest": {
            "type": "object",
            "properties": {
                "diagnose": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "prescription": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
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
        "request.NurseMedicRequest": {
            "type": "object",
            "properties": {
                "blood_tesion": {
                    "type": "integer"
                },
                "body_temp": {
                    "type": "integer"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MessageData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MessageOnly": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.OutPatientResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "complaint": {
                    "type": "string"
                },
                "date_check": {
                    "type": "string"
                },
                "doctor": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "queue": {
                    "type": "integer"
                },
                "session": {
                    "type": "string"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "facility": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "roles": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKey": {
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
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Holy Hospital Sever API",
	Description:      "server API for Holy Hospital Application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
