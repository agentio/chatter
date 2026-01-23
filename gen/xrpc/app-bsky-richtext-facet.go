// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.richtext.facet

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/common"
)

// Annotation of a sub-string within rich text.
type AppBskyRichtextFacet struct {
	LexiconTypeID string                                                    `json:"$type,omitempty"`
	Features      []*AppBskyRichtextFacetAppBskyRichtextFacet_Features_Elem `json:"features,omitempty"`
	Index         *AppBskyRichtextFacet_ByteSlice                           `json:"index,omitempty"`
}

type AppBskyRichtextFacetAppBskyRichtextFacet_Features_Elem struct {
	RichtextFacet_Mention *AppBskyRichtextFacet_Mention
	RichtextFacet_Link    *AppBskyRichtextFacet_Link
	RichtextFacet_Tag     *AppBskyRichtextFacet_Tag
}

func (m *AppBskyRichtextFacetAppBskyRichtextFacet_Features_Elem) UnmarshalJSON(data []byte) error {
	recordType := common.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.richtext.facet#mention":
		m.RichtextFacet_Mention = &AppBskyRichtextFacet_Mention{}
		json.Unmarshal(data, m.RichtextFacet_Mention)
	case "app.bsky.richtext.facet#link":
		m.RichtextFacet_Link = &AppBskyRichtextFacet_Link{}
		json.Unmarshal(data, m.RichtextFacet_Link)
	case "app.bsky.richtext.facet#tag":
		m.RichtextFacet_Tag = &AppBskyRichtextFacet_Tag{}
		json.Unmarshal(data, m.RichtextFacet_Tag)
	}
	return nil
}

func (m AppBskyRichtextFacetAppBskyRichtextFacet_Features_Elem) MarshalJSON() ([]byte, error) {
	if m.RichtextFacet_Mention != nil {
		return json.Marshal(m.RichtextFacet_Mention)
	} else if m.RichtextFacet_Link != nil {
		return json.Marshal(m.RichtextFacet_Link)
	} else if m.RichtextFacet_Tag != nil {
		return json.Marshal(m.RichtextFacet_Tag)
	} else {
		return []byte("{}"), nil
	}
}

// Facet feature for mention of another account. The text is usually a handle, including a '@' prefix, but the facet reference is a DID.
type AppBskyRichtextFacet_Mention struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did,omitempty"`
}

// Facet feature for a URL. The text URL may have been simplified or truncated, but the facet reference should be a complete URL.
type AppBskyRichtextFacet_Link struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Uri           string `json:"uri,omitempty"`
}

// Facet feature for a hashtag. The text usually includes a '#' prefix, but the facet reference should not (except in the case of 'double hash tags').
type AppBskyRichtextFacet_Tag struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Tag           string `json:"tag,omitempty"`
}

// Specifies the sub-string range a facet feature applies to. Start index is inclusive, end index is exclusive. Indices are zero-indexed, counting bytes of the UTF-8 encoded text. NOTE: some languages, like Javascript, use UTF-16 or Unicode codepoints for string slice indexing; in these languages, convert to byte arrays before working with facets.
type AppBskyRichtextFacet_ByteSlice struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ByteEnd       int64  `json:"byteEnd,omitempty"`
	ByteStart     int64  `json:"byteStart,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.richtext.facet",
  "description": "",
  "defs": {
    "byteSlice": {
      "type": "object",
      "description": "Specifies the sub-string range a facet feature applies to. Start index is inclusive, end index is exclusive. Indices are zero-indexed, counting bytes of the UTF-8 encoded text. NOTE: some languages, like Javascript, use UTF-16 or Unicode codepoints for string slice indexing; in these languages, convert to byte arrays before working with facets.",
      "required": [
        "byteStart",
        "byteEnd"
      ],
      "properties": {
        "byteEnd": {
          "type": "integer"
        },
        "byteStart": {
          "type": "integer"
        }
      }
    },
    "link": {
      "type": "object",
      "description": "Facet feature for a URL. The text URL may have been simplified or truncated, but the facet reference should be a complete URL.",
      "required": [
        "uri"
      ],
      "properties": {
        "uri": {
          "type": "string"
        }
      }
    },
    "main": {
      "type": "object",
      "description": "Annotation of a sub-string within rich text.",
      "required": [
        "index",
        "features"
      ],
      "properties": {
        "features": {
          "type": "array",
          "items": {
            "type": "union",
            "refs": [
              "#mention",
              "#link",
              "#tag"
            ]
          }
        },
        "index": {
          "type": "ref",
          "ref": "#byteSlice"
        }
      }
    },
    "mention": {
      "type": "object",
      "description": "Facet feature for mention of another account. The text is usually a handle, including a '@' prefix, but the facet reference is a DID.",
      "required": [
        "did"
      ],
      "properties": {
        "did": {
          "type": "string"
        }
      }
    },
    "tag": {
      "type": "object",
      "description": "Facet feature for a hashtag. The text usually includes a '#' prefix, but the facet reference should not (except in the case of 'double hash tags').",
      "required": [
        "tag"
      ],
      "properties": {
        "tag": {
          "type": "string"
        }
      }
    }
  }
}
*/
