// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.postgate

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedPostgate_DisableRule_Description = "Disables embedding of this post."

type AppBskyFeedPostgate_DisableRule struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedPostgate_Description = "Record defining interaction rules for a post. The record key (rkey) of the postgate record must match the record key of the post, and that record must be in the same repository."

type AppBskyFeedPostgate struct {
	LexiconTypeID         string                                                        `json:"$type,omitempty"`
	CreatedAt             string                                                        `json:"createdAt"`
	DetachedEmbeddingUris []string                                                      `json:"detachedEmbeddingUris,omitempty"`
	EmbeddingRules        []*AppBskyFeedPostgateAppBskyFeedPostgate_EmbeddingRules_Elem `json:"embeddingRules,omitempty"`
	Post                  string                                                        `json:"post"`
}

type AppBskyFeedPostgateAppBskyFeedPostgate_EmbeddingRules_Elem struct {
	FeedPostgate_DisableRule *AppBskyFeedPostgate_DisableRule
}

func (m *AppBskyFeedPostgateAppBskyFeedPostgate_EmbeddingRules_Elem) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.feed.postgate#disableRule":
		m.FeedPostgate_DisableRule = &AppBskyFeedPostgate_DisableRule{}
		json.Unmarshal(data, m.FeedPostgate_DisableRule)
	}
	return nil
}

func (m AppBskyFeedPostgateAppBskyFeedPostgate_EmbeddingRules_Elem) MarshalJSON() ([]byte, error) {
	if m.FeedPostgate_DisableRule != nil {
		return json.Marshal(m.FeedPostgate_DisableRule)
	} else {
		return []byte("{}"), nil
	}
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.feed.postgate",
  "description": "",
  "defs": {
    "disableRule": {
      "type": "object",
      "description": "Disables embedding of this post."
    },
    "main": {
      "type": "record",
      "description": "Record defining interaction rules for a post. The record key (rkey) of the postgate record must match the record key of the post, and that record must be in the same repository.",
      "record": {
        "type": "object",
        "required": [
          "post",
          "createdAt"
        ],
        "properties": {
          "createdAt": {
            "type": "string"
          },
          "detachedEmbeddingUris": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "embeddingRules": {
            "type": "array",
            "items": {
              "type": "union",
              "refs": [
                "#disableRule"
              ]
            }
          },
          "post": {
            "type": "string"
          }
        }
      }
    }
  }
}
*/
