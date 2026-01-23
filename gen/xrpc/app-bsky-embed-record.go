// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.record

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/common"
)

type AppBskyEmbedRecord struct {
	LexiconTypeID string         `json:"$type,omitempty"`
	Record        *RepoStrongRef `json:"record,omitempty"`
}

type AppBskyEmbedRecord_View struct {
	LexiconTypeID string                                            `json:"$type,omitempty"`
	Record        *AppBskyEmbedRecordAppBskyEmbedRecord_View_Record `json:"record,omitempty"`
}

type AppBskyEmbedRecordAppBskyEmbedRecord_View_Record struct {
	EmbedRecord_ViewRecord         *AppBskyEmbedRecord_ViewRecord
	EmbedRecord_ViewNotFound       *AppBskyEmbedRecord_ViewNotFound
	EmbedRecord_ViewBlocked        *AppBskyEmbedRecord_ViewBlocked
	EmbedRecord_ViewDetached       *AppBskyEmbedRecord_ViewDetached
	FeedDefs_GeneratorView         *AppBskyFeedDefs_GeneratorView
	GraphDefs_ListView             *AppBskyGraphDefs_ListView
	LabelerDefs_LabelerView        *AppBskyLabelerDefs_LabelerView
	GraphDefs_StarterPackViewBasic *AppBskyGraphDefs_StarterPackViewBasic
}

func (m *AppBskyEmbedRecordAppBskyEmbedRecord_View_Record) UnmarshalJSON(data []byte) error {
	recordType := common.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#viewRecord":
		m.EmbedRecord_ViewRecord = &AppBskyEmbedRecord_ViewRecord{}
		json.Unmarshal(data, m.EmbedRecord_ViewRecord)
	case "#viewNotFound":
		m.EmbedRecord_ViewNotFound = &AppBskyEmbedRecord_ViewNotFound{}
		json.Unmarshal(data, m.EmbedRecord_ViewNotFound)
	case "#viewBlocked":
		m.EmbedRecord_ViewBlocked = &AppBskyEmbedRecord_ViewBlocked{}
		json.Unmarshal(data, m.EmbedRecord_ViewBlocked)
	case "#viewDetached":
		m.EmbedRecord_ViewDetached = &AppBskyEmbedRecord_ViewDetached{}
		json.Unmarshal(data, m.EmbedRecord_ViewDetached)
	case "app.bsky.feed.defs#generatorView":
		m.FeedDefs_GeneratorView = &AppBskyFeedDefs_GeneratorView{}
		json.Unmarshal(data, m.FeedDefs_GeneratorView)
	case "app.bsky.graph.defs#listView":
		m.GraphDefs_ListView = &AppBskyGraphDefs_ListView{}
		json.Unmarshal(data, m.GraphDefs_ListView)
	case "app.bsky.labeler.defs#labelerView":
		m.LabelerDefs_LabelerView = &AppBskyLabelerDefs_LabelerView{}
		json.Unmarshal(data, m.LabelerDefs_LabelerView)
	case "app.bsky.graph.defs#starterPackViewBasic":
		m.GraphDefs_StarterPackViewBasic = &AppBskyGraphDefs_StarterPackViewBasic{}
		json.Unmarshal(data, m.GraphDefs_StarterPackViewBasic)
	}
	return nil
}

func (m AppBskyEmbedRecordAppBskyEmbedRecord_View_Record) MarshalJSON() ([]byte, error) {
	if m.EmbedRecord_ViewRecord != nil {
		return json.Marshal(m.EmbedRecord_ViewRecord)
	} else if m.EmbedRecord_ViewNotFound != nil {
		return json.Marshal(m.EmbedRecord_ViewNotFound)
	} else if m.EmbedRecord_ViewBlocked != nil {
		return json.Marshal(m.EmbedRecord_ViewBlocked)
	} else if m.EmbedRecord_ViewDetached != nil {
		return json.Marshal(m.EmbedRecord_ViewDetached)
	} else if m.FeedDefs_GeneratorView != nil {
		return json.Marshal(m.FeedDefs_GeneratorView)
	} else if m.GraphDefs_ListView != nil {
		return json.Marshal(m.GraphDefs_ListView)
	} else if m.LabelerDefs_LabelerView != nil {
		return json.Marshal(m.LabelerDefs_LabelerView)
	} else if m.GraphDefs_StarterPackViewBasic != nil {
		return json.Marshal(m.GraphDefs_StarterPackViewBasic)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyEmbedRecord_ViewRecord struct {
	LexiconTypeID string                                                         `json:"$type,omitempty"`
	Author        *AppBskyActorDefs_ProfileViewBasic                             `json:"author,omitempty"`
	Cid           string                                                         `json:"cid,omitempty"`
	Embeds        []*AppBskyEmbedRecordAppBskyEmbedRecord_ViewRecord_Embeds_Elem `json:"embeds,omitempty"`
	IndexedAt     string                                                         `json:"indexedAt,omitempty"`
	Labels        []*LabelDefs_Label                                             `json:"labels,omitempty"`
	LikeCount     *int64                                                         `json:"likeCount,omitempty"`
	QuoteCount    *int64                                                         `json:"quoteCount,omitempty"`
	ReplyCount    *int64                                                         `json:"replyCount,omitempty"`
	RepostCount   *int64                                                         `json:"repostCount,omitempty"`
	Uri           string                                                         `json:"uri,omitempty"`
	Value         any                                                            `json:"value,omitempty"`
}

type AppBskyEmbedRecordAppBskyEmbedRecord_ViewRecord_Embeds_Elem struct {
	EmbedImages_View          *AppBskyEmbedImages_View
	EmbedVideo_View           *AppBskyEmbedVideo_View
	EmbedExternal_View        *AppBskyEmbedExternal_View
	EmbedRecord_View          *AppBskyEmbedRecord_View
	EmbedRecordWithMedia_View *AppBskyEmbedRecordWithMedia_View
}

func (m *AppBskyEmbedRecordAppBskyEmbedRecord_ViewRecord_Embeds_Elem) UnmarshalJSON(data []byte) error {
	recordType := common.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.images#view":
		m.EmbedImages_View = &AppBskyEmbedImages_View{}
		json.Unmarshal(data, m.EmbedImages_View)
	case "app.bsky.embed.video#view":
		m.EmbedVideo_View = &AppBskyEmbedVideo_View{}
		json.Unmarshal(data, m.EmbedVideo_View)
	case "app.bsky.embed.external#view":
		m.EmbedExternal_View = &AppBskyEmbedExternal_View{}
		json.Unmarshal(data, m.EmbedExternal_View)
	case "app.bsky.embed.record#view":
		m.EmbedRecord_View = &AppBskyEmbedRecord_View{}
		json.Unmarshal(data, m.EmbedRecord_View)
	case "app.bsky.embed.recordWithMedia#view":
		m.EmbedRecordWithMedia_View = &AppBskyEmbedRecordWithMedia_View{}
		json.Unmarshal(data, m.EmbedRecordWithMedia_View)
	}
	return nil
}

func (m AppBskyEmbedRecordAppBskyEmbedRecord_ViewRecord_Embeds_Elem) MarshalJSON() ([]byte, error) {
	if m.EmbedImages_View != nil {
		return json.Marshal(m.EmbedImages_View)
	} else if m.EmbedVideo_View != nil {
		return json.Marshal(m.EmbedVideo_View)
	} else if m.EmbedExternal_View != nil {
		return json.Marshal(m.EmbedExternal_View)
	} else if m.EmbedRecord_View != nil {
		return json.Marshal(m.EmbedRecord_View)
	} else if m.EmbedRecordWithMedia_View != nil {
		return json.Marshal(m.EmbedRecordWithMedia_View)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyEmbedRecord_ViewNotFound struct {
	LexiconTypeID string `json:"$type,omitempty"`
	NotFound      bool   `json:"notFound"`
	Uri           string `json:"uri,omitempty"`
}

type AppBskyEmbedRecord_ViewBlocked struct {
	LexiconTypeID string                         `json:"$type,omitempty"`
	Author        *AppBskyFeedDefs_BlockedAuthor `json:"author,omitempty"`
	Blocked       bool                           `json:"blocked"`
	Uri           string                         `json:"uri,omitempty"`
}

type AppBskyEmbedRecord_ViewDetached struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Detached      bool   `json:"detached"`
	Uri           string `json:"uri,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.record",
  "description": "A representation of a record embedded in a Bluesky record (eg, a post). For example, a quote-post, or sharing a feed generator record.",
  "defs": {
    "main": {
      "type": "object",
      "description": "",
      "required": [
        "record"
      ],
      "properties": {
        "record": {
          "type": "ref",
          "ref": "com.atproto.repo.strongRef"
        }
      }
    },
    "view": {
      "type": "object",
      "description": "",
      "required": [
        "record"
      ],
      "properties": {
        "record": {
          "type": "union",
          "refs": [
            "#viewRecord",
            "#viewNotFound",
            "#viewBlocked",
            "#viewDetached",
            "app.bsky.feed.defs#generatorView",
            "app.bsky.graph.defs#listView",
            "app.bsky.labeler.defs#labelerView",
            "app.bsky.graph.defs#starterPackViewBasic"
          ]
        }
      }
    },
    "viewBlocked": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "blocked",
        "author"
      ],
      "properties": {
        "author": {
          "type": "ref",
          "ref": "app.bsky.feed.defs#blockedAuthor"
        },
        "blocked": {
          "type": "boolean"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "viewDetached": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "detached"
      ],
      "properties": {
        "detached": {
          "type": "boolean"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "viewNotFound": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "notFound"
      ],
      "properties": {
        "notFound": {
          "type": "boolean"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "viewRecord": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "author",
        "value",
        "indexedAt"
      ],
      "properties": {
        "author": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "cid": {
          "type": "string"
        },
        "embeds": {
          "type": "array",
          "items": {
            "type": "union",
            "refs": [
              "app.bsky.embed.images#view",
              "app.bsky.embed.video#view",
              "app.bsky.embed.external#view",
              "app.bsky.embed.record#view",
              "app.bsky.embed.recordWithMedia#view"
            ]
          }
        },
        "indexedAt": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "likeCount": {
          "type": "integer"
        },
        "quoteCount": {
          "type": "integer"
        },
        "replyCount": {
          "type": "integer"
        },
        "repostCount": {
          "type": "integer"
        },
        "uri": {
          "type": "string"
        },
        "value": {
          "type": "unknown"
        }
      }
    }
  }
}
*/
