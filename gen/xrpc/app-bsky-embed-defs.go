// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.defs

const AppBskyEmbedDefs_AspectRatio_Description = "width:height represents an aspect ratio. It may be approximate, and may not correspond to absolute dimensions in any given unit."

// AppBskyEmbedDefs_AspectRatio is a record with Lexicon type app.bsky.embed.defs#aspectRatio
type AppBskyEmbedDefs_AspectRatio struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Height        int64  `json:"height"`
	Width         int64  `json:"width"`
}
