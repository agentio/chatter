// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.recordWithMedia

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/common"
)

type AppBskyEmbedRecordWithMedia_View struct {
	LexiconTypeID string                                                             `json:"$type,omitempty"`
	Media         *AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_View_Media `json:"media,omitempty"`
	Record        *AppBskyEmbedRecord_View                                           `json:"record,omitempty"`
}

type AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_View_Media struct {
	EmbedImages_View   *AppBskyEmbedImages_View
	EmbedVideo_View    *AppBskyEmbedVideo_View
	EmbedExternal_View *AppBskyEmbedExternal_View
}

func (m *AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_View_Media) UnmarshalJSON(data []byte) error {
	recordType := common.LexiconTypeFromJSONBytes(data)
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

func (m AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_View_Media) MarshalJSON() ([]byte, error) {
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

type AppBskyEmbedRecordWithMedia struct {
	LexiconTypeID string                                                        `json:"$type,omitempty"`
	Media         *AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_Media `json:"media,omitempty"`
	Record        *AppBskyEmbedRecord                                           `json:"record,omitempty"`
}

type AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_Media struct {
	EmbedImages   *AppBskyEmbedImages
	EmbedVideo    *AppBskyEmbedVideo
	EmbedExternal *AppBskyEmbedExternal
}

func (m *AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_Media) UnmarshalJSON(data []byte) error {
	recordType := common.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.images":
		m.EmbedImages = &AppBskyEmbedImages{}
		json.Unmarshal(data, m.EmbedImages)
	case "app.bsky.embed.video":
		m.EmbedVideo = &AppBskyEmbedVideo{}
		json.Unmarshal(data, m.EmbedVideo)
	case "app.bsky.embed.external":
		m.EmbedExternal = &AppBskyEmbedExternal{}
		json.Unmarshal(data, m.EmbedExternal)
	}
	return nil
}

func (m AppBskyEmbedRecordWithMediaAppBskyEmbedRecordWithMedia_Media) MarshalJSON() ([]byte, error) {
	if m.EmbedImages != nil {
		return json.Marshal(m.EmbedImages)
	} else if m.EmbedVideo != nil {
		return json.Marshal(m.EmbedVideo)
	} else if m.EmbedExternal != nil {
		return json.Marshal(m.EmbedExternal)
	} else {
		return []byte("{}"), nil
	}
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.embed.recordWithMedia",
  "description": "A representation of a record embedded in a Bluesky record (eg, a post), alongside other compatible embeds. For example, a quote post and image, or a quote post and external URL card.",
  "defs": {
    "main": {
      "type": "object",
      "description": "",
      "required": [
        "record",
        "media"
      ],
      "properties": {
        "media": {
          "type": "union",
          "refs": [
            "app.bsky.embed.images",
            "app.bsky.embed.video",
            "app.bsky.embed.external"
          ]
        },
        "record": {
          "type": "ref",
          "ref": "app.bsky.embed.record"
        }
      }
    },
    "view": {
      "type": "object",
      "description": "",
      "required": [
        "record",
        "media"
      ],
      "properties": {
        "media": {
          "type": "union",
          "refs": [
            "app.bsky.embed.images#view",
            "app.bsky.embed.video#view",
            "app.bsky.embed.external#view"
          ]
        },
        "record": {
          "type": "ref",
          "ref": "app.bsky.embed.record#view"
        }
      }
    }
  }
}
*/
