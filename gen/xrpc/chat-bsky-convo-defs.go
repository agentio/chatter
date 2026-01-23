// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.convo.defs

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

type ChatBskyConvoDefs_ReactionViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did,omitempty"`
}

type ChatBskyConvoDefs_MessageAndReactionView struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Message       *ChatBskyConvoDefs_MessageView  `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView `json:"reaction,omitempty"`
}

type ChatBskyConvoDefs_LogBeginConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Rev           string `json:"rev,omitempty"`
}

type ChatBskyConvoDefs_LogAcceptConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Rev           string `json:"rev,omitempty"`
}

type ChatBskyConvoDefs_LogCreateMessage struct {
	LexiconTypeID string                                                       `json:"$type,omitempty"`
	ConvoId       string                                                       `json:"convoId,omitempty"`
	Message       *ChatBskyConvoDefsChatBskyConvoDefs_LogCreateMessage_Message `json:"message,omitempty"`
	Rev           string                                                       `json:"rev,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_LogCreateMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_LogCreateMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_LogCreateMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_LogReadMessage struct {
	LexiconTypeID string                                                     `json:"$type,omitempty"`
	ConvoId       string                                                     `json:"convoId,omitempty"`
	Message       *ChatBskyConvoDefsChatBskyConvoDefs_LogReadMessage_Message `json:"message,omitempty"`
	Rev           string                                                     `json:"rev,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_LogReadMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_LogReadMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_LogReadMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_LogLeaveConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Rev           string `json:"rev,omitempty"`
}

type ChatBskyConvoDefs_LogUnmuteConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Rev           string `json:"rev,omitempty"`
}

type ChatBskyConvoDefs_LogAddReaction struct {
	LexiconTypeID string                                                     `json:"$type,omitempty"`
	ConvoId       string                                                     `json:"convoId,omitempty"`
	Message       *ChatBskyConvoDefsChatBskyConvoDefs_LogAddReaction_Message `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView                            `json:"reaction,omitempty"`
	Rev           string                                                     `json:"rev,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_LogAddReaction_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_LogAddReaction_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_LogAddReaction_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_LogRemoveReaction struct {
	LexiconTypeID string                                                        `json:"$type,omitempty"`
	ConvoId       string                                                        `json:"convoId,omitempty"`
	Message       *ChatBskyConvoDefsChatBskyConvoDefs_LogRemoveReaction_Message `json:"message,omitempty"`
	Reaction      *ChatBskyConvoDefs_ReactionView                               `json:"reaction,omitempty"`
	Rev           string                                                        `json:"rev,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_LogRemoveReaction_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_LogRemoveReaction_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_LogRemoveReaction_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_MessageInput struct {
	LexiconTypeID string                                                 `json:"$type,omitempty"`
	Embed         *ChatBskyConvoDefsChatBskyConvoDefs_MessageInput_Embed `json:"embed,omitempty"`
	Facets        []*AppBskyRichtextFacet                                `json:"facets,omitempty"`
	Text          string                                                 `json:"text,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_MessageInput_Embed struct {
	EmbedRecord *AppBskyEmbedRecord
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_MessageInput_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.record":
		m.EmbedRecord = &AppBskyEmbedRecord{}
		json.Unmarshal(data, m.EmbedRecord)
	}
	return nil
}

func (m ChatBskyConvoDefsChatBskyConvoDefs_MessageInput_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedRecord != nil {
		return json.Marshal(m.EmbedRecord)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_MessageView struct {
	LexiconTypeID string                                                `json:"$type,omitempty"`
	Embed         *ChatBskyConvoDefsChatBskyConvoDefs_MessageView_Embed `json:"embed,omitempty"`
	Facets        []*AppBskyRichtextFacet                               `json:"facets,omitempty"`
	Id            string                                                `json:"id,omitempty"`
	Reactions     []*ChatBskyConvoDefs_ReactionView                     `json:"reactions,omitempty"`
	Rev           string                                                `json:"rev,omitempty"`
	Sender        *ChatBskyConvoDefs_MessageViewSender                  `json:"sender,omitempty"`
	SentAt        string                                                `json:"sentAt,omitempty"`
	Text          string                                                `json:"text,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_MessageView_Embed struct {
	EmbedRecord_View *AppBskyEmbedRecord_View
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_MessageView_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.record#view":
		m.EmbedRecord_View = &AppBskyEmbedRecord_View{}
		json.Unmarshal(data, m.EmbedRecord_View)
	}
	return nil
}

func (m ChatBskyConvoDefsChatBskyConvoDefs_MessageView_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedRecord_View != nil {
		return json.Marshal(m.EmbedRecord_View)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_DeletedMessageView struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Id            string                               `json:"id,omitempty"`
	Rev           string                               `json:"rev,omitempty"`
	Sender        *ChatBskyConvoDefs_MessageViewSender `json:"sender,omitempty"`
	SentAt        string                               `json:"sentAt,omitempty"`
}

type ChatBskyConvoDefs_ReactionView struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	CreatedAt     string                                `json:"createdAt,omitempty"`
	Sender        *ChatBskyConvoDefs_ReactionViewSender `json:"sender,omitempty"`
	Value         string                                `json:"value,omitempty"`
}

type ChatBskyConvoDefs_MessageRef struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Did           string `json:"did,omitempty"`
	MessageId     string `json:"messageId,omitempty"`
}

type ChatBskyConvoDefs_MessageViewSender struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did,omitempty"`
}

type ChatBskyConvoDefs_ConvoView struct {
	LexiconTypeID string                                                     `json:"$type,omitempty"`
	Id            string                                                     `json:"id,omitempty"`
	LastMessage   *ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastMessage  `json:"lastMessage,omitempty"`
	LastReaction  *ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastReaction `json:"lastReaction,omitempty"`
	Members       []*ChatBskyActorDefs_ProfileViewBasic                      `json:"members,omitempty"`
	Muted         bool                                                       `json:"muted"`
	Rev           string                                                     `json:"rev,omitempty"`
	Status        *string                                                    `json:"status,omitempty"`
	UnreadCount   int64                                                      `json:"unreadCount,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastMessage struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastMessage) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastMessage) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastReaction struct {
	ConvoDefs_MessageAndReactionView *ChatBskyConvoDefs_MessageAndReactionView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastReaction) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#messageAndReactionView":
		m.ConvoDefs_MessageAndReactionView = &ChatBskyConvoDefs_MessageAndReactionView{}
		json.Unmarshal(data, m.ConvoDefs_MessageAndReactionView)
	}
	return nil
}

func (m ChatBskyConvoDefsChatBskyConvoDefs_ConvoView_LastReaction) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageAndReactionView != nil {
		return json.Marshal(m.ConvoDefs_MessageAndReactionView)
	} else {
		return []byte("{}"), nil
	}
}

type ChatBskyConvoDefs_LogMuteConvo struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ConvoId       string `json:"convoId,omitempty"`
	Rev           string `json:"rev,omitempty"`
}

type ChatBskyConvoDefs_LogDeleteMessage struct {
	LexiconTypeID string                                                       `json:"$type,omitempty"`
	ConvoId       string                                                       `json:"convoId,omitempty"`
	Message       *ChatBskyConvoDefsChatBskyConvoDefs_LogDeleteMessage_Message `json:"message,omitempty"`
	Rev           string                                                       `json:"rev,omitempty"`
}

type ChatBskyConvoDefsChatBskyConvoDefs_LogDeleteMessage_Message struct {
	ConvoDefs_MessageView        *ChatBskyConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ChatBskyConvoDefs_DeletedMessageView
}

func (m *ChatBskyConvoDefsChatBskyConvoDefs_LogDeleteMessage_Message) UnmarshalJSON(data []byte) error {
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

func (m ChatBskyConvoDefsChatBskyConvoDefs_LogDeleteMessage_Message) MarshalJSON() ([]byte, error) {
	if m.ConvoDefs_MessageView != nil {
		return json.Marshal(m.ConvoDefs_MessageView)
	} else if m.ConvoDefs_DeletedMessageView != nil {
		return json.Marshal(m.ConvoDefs_DeletedMessageView)
	} else {
		return []byte("{}"), nil
	}
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
