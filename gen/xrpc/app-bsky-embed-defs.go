// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.defs

// width:height represents an aspect ratio. It may be approximate, and may not correspond to absolute dimensions in any given unit.
type AppBskyEmbedDefs_AspectRatio struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Height        int64  `json:"height"`
	Width         int64  `json:"width"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.defs",
  "description": "",
  "defs": {
    "aspectRatio": {
      "type": "object",
      "description": "width:height represents an aspect ratio. It may be approximate, and may not correspond to absolute dimensions in any given unit.",
      "required": [
        "width",
        "height"
      ],
      "properties": {
        "height": {
          "type": "integer",
          "minimum": 1
        },
        "width": {
          "type": "integer",
          "minimum": 1
        }
      }
    }
  }
}
*/
