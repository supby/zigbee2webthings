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
        "summary": "Collection of web things exposed by server.",
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
    "/things/{deviceName}": {
      "get": {
        "description": "Returns a description for device identified by ` + "`" + `id` + "`" + `.\n",
        "summary": "Returns a description for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceName",
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
    "/things/{deviceName}/actions": {
      "get": {
        "summary": "Actions Get or Create",
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
          }
        ],
        "responses": {
          "201": {
            "description": "Created action to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/actions/{actionName}/{actionId}": {
      "get": {
        "summary": "Single Action Status Get",
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "put": {
        "summary": "Update Single Action",
        "parameters": [
          {
            "description": "Action to execute.",
            "name": "action",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionPostBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "delete": {
        "summary": "Cancel created action",
        "responses": {
          "204": {
            "description": "No Content"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Name of action",
          "name": "actionName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Id of created action",
          "name": "actionId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/events": {
      "get": {
        "summary": "Get list of events",
        "responses": {
          "200": {
            "description": "List of EventObject",
            "schema": {
              "$ref": "#/definitions/EventsResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/events/{eventType}": {
      "get": {
        "summary": "Get list of events for particular event type",
        "responses": {
          "200": {
            "description": "List of EventObject",
            "schema": {
              "$ref": "#/definitions/EventsResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Type of event",
          "name": "eventType",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/properties": {
      "get": {
        "summary": "Returns a list of properties for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceName",
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
    "/things/{deviceName}/properties/{propertyName}": {
      "get": {
        "summary": "Returns a property of device.",
        "responses": {
          "200": {
            "description": "A property value",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      },
      "put": {
        "summary": "Set a property of device.",
        "parameters": [
          {
            "description": "Set property.",
            "name": "property",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PropertyPutBody"
            }
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
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
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
      ]
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
    "ActionObject": {
      "type": "object",
      "additionalProperties": {
        "type": "object",
        "properties": {
          "href": {
            "type": "string"
          },
          "id": {
            "type": "string"
          },
          "input": {
            "type": "object",
            "additionalProperties": {
              "type": [
                "string",
                "integer"
              ]
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
    "ActionPostBody": {
      "type": "object",
      "additionalProperties": {
        "description": "Name of action to execute",
        "type": "object",
        "properties": {
          "input": {
            "type": "object",
            "additionalProperties": {
              "type": [
                "string",
                "integer"
              ]
            }
          }
        }
      }
    },
    "ActionsGetResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/ActionObject"
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
    "EventObject": {
      "type": "object",
      "additionalProperties": {
        "type": "object",
        "additionalProperties": {
          "type": [
            "string",
            "number"
          ]
        }
      }
    },
    "EventsResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/EventObject"
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
        "type": [
          "string",
          "number"
        ]
      }
    },
    "PropertyPutBody": {
      "type": "object",
      "additionalProperties": {
        "type": [
          "string",
          "number"
        ]
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
        "summary": "Collection of web things exposed by server.",
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
    "/things/{deviceName}": {
      "get": {
        "description": "Returns a description for device identified by ` + "`" + `id` + "`" + `.\n",
        "summary": "Returns a description for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceName",
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
    "/things/{deviceName}/actions": {
      "get": {
        "summary": "Actions Get or Create",
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
          }
        ],
        "responses": {
          "201": {
            "description": "Created action to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/actions/{actionName}/{actionId}": {
      "get": {
        "summary": "Single Action Status Get",
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "put": {
        "summary": "Update Single Action",
        "parameters": [
          {
            "description": "Action to execute.",
            "name": "action",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ActionPostBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A queue of actions to be executed on a device.\n",
            "schema": {
              "$ref": "#/definitions/ActionObject"
            }
          }
        }
      },
      "delete": {
        "summary": "Cancel created action",
        "responses": {
          "204": {
            "description": "No Content"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Name of action",
          "name": "actionName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Id of created action",
          "name": "actionId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/events": {
      "get": {
        "summary": "Get list of events",
        "responses": {
          "200": {
            "description": "List of EventObject",
            "schema": {
              "$ref": "#/definitions/EventsResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/events/{eventType}": {
      "get": {
        "summary": "Get list of events for particular event type",
        "responses": {
          "200": {
            "description": "List of EventObject",
            "schema": {
              "$ref": "#/definitions/EventsResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "Type of event",
          "name": "eventType",
          "in": "path",
          "required": true
        }
      ]
    },
    "/things/{deviceName}/properties": {
      "get": {
        "summary": "Returns a list of properties for device.",
        "parameters": [
          {
            "type": "string",
            "description": "Id of device",
            "name": "deviceName",
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
    "/things/{deviceName}/properties/{propertyName}": {
      "get": {
        "summary": "Returns a property of device.",
        "responses": {
          "200": {
            "description": "A property value",
            "schema": {
              "$ref": "#/definitions/PropertyGetResponse"
            }
          }
        }
      },
      "put": {
        "summary": "Set a property of device.",
        "parameters": [
          {
            "description": "Set property.",
            "name": "property",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PropertyPutBody"
            }
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
      },
      "parameters": [
        {
          "type": "string",
          "description": "Id of device",
          "name": "deviceName",
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
      ]
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
    "ActionObject": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/ActionObjectAnon"
      }
    },
    "ActionObjectAnon": {
      "type": "object",
      "properties": {
        "href": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "input": {
          "type": "object",
          "additionalProperties": {
            "type": [
              "string",
              "integer"
            ]
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
            "type": [
              "string",
              "integer"
            ]
          }
        }
      }
    },
    "ActionsGetResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/ActionObject"
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
    "EventObject": {
      "type": "object",
      "additionalProperties": {
        "type": "object",
        "additionalProperties": {
          "type": [
            "string",
            "number"
          ]
        }
      }
    },
    "EventsResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/EventObject"
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
        "type": [
          "string",
          "number"
        ]
      }
    },
    "PropertyPutBody": {
      "type": "object",
      "additionalProperties": {
        "type": [
          "string",
          "number"
        ]
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