// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.defs

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const ChatBskyConvoDefs_ConvoView_Description = ""

// ChatBskyConvoDefs_ConvoView is a record with Lexicon type chat.bsky.convo.defs#convoView
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

// ChatBskyConvoDefs_ConvoView_LastMessage is a union with these possible values:
// - ChatBskyConvoDefs_MessageView (chat.bsky.convo.defs#messageView)
// - ChatBskyConvoDefs_DeletedMessageView (chat.bsky.convo.defs#deletedMessageView)
type ChatBskyConvoDefs_ConvoView_LastMessage struct {
	Wrapper ChatBskyConvoDefs_ConvoView_LastMessage_Wrapper
}

// Value wrappers must conform to ChatBskyConvoDefs_ConvoView_LastMessage_Wrapper
type ChatBskyConvoDefs_ConvoView_LastMessage_Wrapper interface {
	isChatBskyConvoDefs_ConvoView_LastMessage()
}

// ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_MessageView wraps values of type *ChatBskyConvoDefs_MessageView
type ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_MessageView struct {
	Value *ChatBskyConvoDefs_MessageView
}

func (t ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_MessageView) isChatBskyConvoDefs_ConvoView_LastMessage() {
}

// ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_DeletedMessageView wraps values of type *ChatBskyConvoDefs_DeletedMessageView
type ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_DeletedMessageView struct {
	Value *ChatBskyConvoDefs_DeletedMessageView
}

func (t ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_DeletedMessageView) isChatBskyConvoDefs_ConvoView_LastMessage() {
}

func (u ChatBskyConvoDefs_ConvoView_LastMessage) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_MessageView:
		return slink.MarshalWithLexiconType("chat.bsky.convo.defs#messageView", v.Value)
	case ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_DeletedMessageView:
		return slink.MarshalWithLexiconType("chat.bsky.convo.defs#deletedMessageView", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *ChatBskyConvoDefs_ConvoView_LastMessage) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "chat.bsky.convo.defs#messageView":
		var v ChatBskyConvoDefs_MessageView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_MessageView{Value: &v}
		return nil
	case "chat.bsky.convo.defs#deletedMessageView":
		var v ChatBskyConvoDefs_DeletedMessageView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = ChatBskyConvoDefs_ConvoView_LastMessage__ChatBskyConvoDefs_DeletedMessageView{Value: &v}
		return nil
	default:
		return nil
	}
}

// ChatBskyConvoDefs_ConvoView_LastReaction is a union with these possible values:
// - ChatBskyConvoDefs_MessageAndReactionView (chat.bsky.convo.defs#messageAndReactionView)
type ChatBskyConvoDefs_ConvoView_LastReaction struct {
	Wrapper ChatBskyConvoDefs_ConvoView_LastReaction_Wrapper
}

// Value wrappers must conform to ChatBskyConvoDefs_ConvoView_LastReaction_Wrapper
type ChatBskyConvoDefs_ConvoView_LastReaction_Wrapper interface {
	isChatBskyConvoDefs_ConvoView_LastReaction()
}

// ChatBskyConvoDefs_ConvoView_LastReaction__ChatBskyConvoDefs_MessageAndReactionView wraps values of type *ChatBskyConvoDefs_MessageAndReactionView
type ChatBskyConvoDefs_ConvoView_LastReaction__ChatBskyConvoDefs_MessageAndReactionView struct {
	Value *ChatBskyConvoDefs_MessageAndReactionView
}

func (t ChatBskyConvoDefs_ConvoView_LastReaction__ChatBskyConvoDefs_MessageAndReactionView) isChatBskyConvoDefs_ConvoView_LastReaction() {
}

func (u ChatBskyConvoDefs_ConvoView_LastReaction) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case ChatBskyConvoDefs_ConvoView_LastReaction__ChatBskyConvoDefs_MessageAndReactionView:
		return slink.MarshalWithLexiconType("chat.bsky.convo.defs#messageAndReactionView", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *ChatBskyConvoDefs_ConvoView_LastReaction) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "chat.bsky.convo.defs#messageAndReactionView":
		var v ChatBskyConvoDefs_MessageAndReactionView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = ChatBskyConvoDefs_ConvoView_LastReaction__ChatBskyConvoDefs_MessageAndReactionView{Value: &v}
		return nil
	default:
		return nil
	}
}

const ChatBskyConvoDefs_DeletedMessageView_Description = ""

// ChatBskyConvoDefs_DeletedMessageView is a record with Lexicon type chat.bsky.convo.defs#deletedMessageView
type ChatBskyConvoDefs_DeletedMessageView struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Id            string                               `json:"id"`
	Rev           string                               `json:"rev"`
	Sender        *ChatBskyConvoDefs_MessageViewSender `json:"sender,omitempty"`
	SentAt        string                               `json:"sentAt"`
}

const ChatBskyConvoDefs_MessageAndReactionView_Description = ""

// ChatBskyConvoDefs_MessageAndReactionView is a record with Lexicon type chat.bsky.convo.defs#messageAndReactionView
type ChatBskyConvoDefs_MessageAndReactionView struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Message       *ChatBskyConvoDefs_MessageView  `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView `json:"reaction,omitempty"`
}

const ChatBskyConvoDefs_MessageInput_Description = ""

// ChatBskyConvoDefs_MessageInput is a record with Lexicon type chat.bsky.convo.defs#messageInput
type ChatBskyConvoDefs_MessageInput struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	Embed         *ChatBskyConvoDefs_MessageInput_Embed `json:"embed,omitempty"`
	Facets        []*AppBskyRichtextFacet               `json:"facets,omitempty"`
	Text          string                                `json:"text"`
}

// ChatBskyConvoDefs_MessageInput_Embed is a union with these possible values:
// - AppBskyEmbedRecord (app.bsky.embed.record)
type ChatBskyConvoDefs_MessageInput_Embed struct {
	Wrapper ChatBskyConvoDefs_MessageInput_Embed_Wrapper
}

// Value wrappers must conform to ChatBskyConvoDefs_MessageInput_Embed_Wrapper
type ChatBskyConvoDefs_MessageInput_Embed_Wrapper interface {
	isChatBskyConvoDefs_MessageInput_Embed()
}

// ChatBskyConvoDefs_MessageInput_Embed__AppBskyEmbedRecord wraps values of type *AppBskyEmbedRecord
type ChatBskyConvoDefs_MessageInput_Embed__AppBskyEmbedRecord struct {
	Value *AppBskyEmbedRecord
}

func (t ChatBskyConvoDefs_MessageInput_Embed__AppBskyEmbedRecord) isChatBskyConvoDefs_MessageInput_Embed() {
}

func (u ChatBskyConvoDefs_MessageInput_Embed) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case ChatBskyConvoDefs_MessageInput_Embed__AppBskyEmbedRecord:
		return slink.MarshalWithLexiconType("app.bsky.embed.record", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *ChatBskyConvoDefs_MessageInput_Embed) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.record":
		var v AppBskyEmbedRecord
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = ChatBskyConvoDefs_MessageInput_Embed__AppBskyEmbedRecord{Value: &v}
		return nil
	default:
		return nil
	}
}

const ChatBskyConvoDefs_MessageView_Description = ""

// ChatBskyConvoDefs_MessageView is a record with Lexicon type chat.bsky.convo.defs#messageView
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

// ChatBskyConvoDefs_MessageView_Embed is a union with these possible values:
// - AppBskyEmbedRecord_View (app.bsky.embed.record#view)
type ChatBskyConvoDefs_MessageView_Embed struct {
	Wrapper ChatBskyConvoDefs_MessageView_Embed_Wrapper
}

// Value wrappers must conform to ChatBskyConvoDefs_MessageView_Embed_Wrapper
type ChatBskyConvoDefs_MessageView_Embed_Wrapper interface {
	isChatBskyConvoDefs_MessageView_Embed()
}

// ChatBskyConvoDefs_MessageView_Embed__AppBskyEmbedRecord_View wraps values of type *AppBskyEmbedRecord_View
type ChatBskyConvoDefs_MessageView_Embed__AppBskyEmbedRecord_View struct {
	Value *AppBskyEmbedRecord_View
}

func (t ChatBskyConvoDefs_MessageView_Embed__AppBskyEmbedRecord_View) isChatBskyConvoDefs_MessageView_Embed() {
}

func (u ChatBskyConvoDefs_MessageView_Embed) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case ChatBskyConvoDefs_MessageView_Embed__AppBskyEmbedRecord_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#view", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *ChatBskyConvoDefs_MessageView_Embed) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.record#view":
		var v AppBskyEmbedRecord_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = ChatBskyConvoDefs_MessageView_Embed__AppBskyEmbedRecord_View{Value: &v}
		return nil
	default:
		return nil
	}
}

const ChatBskyConvoDefs_MessageViewSender_Description = ""

// ChatBskyConvoDefs_MessageViewSender is a record with Lexicon type chat.bsky.convo.defs#messageViewSender
type ChatBskyConvoDefs_MessageViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}

const ChatBskyConvoDefs_ReactionView_Description = ""

// ChatBskyConvoDefs_ReactionView is a record with Lexicon type chat.bsky.convo.defs#reactionView
type ChatBskyConvoDefs_ReactionView struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	CreatedAt     string                                `json:"createdAt"`
	Sender        *ChatBskyConvoDefs_ReactionViewSender `json:"sender,omitempty"`
	Value         string                                `json:"value"`
}

const ChatBskyConvoDefs_ReactionViewSender_Description = ""

// ChatBskyConvoDefs_ReactionViewSender is a record with Lexicon type chat.bsky.convo.defs#reactionViewSender
type ChatBskyConvoDefs_ReactionViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}
