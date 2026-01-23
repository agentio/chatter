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

const ChatBskyConvoDefs_LogAcceptConvo_Description = ""

type ChatBskyConvoDefs_LogAcceptConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Rev           string `json:"rev"`
}

const ChatBskyConvoDefs_LogAddReaction_Description = ""

type ChatBskyConvoDefs_LogAddReaction struct {
	LexiconTypeID string                                    `json:"$type,omitempty"`
	ConvoId       string                                    `json:"convoId"`
	Message       *ChatBskyConvoDefs_LogAddReaction_Message `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView           `json:"reaction,omitempty"`
	Rev           string                                    `json:"rev"`
}

type ChatBskyConvoDefs_LogAddReaction_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_LogAddReaction_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefs_LogAddReaction_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_LogBeginConvo_Description = ""

type ChatBskyConvoDefs_LogBeginConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Rev           string `json:"rev"`
}

const ChatBskyConvoDefs_LogCreateMessage_Description = ""

type ChatBskyConvoDefs_LogCreateMessage struct {
	LexiconTypeID string                                      `json:"$type,omitempty"`
	ConvoId       string                                      `json:"convoId"`
	Message       *ChatBskyConvoDefs_LogCreateMessage_Message `json:"message,omitempty"`
	Rev           string                                      `json:"rev"`
}

type ChatBskyConvoDefs_LogCreateMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_LogCreateMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefs_LogCreateMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_LogDeleteMessage_Description = ""

type ChatBskyConvoDefs_LogDeleteMessage struct {
	LexiconTypeID string                                      `json:"$type,omitempty"`
	ConvoId       string                                      `json:"convoId"`
	Message       *ChatBskyConvoDefs_LogDeleteMessage_Message `json:"message,omitempty"`
	Rev           string                                      `json:"rev"`
}

type ChatBskyConvoDefs_LogDeleteMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_LogDeleteMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefs_LogDeleteMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_LogLeaveConvo_Description = ""

type ChatBskyConvoDefs_LogLeaveConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Rev           string `json:"rev"`
}

const ChatBskyConvoDefs_LogMuteConvo_Description = ""

type ChatBskyConvoDefs_LogMuteConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Rev           string `json:"rev"`
}

const ChatBskyConvoDefs_LogReadMessage_Description = ""

type ChatBskyConvoDefs_LogReadMessage struct {
	LexiconTypeID string                                    `json:"$type,omitempty"`
	ConvoId       string                                    `json:"convoId"`
	Message       *ChatBskyConvoDefs_LogReadMessage_Message `json:"message,omitempty"`
	Rev           string                                    `json:"rev"`
}

type ChatBskyConvoDefs_LogReadMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_LogReadMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefs_LogReadMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_LogRemoveReaction_Description = ""

type ChatBskyConvoDefs_LogRemoveReaction struct {
	LexiconTypeID string                                       `json:"$type,omitempty"`
	ConvoId       string                                       `json:"convoId"`
	Message       *ChatBskyConvoDefs_LogRemoveReaction_Message `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView              `json:"reaction,omitempty"`
	Rev           string                                       `json:"rev"`
}

type ChatBskyConvoDefs_LogRemoveReaction_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefs_LogRemoveReaction_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefs_LogRemoveReaction_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

const ChatBskyConvoDefs_LogUnmuteConvo_Description = ""

type ChatBskyConvoDefs_LogUnmuteConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Rev           string `json:"rev"`
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

const ChatBskyConvoDefs_MessageRef_Description = ""

type ChatBskyConvoDefs_MessageRef struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId"`
	Did           string `json:"did"`
	MessageId     string `json:"messageId"`
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

/*
{
  "lexicon": 1,
  "id": "chat.bsky.convo.defs",
  "description": "",
  "defs": {
    "convoView": {
      "type": "object",
      "description": "",
      "required": [
        "id",
        "rev",
        "members",
        "muted",
        "unreadCount"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "lastMessage": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "lastReaction": {
          "type": "union",
          "refs": [
            "#messageAndReactionView"
          ]
        },
        "members": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "chat.bsky.actor.defs#profileViewBasic"
          }
        },
        "muted": {
          "type": "boolean"
        },
        "rev": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "unreadCount": {
          "type": "integer"
        }
      }
    },
    "deletedMessageView": {
      "type": "object",
      "description": "",
      "required": [
        "id",
        "rev",
        "sender",
        "sentAt"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        },
        "sender": {
          "type": "ref",
          "ref": "#messageViewSender"
        },
        "sentAt": {
          "type": "string"
        }
      }
    },
    "logAcceptConvo": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logAddReaction": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId",
        "message",
        "reaction"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "message": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "reaction": {
          "type": "ref",
          "ref": "#reactionView"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logBeginConvo": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logCreateMessage": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId",
        "message"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "message": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logDeleteMessage": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId",
        "message"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "message": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logLeaveConvo": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logMuteConvo": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logReadMessage": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId",
        "message"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "message": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logRemoveReaction": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId",
        "message",
        "reaction"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "message": {
          "type": "union",
          "refs": [
            "#messageView",
            "#deletedMessageView"
          ]
        },
        "reaction": {
          "type": "ref",
          "ref": "#reactionView"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "logUnmuteConvo": {
      "type": "object",
      "description": "",
      "required": [
        "rev",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "rev": {
          "type": "string"
        }
      }
    },
    "messageAndReactionView": {
      "type": "object",
      "description": "",
      "required": [
        "message",
        "reaction"
      ],
      "properties": {
        "message": {
          "type": "ref",
          "ref": "#messageView"
        },
        "reaction": {
          "type": "ref",
          "ref": "#reactionView"
        }
      }
    },
    "messageInput": {
      "type": "object",
      "description": "",
      "required": [
        "text"
      ],
      "properties": {
        "embed": {
          "type": "union",
          "refs": [
            "app.bsky.embed.record"
          ]
        },
        "facets": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.richtext.facet"
          }
        },
        "text": {
          "type": "string"
        }
      }
    },
    "messageRef": {
      "type": "object",
      "description": "",
      "required": [
        "did",
        "messageId",
        "convoId"
      ],
      "properties": {
        "convoId": {
          "type": "string"
        },
        "did": {
          "type": "string"
        },
        "messageId": {
          "type": "string"
        }
      }
    },
    "messageView": {
      "type": "object",
      "description": "",
      "required": [
        "id",
        "rev",
        "text",
        "sender",
        "sentAt"
      ],
      "properties": {
        "embed": {
          "type": "union",
          "refs": [
            "app.bsky.embed.record#view"
          ]
        },
        "facets": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.richtext.facet"
          }
        },
        "id": {
          "type": "string"
        },
        "reactions": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#reactionView"
          }
        },
        "rev": {
          "type": "string"
        },
        "sender": {
          "type": "ref",
          "ref": "#messageViewSender"
        },
        "sentAt": {
          "type": "string"
        },
        "text": {
          "type": "string"
        }
      }
    },
    "messageViewSender": {
      "type": "object",
      "description": "",
      "required": [
        "did"
      ],
      "properties": {
        "did": {
          "type": "string"
        }
      }
    },
    "reactionView": {
      "type": "object",
      "description": "",
      "required": [
        "value",
        "sender",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "sender": {
          "type": "ref",
          "ref": "#reactionViewSender"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "reactionViewSender": {
      "type": "object",
      "description": "",
      "required": [
        "did"
      ],
      "properties": {
        "did": {
          "type": "string"
        }
      }
    }
  }
}
*/
