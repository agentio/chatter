// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.recordWithMedia

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyEmbedRecordWithMedia_View_Description = ""

type AppBskyEmbedRecordWithMedia_View struct {
	LexiconTypeID string                                  `json:"$type,omitempty"`
	Media         *AppBskyEmbedRecordWithMedia_View_Media `json:"media,omitempty"`
	Record        *AppBskyEmbedRecord_View                `json:"record,omitempty"`
}

type AppBskyEmbedRecordWithMedia_View_Media struct {
	EmbedImages_View   *AppBskyEmbedImages_View
	EmbedVideo_View    *AppBskyEmbedVideo_View
	EmbedExternal_View *AppBskyEmbedExternal_View
}

func (m *AppBskyEmbedRecordWithMedia_View_Media) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.images#view":
		m.EmbedImages_View = &AppBskyEmbedImages_View{}
		json.Unmarshal(data, m.EmbedImages_View)
	case "app.bsky.embed.video#view":
		m.EmbedVideo_View = &AppBskyEmbedVideo_View{}
		json.Unmarshal(data, m.EmbedVideo_View)
	case "app.bsky.embed.external#view":
		m.EmbedExternal_View = &AppBskyEmbedExternal_View{}
		json.Unmarshal(data, m.EmbedExternal_View)
	}
	return nil
}

func (m AppBskyEmbedRecordWithMedia_View_Media) MarshalJSON() ([]byte, error) {
	if m.EmbedImages_View != nil {
		return json.Marshal(m.EmbedImages_View)
	} else if m.EmbedVideo_View != nil {
		return json.Marshal(m.EmbedVideo_View)
	} else if m.EmbedExternal_View != nil {
		return json.Marshal(m.EmbedExternal_View)
	} else {
		return []byte("{}"), nil
	}
}
