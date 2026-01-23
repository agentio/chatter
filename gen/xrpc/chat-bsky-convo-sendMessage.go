// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.sendMessage

import (
	"context"

	"github.com/agentio/slink/pkg/common"
)

const ChatBskyConvoSendMessage_Description = ""

type ChatBskyConvoSendMessage_Input struct {
	ConvoId string                          `json:"convoId,omitempty"`
	Message *ChatBskyConvoDefs_MessageInput `json:"message,omitempty"`
}

type ChatBskyConvoSendMessage_Output = ChatBskyConvoDefs_MessageView

func ChatBskyConvoSendMessage(ctx context.Context, c common.Client, input *ChatBskyConvoSendMessage_Input) (*ChatBskyConvoSendMessage_Output, error) {
	var output ChatBskyConvoSendMessage_Output
	if err := c.Do(ctx, common.Procedure, "application/json", "chat.bsky.convo.sendMessage", nil, input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

/*
{
  "lexicon": 1,
  "id": "chat.bsky.convo.sendMessage",
  "description": "",
  "defs": {
    "main": {
      "type": "procedure",
      "description": "",
      "output": {
        "encoding": "application/json",
        "schema": {
          "type": "ref",
          "ref": "chat.bsky.convo.defs#messageView"
        }
      },
      "input": {
        "encoding": "application/json",
        "schema": {
          "type": "object",
          "required": [
            "convoId",
            "message"
          ],
          "properties": {
            "convoId": {
              "type": "string"
            },
            "message": {
              "type": "ref",
              "ref": "chat.bsky.convo.defs#messageInput"
            }
          }
        }
      }
    }
  }
}
*/
