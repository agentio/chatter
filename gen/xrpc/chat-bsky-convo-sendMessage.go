// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.sendMessage

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ChatBskyConvoSendMessage_Description = ""

const ChatBskyConvoSendMessage_Input_Description = "Input for ChatBskyConvoSendMessage"

type ChatBskyConvoSendMessage_Input struct {
	ConvoId string                          `json:"convoId"`
	Message *ChatBskyConvoDefs_MessageInput `json:"message,omitempty"`
}

type ChatBskyConvoSendMessage_Output = ChatBskyConvoDefs_MessageView

func ChatBskyConvoSendMessage(ctx context.Context, c slink.Client, input *ChatBskyConvoSendMessage_Input) (*ChatBskyConvoSendMessage_Output, error) {
	var output ChatBskyConvoSendMessage_Output
	if err := c.Do(ctx, slink.Procedure, "application/json", "chat.bsky.convo.sendMessage", nil, input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
