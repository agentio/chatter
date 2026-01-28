// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.video

const AppBskyEmbedVideo_View_Description = ""

// AppBskyEmbedVideo_View is a record with Lexicon type app.bsky.embed.video#view
type AppBskyEmbedVideo_View struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Alt           *string                       `json:"alt,omitempty"`
	AspectRatio   *AppBskyEmbedDefs_AspectRatio `json:"aspectRatio,omitempty"`
	Cid           string                        `json:"cid"`
	Playlist      string                        `json:"playlist"`
	Thumbnail     *string                       `json:"thumbnail,omitempty"`
}
