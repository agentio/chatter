// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.video

import "github.com/agentio/slink/pkg/slink"

type AppBskyEmbedVideo struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           *string                       `json:"alt,omitempty"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Captions      []*AppBskyEmbedVideo_Caption  `json:"captions,omitempty"`
	Video         *slink.Blob                   `json:"video"`
}

type AppBskyEmbedVideo_Caption struct {
	LexiconTypeID string      `json:"$type,omitempty"`
	File          *slink.Blob `json:"file"`
	Lang          string      `json:"lang,omitempty"`
}

type AppBskyEmbedVideo_View struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           *string                       `json:"alt,omitempty"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Cid           string                        `json:"cid,omitempty"`
	Playlist      string                        `json:"playlist,omitempty"`
	Thumbnail     *string                       `json:"thumbnail,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.video",
  "description": "A video embedded in a Bluesky record (eg, a post).",
  "defs": {
    "caption": {
      "type": "object",
      "description": "",
      "required": [
        "lang",
        "file"
      ],
      "properties": {
        "file": {
          "type": "blob"
        },
        "lang": {
          "type": "string"
        }
      }
    },
    "main": {
      "type": "object",
      "description": "",
      "required": [
        "video"
      ],
      "properties": {
        "alt": {
          "type": "string"
        },
        "aspectRatio": {
          "type": "ref",
          "ref": "app.bsky.embed.defs#aspectRatio"
        },
        "captions": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#caption"
          }
        },
        "video": {
          "type": "blob"
        }
      }
    },
    "view": {
      "type": "object",
      "description": "",
      "required": [
        "cid",
        "playlist"
      ],
      "properties": {
        "alt": {
          "type": "string"
        },
        "aspectRatio": {
          "type": "ref",
          "ref": "app.bsky.embed.defs#aspectRatio"
        },
        "cid": {
          "type": "string"
        },
        "playlist": {
          "type": "string"
        },
        "thumbnail": {
          "type": "string"
        }
      }
    }
  }
}
*/
