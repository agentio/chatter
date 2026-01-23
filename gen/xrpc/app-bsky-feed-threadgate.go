// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.threadgate

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedThreadgate_FollowerRule_Description = "Allow replies from actors who follow you."

type AppBskyFeedThreadgate_FollowerRule struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedThreadgate_FollowingRule_Description = "Allow replies from actors you follow."

type AppBskyFeedThreadgate_FollowingRule struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedThreadgate_ListRule_Description = "Allow replies from actors on a list."

type AppBskyFeedThreadgate_ListRule struct {
	LexiconTypeID string `json:"$type,omitempty"`
	List          string `json:"list"`
}

const AppBskyFeedThreadgate_Description = "Record defining interaction gating rules for a thread (aka, reply controls). The record key (rkey) of the threadgate record must match the record key of the thread's root post, and that record must be in the same repository."

type AppBskyFeedThreadgate struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Allow         []*AppBskyFeedThreadgate_Allow_Elem `json:"allow,omitempty"`
	CreatedAt     string                              `json:"createdAt"`
	HiddenReplies []string                            `json:"hiddenReplies,omitempty"`
	Post          string                              `json:"post"`
}

type AppBskyFeedThreadgate_Allow_Elem struct {
	FeedThreadgate_MentionRule   *AppBskyFeedThreadgate_MentionRule
	FeedThreadgate_FollowerRule  *AppBskyFeedThreadgate_FollowerRule
	FeedThreadgate_FollowingRule *AppBskyFeedThreadgate_FollowingRule
	FeedThreadgate_ListRule      *AppBskyFeedThreadgate_ListRule
}

func (m *AppBskyFeedThreadgate_Allow_Elem) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.feed.threadgate#mentionRule":
		m.FeedThreadgate_MentionRule = &AppBskyFeedThreadgate_MentionRule{}
		json.Unmarshal(data, m.FeedThreadgate_MentionRule)
	case "app.bsky.feed.threadgate#followerRule":
		m.FeedThreadgate_FollowerRule = &AppBskyFeedThreadgate_FollowerRule{}
		json.Unmarshal(data, m.FeedThreadgate_FollowerRule)
	case "app.bsky.feed.threadgate#followingRule":
		m.FeedThreadgate_FollowingRule = &AppBskyFeedThreadgate_FollowingRule{}
		json.Unmarshal(data, m.FeedThreadgate_FollowingRule)
	case "app.bsky.feed.threadgate#listRule":
		m.FeedThreadgate_ListRule = &AppBskyFeedThreadgate_ListRule{}
		json.Unmarshal(data, m.FeedThreadgate_ListRule)
	}
	return nil
}

func (m AppBskyFeedThreadgate_Allow_Elem) MarshalJSON() ([]byte, error) {
	if m.FeedThreadgate_MentionRule != nil {
		return json.Marshal(m.FeedThreadgate_MentionRule)
	} else if m.FeedThreadgate_FollowerRule != nil {
		return json.Marshal(m.FeedThreadgate_FollowerRule)
	} else if m.FeedThreadgate_FollowingRule != nil {
		return json.Marshal(m.FeedThreadgate_FollowingRule)
	} else if m.FeedThreadgate_ListRule != nil {
		return json.Marshal(m.FeedThreadgate_ListRule)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyFeedThreadgate_MentionRule_Description = "Allow replies from actors mentioned in your post."

type AppBskyFeedThreadgate_MentionRule struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.feed.threadgate",
  "description": "",
  "defs": {
    "followerRule": {
      "type": "object",
      "description": "Allow replies from actors who follow you."
    },
    "followingRule": {
      "type": "object",
      "description": "Allow replies from actors you follow."
    },
    "listRule": {
      "type": "object",
      "description": "Allow replies from actors on a list.",
      "required": [
        "list"
      ],
      "properties": {
        "list": {
          "type": "string"
        }
      }
    },
    "main": {
      "type": "record",
      "description": "Record defining interaction gating rules for a thread (aka, reply controls). The record key (rkey) of the threadgate record must match the record key of the thread's root post, and that record must be in the same repository.",
      "record": {
        "type": "object",
        "required": [
          "post",
          "createdAt"
        ],
        "properties": {
          "allow": {
            "type": "array",
            "items": {
              "type": "union",
              "refs": [
                "#mentionRule",
                "#followerRule",
                "#followingRule",
                "#listRule"
              ]
            }
          },
          "createdAt": {
            "type": "string"
          },
          "hiddenReplies": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "post": {
            "type": "string"
          }
        }
      }
    },
    "mentionRule": {
      "type": "object",
      "description": "Allow replies from actors mentioned in your post."
    }
  }
}
*/
