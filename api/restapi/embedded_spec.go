// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Web Thing API provides a web services based programming interface with a RESTful design for applications to access the properties of devices, request the execution of actions and access a list of events representing a change in state.",
    "title": "Web Thing REST API",
    "version": "0.1.0"
  },
  "host": "api.example.com",
  "basePath": "/v1",
  "paths": {
    "/things": {
      "get": {
        "description": "Returns a list of thngs on this server.",
        "summary": "Returns a list of thngs on this server.",
        "responses": {
          "200": {
            "description": "A JSON array of thing objects",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ThingDescription"
              }
            }
          }
        }
      }
    },
    "/things/{deviceId}": {
      "get": {
        "description": "Returns a description for device identified by ` + "`" + `id` + "`" + `.\n",
        "summary": "Returns a description for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A thing object",
            "schema": {
              "$ref": "#/definitions/ThingDescription"
            }
          }
        }
      }
    },
    "/things/{deviceId}/actions": {
      "get": {
        "summary": "Actions Get or Create",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionsGetResponse"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "description": "Action to execute.",
            "name": "action",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionPostBody"
            }
          },
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created action to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionPostResponse"
            }
          }
        }
      }
    },
    "/things/{deviceId}/properties": {
      "get": {
        "summary": "Returns a list of properties for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A properties object",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      }
    },
    "/things/{deviceId}/properties/{propertyName}": {
      "get": {
        "summary": "Returns a property of device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Name of property",
            "name": "propertyName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A property value",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "input": {
          "type": "object",
          "properties": {
            "@type": {
              "type": "string"
            },
            "maximum": {
              "type": "integer"
            },
            "minimum": {
              "type": "integer"
            },
            "multipleOf": {
              "type": "integer"
            },
            "type": {
              "type": "string"
            },
            "unit": {
              "type": "string"
            }
          }
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "title": {
          "type": "integer"
        }
      }
    },
    "ActionPostBody": {
      "type": "object",
      "additionalProperties": {
        "description": "Name of action to execute",
        "type": "object",
        "properties": {
          "input": {
            "type": "object",
            "additionalProperties": {
              "type": "string"
            }
          }
        }
      }
    },
    "ActionPostResponse": {
      "type": "object",
      "additionalProperties": {
        "type": "object",
        "properties": {
          "href": {
            "type": "string"
          },
          "input": {
            "type": "object",
            "additionalProperties": {
              "type": "string"
            }
          },
          "status": {
            "type": "string"
          },
          "timeRequested": {
            "type": "string"
          }
        }
      }
    },
    "ActionsGetResponse": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": {
          "type": "object",
          "properties": {
            "href": {
              "type": "string"
            },
            "input": {
              "type": "object",
              "additionalProperties": {
                "type": "string"
              }
            },
            "status": {
              "type": "string"
            },
            "timeRequested": {
              "type": "string"
            }
          }
        }
      }
    },
    "EventDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "maximum": {
          "type": "integer"
        },
        "minimum": {
          "type": "integer"
        },
        "multipleOf": {
          "type": "integer"
        },
        "title": {
          "type": "integer"
        },
        "type": {
          "type": "string"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "LinkDescription": {
      "type": "object",
      "properties": {
        "href": {
          "type": "string"
        },
        "mediaType": {
          "type": "string"
        },
        "rel": {
          "type": "string"
        }
      }
    },
    "PropertyDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "enum": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "maximum": {
          "type": "integer"
        },
        "minimum": {
          "type": "integer"
        },
        "multipleOf": {
          "type": "integer"
        },
        "readOnly": {
          "type": "boolean"
        },
        "title": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "PropertyGetResponse": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "ThingDescription": {
      "type": "object",
      "properties": {
        "@context": {
          "type": "string"
        },
        "@type": {
          "type": "string"
        },
        "actions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionDescription"
          }
        },
        "description": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDescription"
          }
        },
        "id": {
          "type": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "properties": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PropertyDescription"
          }
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Web Thing API provides a web services based programming interface with a RESTful design for applications to access the properties of devices, request the execution of actions and access a list of events representing a change in state.",
    "title": "Web Thing REST API",
    "version": "0.1.0"
  },
  "host": "api.example.com",
  "basePath": "/v1",
  "paths": {
    "/things": {
      "get": {
        "description": "Returns a list of thngs on this server.",
        "summary": "Returns a list of thngs on this server.",
        "responses": {
          "200": {
            "description": "A JSON array of thing objects",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ThingDescription"
              }
            }
          }
        }
      }
    },
    "/things/{deviceId}": {
      "get": {
        "description": "Returns a description for device identified by ` + "`" + `id` + "`" + `.\n",
        "summary": "Returns a description for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A thing object",
            "schema": {
              "$ref": "#/definitions/ThingDescription"
            }
          }
        }
      }
    },
    "/things/{deviceId}/actions": {
      "get": {
        "summary": "Actions Get or Create",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionsGetResponse"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "description": "Action to execute.",
            "name": "action",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionPostBody"
            }
          },
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created action to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionPostResponse"
            }
          }
        }
      }
    },
    "/things/{deviceId}/properties": {
      "get": {
        "summary": "Returns a list of properties for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A properties object",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      }
    },
    "/things/{deviceId}/properties/{propertyName}": {
      "get": {
        "summary": "Returns a property of device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Name of property",
            "name": "propertyName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A property value",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "input": {
          "type": "object",
          "properties": {
            "@type": {
              "type": "string"
            },
            "maximum": {
              "type": "integer"
            },
            "minimum": {
              "type": "integer"
            },
            "multipleOf": {
              "type": "integer"
            },
            "type": {
              "type": "string"
            },
            "unit": {
              "type": "string"
            }
          }
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "title": {
          "type": "integer"
        }
      }
    },
    "ActionDescriptionInput": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "maximum": {
          "type": "integer"
        },
        "minimum": {
          "type": "integer"
        },
        "multipleOf": {
          "type": "integer"
        },
        "type": {
          "type": "string"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "ActionPostBody": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/ActionPostBodyAnon"
      }
    },
    "ActionPostBodyAnon": {
      "description": "Name of action to execute",
      "type": "object",
      "properties": {
        "input": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "ActionPostResponse": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/ActionPostResponseAnon"
      }
    },
    "ActionPostResponseAnon": {
      "type": "object",
      "properties": {
        "href": {
          "type": "string"
        },
        "input": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "status": {
          "type": "string"
        },
        "timeRequested": {
          "type": "string"
        }
      }
    },
    "ActionsGetResponse": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": {
          "$ref": "#/definitions/ActionsGetResponseAnon"
        }
      }
    },
    "ActionsGetResponseAnon": {
      "type": "object",
      "properties": {
        "href": {
          "type": "string"
        },
        "input": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "status": {
          "type": "string"
        },
        "timeRequested": {
          "type": "string"
        }
      }
    },
    "EventDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "maximum": {
          "type": "integer"
        },
        "minimum": {
          "type": "integer"
        },
        "multipleOf": {
          "type": "integer"
        },
        "title": {
          "type": "integer"
        },
        "type": {
          "type": "string"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "LinkDescription": {
      "type": "object",
      "properties": {
        "href": {
          "type": "string"
        },
        "mediaType": {
          "type": "string"
        },
        "rel": {
          "type": "string"
        }
      }
    },
    "PropertyDescription": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "enum": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "maximum": {
          "type": "integer"
        },
        "minimum": {
          "type": "integer"
        },
        "multipleOf": {
          "type": "integer"
        },
        "readOnly": {
          "type": "boolean"
        },
        "title": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "PropertyGetResponse": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "ThingDescription": {
      "type": "object",
      "properties": {
        "@context": {
          "type": "string"
        },
        "@type": {
          "type": "string"
        },
        "actions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionDescription"
          }
        },
        "description": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDescription"
          }
        },
        "id": {
          "type": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/LinkDescription"
          }
        },
        "properties": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PropertyDescription"
          }
        },
        "title": {
          "type": "string"
        }
      }
    }
  }
}`))
}