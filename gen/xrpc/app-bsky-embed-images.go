// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.images

import "github.com/agentio/slink/pkg/slink"

type AppBskyEmbedImages_Image struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           string                        `json:"alt"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Image         *slink.Blob                   `json:"image"`
}

type AppBskyEmbedImages struct {
	LexiconTypeID string                      `json:"$type,omitempty"`
	Images        []*AppBskyEmbedImages_Image `json:"images,omitempty"`
}

type AppBskyEmbedImages_View struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Images        []*AppBskyEmbedImages_ViewImage `json:"images,omitempty"`
}

type AppBskyEmbedImages_ViewImage struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           string                        `json:"alt"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Fullsize      string                        `json:"fullsize"`
	Thumb         string                        `json:"thumb"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.images",
  "description": "A set of images embedded in a Bluesky record (eg, a post).",
  "defs": {
    "image": {
      "type": "object",
      "description": "",
      "required": [
        "image",
        "alt"
      ],
      "properties": {
        "alt": {
          "type": "string"
        },
        "aspectRatio": {
          "type": "ref",
          "ref": "app.bsky.embed.defs#aspectRatio"
        },
        "image": {
          "type": "blob"
        }
      }
    },
    "main": {
      "type": "object",
      "description": "",
      "required": [
        "images"
      ],
      "properties": {
        "images": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#image"
          }
        }
      }
    },
    "view": {
      "type": "object",
      "description": "",
      "required": [
        "images"
      ],
      "properties": {
        "images": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#viewImage"
          }
        }
      }
    },
    "viewImage": {
      "type": "object",
      "description": "",
      "required": [
        "thumb",
        "fullsize",
        "alt"
      ],
      "properties": {
        "alt": {
          "type": "string"
        },
        "aspectRatio": {
          "type": "ref",
          "ref": "app.bsky.embed.defs#aspectRatio"
        },
        "fullsize": {
          "type": "string"
        },
        "thumb": {
          "type": "string"
        }
      }
    }
  }
}
*/
