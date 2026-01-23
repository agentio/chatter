// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.external

import "github.com/agentio/slink/pkg/common"

// A representation of some externally linked content (eg, a URL and 'card'), embedded in a Bluesky record (eg, a post).
type AppBskyEmbedExternal struct {
	LexiconTypeID string                         `json:"$type,omitempty"`
	External      *AppBskyEmbedExternal_External `json:"external,omitempty"`
}

type AppBskyEmbedExternal_External struct {
	LexiconTypeID string       `json:"$type,omitempty"`
	Description   string       `json:"description,omitempty"`
	Thumb         *common.Blob `json:"thumb,omitempty"`
	Title         string       `json:"title,omitempty"`
	Uri           string       `json:"uri,omitempty"`
}

type AppBskyEmbedExternal_View struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	External      *AppBskyEmbedExternal_ViewExternal `json:"external,omitempty"`
}

type AppBskyEmbedExternal_ViewExternal struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Description   string  `json:"description,omitempty"`
	Thumb         *string `json:"thumb,omitempty"`
	Title         string  `json:"title,omitempty"`
	Uri           string  `json:"uri,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.external",
  "description": "",
  "defs": {
    "external": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "title",
        "description"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "thumb": {
          "type": "blob"
        },
        "title": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "main": {
      "type": "object",
      "description": "A representation of some externally linked content (eg, a URL and 'card'), embedded in a Bluesky record (eg, a post).",
      "required": [
        "external"
      ],
      "properties": {
        "external": {
          "type": "ref",
          "ref": "#external"
        }
      }
    },
    "view": {
      "type": "object",
      "description": "",
      "required": [
        "external"
      ],
      "properties": {
        "external": {
          "type": "ref",
          "ref": "#viewExternal"
        }
      }
    },
    "viewExternal": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "title",
        "description"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "thumb": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    }
  }
}
*/
