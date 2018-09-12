// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/app"
)

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "myApp",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.0.0",
  "triggers": [
    {
      "id": "receive_http_message",
      "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/rest",
      "name": "Receive HTTP Message",
      "description": "Simple REST Trigger",
      "settings": {
        "port": 9999
      },
      "handlers": [
        {
          "action": {
            "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
            "data": {
              "flowURI": "res://flow:get_books"
            },
            "mappings": {
              "input": [
                {
                  "mapTo": "isbn",
                  "type": "expression",
                  "value": "string.concat(\"isbn:\", $.pathParams.isbn)"
                }
              ],
              "output": [
                {
                  "mapTo": "code",
                  "type": "assign",
                  "value": "$.code"
                },
                {
                  "mapTo": "data",
                  "type": "assign",
                  "value": "$.message"
                }
              ]
            }
          },
          "settings": {
            "method": "GET",
            "path": "/books/:isbn"
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:get_books",
      "data": {
        "name": "GetBooks",
        "metadata": {
          "input": [
            {
              "name": "isbn",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "code",
              "type": "integer"
            },
            {
              "name": "message",
              "type": "any"
            }
          ]
        },
        "tasks": [
          {
            "id": "log_2",
            "name": "Log Message",
            "description": "Simple Log Activity",
            "activity": {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/log",
              "input": {
                "message": "",
                "flowInfo": "false",
                "addToFlow": "false"
              },
              "mappings": {
                "input": [
                  {
                    "type": "expression",
                    "value": "string.concat(\"Getting book data for \", $flow.isbn)",
                    "mapTo": "message"
                  }
                ]
              }
            }
          },
          {
            "id": "rest_3",
            "name": "Invoke REST Service",
            "description": "Simple REST Activity",
            "activity": {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/rest",
              "input": {
                "method": "",
                "uri": "",
                "proxy": "",
                "pathParams": null,
                "queryParams": null,
                "header": null,
                "skipSsl": "false",
                "content": null
              },
              "mappings": {
                "input": [
                  {
                    "type": "literal",
                    "value": "GET",
                    "mapTo": "method"
                  },
                  {
                    "type": "literal",
                    "value": "https://www.googleapis.com/books/v1/volumes",
                    "mapTo": "uri"
                  },
                  {
                    "type": "object",
                    "value": {
                      "q": "{{$flow.isbn}}"
                    },
                    "mapTo": "queryParams"
                  }
                ]
              }
            }
          },
          {
            "id": "actreturn_4",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/actreturn",
              "input": {
                "mappings": [
                  {
                    "mapTo": "code",
                    "type": "literal",
                    "value": 200
                  },
                  {
                    "mapTo": "message",
                    "type": "object",
                    "value": {
                      "title": "{{$activity[rest_3].result.items[0].volumeInfo.title}}",
                      "publishedDate": "{{$activity[rest_3].result.items[0].volumeInfo.publishedDate}}",
                      "description": "{{$activity[rest_3].result.items[0].volumeInfo.description}}"
                    }
                  }
                ]
              }
            }
          }
        ],
        "links": [
          {
            "from": "log_2",
            "to": "rest_3"
          },
          {
            "from": "rest_3",
            "to": "actreturn_4"
          }
        ]
      }
    }
  ]
}
`

func init () {
	cp = EmbeddedProvider()
}

// embeddedConfigProvider implementation of ConfigProvider
type embeddedProvider struct {
}

//EmbeddedProvider returns an app config from a compiled json file
func EmbeddedProvider() (app.ConfigProvider){
	return &embeddedProvider{}
}

// GetApp returns the app configuration
func (d *embeddedProvider) GetApp() (*app.Config, error){

	app := &app.Config{}
	err := json.Unmarshal([]byte(flogoJSON), app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
