// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.getConvoForMembers

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ChatBskyConvoGetConvoForMembers_Description = ""

type ChatBskyConvoGetConvoForMembers_Output struct {
	Convo *ChatBskyConvoDefs_ConvoView `json:"convo,omitempty"`
}

func ChatBskyConvoGetConvoForMembers(ctx context.Context, c slink.Client, members []string) (*ChatBskyConvoGetConvoForMembers_Output, error) {
	var output ChatBskyConvoGetConvoForMembers_Output
	params := map[string]any{
		"members": members,
	}
	if err := c.Do(ctx, slink.Query, "", "chat.bsky.convo.getConvoForMembers", params, nil, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

/*
{
  "lexicon": 1,
  "id": "chat.bsky.convo.getConvoForMembers",
  "description": "",
  "defs": {
    "main": {
      "type": "query",
      "description": "",
      "parameters": {
        "type": "params",
        "properties": {
          "members": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "output": {
        "encoding": "application/json",
        "schema": {
          "type": "object",
          "required": [
            "convo"
          ],
          "properties": {
            "convo": {
              "type": "ref",
              "ref": "chat.bsky.convo.defs#convoView"
            }
          }
        }
      }
    }
  }
}
*/
