// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.defs

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const ChatBskyConvoDefs_ConvoView_Description = ""

type ChatBskyConvoDefs_ConvoView struct {
	LexiconTypeID string                                    `json:"$type,omitempty"`
	Id            string                                    `json:"id"`
	LastMessage   *ChatBskyConvoDefs_ConvoView_LastMessage  `json:"lastMessage,omitempty"`
	LastReaction  *ChatBskyConvoDefs_ConvoView_LastReaction `json:"lastReaction,omitempty"`
	Members       []*ChatBskyActorDefs_ProfileViewBasic     `json:"members,omitempty"`
	Muted         bool                                      `json:"muted"`
	Rev           string                                    `json:"rev"`
	Status        *string                                   `json:"status,omitempty"`
	UnreadCount   int64                                     `json:"unreadCount"`
}

type ChatBskyConvoDefs_ConvoView_LastMessage struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_ConvoView_LastMessage) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#messageView":
		m.ConvoDefs_MessageView = &ChatBskyConvoDefs_MessageView{}
		json.Unmarshal(data, m.ConvoDefs_MessageView)
	case "#deletedMessageView":
		m.ConvoDefs_DeletedMessageView = &ChatBskyConvoDefs_DeletedMessageView{}
		json.Unmarshal(data, m.ConvoDefs_DeletedMessageView)
	}
	return nil
}

func (m ChatBskyConvoDefs_ConvoView_LastMessage) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_ConvoView_LastReaction struct {
	ConvoDefs_MessageAndReactionView *ChatBskyConvoDefs_MessageAndReactionView
}

func (m *ChatBskyConvoDefs_ConvoView_LastReaction) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#messageAndReactionView":
		m.ConvoDefs_MessageAndReactionView = &ChatBskyConvoDefs_MessageAndReactionView{}
		json.Unmarshal(data, m.ConvoDefs_MessageAndReactionView)
	}
	return nil
}

func (m ChatBskyConvoDefs_ConvoView_LastReaction) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageAndReactionView != nil {
		return json.Marshal(m.ConvoDefs_MessageAndReactionView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_DeletedMessageView_Description = ""

type ChatBskyConvoDefs_DeletedMessageView struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Id            string                               `json:"id"`
	Rev           string                               `json:"rev"`
	Sender        *ChatBskyConvoDefs_MessageViewSender `json:"sender,omitempty"`
	SentAt        string                               `json:"sentAt"`
}

const ChatBskyConvoDefs_MessageAndReactionView_Description = ""

type ChatBskyConvoDefs_MessageAndReactionView struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Message       *ChatBskyConvoDefs_MessageView  `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView `json:"reaction,omitempty"`
}

const ChatBskyConvoDefs_MessageInput_Description = ""

type ChatBskyConvoDefs_MessageInput struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	Embed         *ChatBskyConvoDefs_MessageInput_Embed `json:"embed,omitempty"`
	Facets        []*AppBskyRichtextFacet               `json:"facets,omitempty"`
	Text          string                                `json:"text"`
}

type ChatBskyConvoDefs_MessageInput_Embed struct {
	EmbedRecord *AppBskyEmbedRecord
}

func (m *ChatBskyConvoDefs_MessageInput_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.record":
		m.EmbedRecord = &AppBskyEmbedRecord{}
		json.Unmarshal(data, m.EmbedRecord)
	}
	return nil
}

func (m ChatBskyConvoDefs_MessageInput_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedRecord != nil {
		return json.Marshal(m.EmbedRecord)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_MessageView_Description = ""

type ChatBskyConvoDefs_MessageView struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Embed         *ChatBskyConvoDefs_MessageView_Embed `json:"embed,omitempty"`
	Facets        []*AppBskyRichtextFacet              `json:"facets,omitempty"`
	Id            string                               `json:"id"`
	Reactions     []*ChatBskyConvoDefs_ReactionView    `json:"reactions,omitempty"`
	Rev           string                               `json:"rev"`
	Sender        *ChatBskyConvoDefs_MessageViewSender `json:"sender,omitempty"`
	SentAt        string                               `json:"sentAt"`
	Text          string                               `json:"text"`
}

type ChatBskyConvoDefs_MessageView_Embed struct {
	EmbedRecord_View *AppBskyEmbedRecord_View
}

func (m *ChatBskyConvoDefs_MessageView_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.record#view":
		m.EmbedRecord_View = &AppBskyEmbedRecord_View{}
		json.Unmarshal(data, m.EmbedRecord_View)
	}
	return nil
}

func (m ChatBskyConvoDefs_MessageView_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedRecord_View != nil {
		return json.Marshal(m.EmbedRecord_View)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_MessageViewSender_Description = ""

type ChatBskyConvoDefs_MessageViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}

const ChatBskyConvoDefs_ReactionView_Description = ""

type ChatBskyConvoDefs_ReactionView struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	CreatedAt     string                                `json:"createdAt"`
	Sender        *ChatBskyConvoDefs_ReactionViewSender `json:"sender,omitempty"`
	Value         string                                `json:"value"`
}

const ChatBskyConvoDefs_ReactionViewSender_Description = ""

type ChatBskyConvoDefs_ReactionViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}
