// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.external

import "github.com/agentio/slink/pkg/slink"

const AppBskyEmbedExternal_External_Description = ""

type AppBskyEmbedExternal_External struct {
	LexiconTypeID string      `json:"$type,omitempty"`
	Description   string      `json:"description"`
	Thumb         *slink.Blob `json:"thumb,omitempty"`
	Title         string      `json:"title"`
	Uri           string      `json:"uri"`
}

const AppBskyEmbedExternal_Description = "A representation of some externally linked content (eg, a URL and 'card'), embedded in a Bluesky record (eg, a post)."

type AppBskyEmbedExternal struct {
	LexiconTypeID string                         `json:"$type,omitempty"`
	External      *AppBskyEmbedExternal_External `json:"external,omitempty"`
}

const AppBskyEmbedExternal_View_Description = ""

type AppBskyEmbedExternal_View struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	External      *AppBskyEmbedExternal_ViewExternal `json:"external,omitempty"`
}

const AppBskyEmbedExternal_ViewExternal_Description = ""

type AppBskyEmbedExternal_ViewExternal struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Description   string  `json:"description"`
	Thumb         *string `json:"thumb,omitempty"`
	Title         string  `json:"title"`
	Uri           string  `json:"uri"`
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
