// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.external

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
