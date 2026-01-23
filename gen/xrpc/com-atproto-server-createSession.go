// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.server.createSession

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ServerCreateSession_Description = "Create an authentication session."

type ServerCreateSession_Input struct {
	AllowTakendown  *bool   `json:"allowTakendown,omitempty"`
	AuthFactorToken *string `json:"authFactorToken,omitempty"`
	Identifier      string  `json:"identifier"`
	Password        string  `json:"password"`
}

type ServerCreateSession_Output struct {
	AccessJwt       string  `json:"accessJwt"`
	Active          *bool   `json:"active,omitempty"`
	Did             string  `json:"did"`
	DidDoc          *any    `json:"didDoc,omitempty"`
	Email           *string `json:"email,omitempty"`
	EmailAuthFactor *bool   `json:"emailAuthFactor,omitempty"`
	EmailConfirmed  *bool   `json:"emailConfirmed,omitempty"`
	Handle          string  `json:"handle"`
	RefreshJwt      string  `json:"refreshJwt"`
	Status          *string `json:"status,omitempty"`
}

// Create an authentication session.
func ServerCreateSession(ctx context.Context, c slink.Client, input *ServerCreateSession_Input) (*ServerCreateSession_Output, error) {
	var output ServerCreateSession_Output
	if err := c.Do(ctx, slink.Procedure, "application/json", "com.atproto.server.createSession", nil, input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

/*
{
  "lexicon": 1,
  "id": "com.atproto.server.createSession",
  "description": "",
  "defs": {
    "main": {
      "type": "procedure",
      "description": "Create an authentication session.",
      "output": {
        "encoding": "application/json",
        "schema": {
          "type": "object",
          "required": [
            "accessJwt",
            "refreshJwt",
            "handle",
            "did"
          ],
          "properties": {
            "accessJwt": {
              "type": "string"
            },
            "active": {
              "type": "boolean"
            },
            "did": {
              "type": "string"
            },
            "didDoc": {
              "type": "unknown"
            },
            "email": {
              "type": "string"
            },
            "emailAuthFactor": {
              "type": "boolean"
            },
            "emailConfirmed": {
              "type": "boolean"
            },
            "handle": {
              "type": "string"
            },
            "refreshJwt": {
              "type": "string"
            },
            "status": {
              "type": "string"
            }
          }
        }
      },
      "input": {
        "encoding": "application/json",
        "schema": {
          "type": "object",
          "required": [
            "identifier",
            "password"
          ],
          "properties": {
            "allowTakendown": {
              "type": "boolean"
            },
            "authFactorToken": {
              "type": "string"
            },
            "identifier": {
              "type": "string"
            },
            "password": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}
*/
