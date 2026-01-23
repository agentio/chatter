// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.video

const AppBskyEmbedVideo_View_Description = ""

type AppBskyEmbedVideo_View struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           *string                       `json:"alt,omitempty"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Cid           string                        `json:"cid"`
	Playlist      string                        `json:"playlist"`
	Thumbnail     *string                       `json:"thumbnail,omitempty"`
}
