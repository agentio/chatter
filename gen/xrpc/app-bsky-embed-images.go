// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.images

const AppBskyEmbedImages_View_Description = ""

type AppBskyEmbedImages_View struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Images        []*AppBskyEmbedImages_ViewImage `json:"images,omitempty"`
}

const AppBskyEmbedImages_ViewImage_Description = ""

type AppBskyEmbedImages_ViewImage struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           string                        `json:"alt"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Fullsize      string                        `json:"fullsize"`
	Thumb         string                        `json:"thumb"`
}
