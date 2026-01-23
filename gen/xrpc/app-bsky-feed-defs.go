// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.defs

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedDefs_BlockedAuthor_Description = ""

type AppBskyFeedDefs_BlockedAuthor struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Did           string                        `json:"did"`
	Viewer        *AppBskyActorDefs_ViewerState `json:"viewer,omitempty"`
}

const AppBskyFeedDefs_BlockedPost_Description = ""

type AppBskyFeedDefs_BlockedPost struct {
	LexiconTypeID string                         `json:"$type,omitempty"`
	Author        *AppBskyFeedDefs_BlockedAuthor `json:"author,omitempty"`
	Blocked       bool                           `json:"blocked"`
	Uri           string                         `json:"uri"`
}

// User clicked through to the author of the feed item
const AppBskyFeedDefs_ClickthroughAuthor string = "clickthroughAuthor"

// User clicked through to the embedded content of the feed item
const AppBskyFeedDefs_ClickthroughEmbed string = "clickthroughEmbed"

// User clicked through to the feed item
const AppBskyFeedDefs_ClickthroughItem string = "clickthroughItem"

// User clicked through to the reposter of the feed item
const AppBskyFeedDefs_ClickthroughReposter string = "clickthroughReposter"

// Declares the feed generator returns any types of posts.
const AppBskyFeedDefs_ContentModeUnspecified string = "contentModeUnspecified"

// Declares the feed generator returns posts containing app.bsky.embed.video embeds.
const AppBskyFeedDefs_ContentModeVideo string = "contentModeVideo"

const AppBskyFeedDefs_FeedViewPost_Description = ""

type AppBskyFeedDefs_FeedViewPost struct {
	LexiconTypeID string                                              `json:"$type,omitempty"`
	FeedContext   *string                                             `json:"feedContext,omitempty"`
	Post          *AppBskyFeedDefs_PostView                           `json:"post,omitempty"`
	Reason        *AppBskyFeedDefsAppBskyFeedDefs_FeedViewPost_Reason `json:"reason,omitempty"`
	Reply         *AppBskyFeedDefs_ReplyRef                           `json:"reply,omitempty"`
	ReqId         *string                                             `json:"reqId,omitempty"`
}

type AppBskyFeedDefsAppBskyFeedDefs_FeedViewPost_Reason struct {
	FeedDefs_ReasonRepost *AppBskyFeedDefs_ReasonRepost
	FeedDefs_ReasonPin    *AppBskyFeedDefs_ReasonPin
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_FeedViewPost_Reason) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#reasonRepost":
		m.FeedDefs_ReasonRepost = &AppBskyFeedDefs_ReasonRepost{}
		json.Unmarshal(data, m.FeedDefs_ReasonRepost)
	case "#reasonPin":
		m.FeedDefs_ReasonPin = &AppBskyFeedDefs_ReasonPin{}
		json.Unmarshal(data, m.FeedDefs_ReasonPin)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_FeedViewPost_Reason) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_ReasonRepost != nil {
		return json.Marshal(m.FeedDefs_ReasonRepost)
	} else if m.FeedDefs_ReasonPin != nil {
		return json.Marshal(m.FeedDefs_ReasonPin)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyFeedDefs_GeneratorView_Description = ""

type AppBskyFeedDefs_GeneratorView struct {
	LexiconTypeID       string                                `json:"$type,omitempty"`
	AcceptsInteractions *bool                                 `json:"acceptsInteractions,omitempty"`
	Avatar              *string                               `json:"avatar,omitempty"`
	Cid                 string                                `json:"cid"`
	ContentMode         *string                               `json:"contentMode,omitempty"`
	Creator             *AppBskyActorDefs_ProfileView         `json:"creator,omitempty"`
	Description         *string                               `json:"description,omitempty"`
	DescriptionFacets   []*AppBskyRichtextFacet               `json:"descriptionFacets,omitempty"`
	Did                 string                                `json:"did"`
	DisplayName         string                                `json:"displayName"`
	IndexedAt           string                                `json:"indexedAt"`
	Labels              []*LabelDefs_Label                    `json:"labels,omitempty"`
	LikeCount           *int64                                `json:"likeCount,omitempty"`
	Uri                 string                                `json:"uri"`
	Viewer              *AppBskyFeedDefs_GeneratorViewerState `json:"viewer,omitempty"`
}

const AppBskyFeedDefs_GeneratorViewerState_Description = ""

type AppBskyFeedDefs_GeneratorViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Like          *string `json:"like,omitempty"`
}

const AppBskyFeedDefs_Interaction_Description = ""

type AppBskyFeedDefs_Interaction struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Event         *string `json:"event,omitempty"`
	FeedContext   *string `json:"feedContext,omitempty"`
	Item          *string `json:"item,omitempty"`
	ReqId         *string `json:"reqId,omitempty"`
}

// User liked the feed item
const AppBskyFeedDefs_InteractionLike string = "interactionLike"

// User quoted the feed item
const AppBskyFeedDefs_InteractionQuote string = "interactionQuote"

// User replied to the feed item
const AppBskyFeedDefs_InteractionReply string = "interactionReply"

// User reposted the feed item
const AppBskyFeedDefs_InteractionRepost string = "interactionRepost"

// Feed item was seen by user
const AppBskyFeedDefs_InteractionSeen string = "interactionSeen"

// User shared the feed item
const AppBskyFeedDefs_InteractionShare string = "interactionShare"

const AppBskyFeedDefs_NotFoundPost_Description = ""

type AppBskyFeedDefs_NotFoundPost struct {
	LexiconTypeID string `json:"$type,omitempty"`
	NotFound      bool   `json:"notFound"`
	Uri           string `json:"uri"`
}

const AppBskyFeedDefs_PostView_Description = ""

type AppBskyFeedDefs_PostView struct {
	LexiconTypeID string                                         `json:"$type,omitempty"`
	Author        *AppBskyActorDefs_ProfileViewBasic             `json:"author,omitempty"`
	BookmarkCount *int64                                         `json:"bookmarkCount,omitempty"`
	Cid           string                                         `json:"cid"`
	Debug         *any                                           `json:"debug,omitempty"`
	Embed         *AppBskyFeedDefsAppBskyFeedDefs_PostView_Embed `json:"embed,omitempty"`
	IndexedAt     string                                         `json:"indexedAt"`
	Labels        []*LabelDefs_Label                             `json:"labels,omitempty"`
	LikeCount     *int64                                         `json:"likeCount,omitempty"`
	QuoteCount    *int64                                         `json:"quoteCount,omitempty"`
	Record        any                                            `json:"record"`
	ReplyCount    *int64                                         `json:"replyCount,omitempty"`
	RepostCount   *int64                                         `json:"repostCount,omitempty"`
	Threadgate    *AppBskyFeedDefs_ThreadgateView                `json:"threadgate,omitempty"`
	Uri           string                                         `json:"uri"`
	Viewer        *AppBskyFeedDefs_ViewerState                   `json:"viewer,omitempty"`
}

type AppBskyFeedDefsAppBskyFeedDefs_PostView_Embed struct {
	EmbedImages_View          *AppBskyEmbedImages_View
	EmbedVideo_View           *AppBskyEmbedVideo_View
	EmbedExternal_View        *AppBskyEmbedExternal_View
	EmbedRecord_View          *AppBskyEmbedRecord_View
	EmbedRecordWithMedia_View *AppBskyEmbedRecordWithMedia_View
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_PostView_Embed) UnmarshalJSON(data []byte) error {
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
	case "app.bsky.embed.record#view":
		m.EmbedRecord_View = &AppBskyEmbedRecord_View{}
		json.Unmarshal(data, m.EmbedRecord_View)
	case "app.bsky.embed.recordWithMedia#view":
		m.EmbedRecordWithMedia_View = &AppBskyEmbedRecordWithMedia_View{}
		json.Unmarshal(data, m.EmbedRecordWithMedia_View)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_PostView_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedImages_View != nil {
		return json.Marshal(m.EmbedImages_View)
	} else if m.EmbedVideo_View != nil {
		return json.Marshal(m.EmbedVideo_View)
	} else if m.EmbedExternal_View != nil {
		return json.Marshal(m.EmbedExternal_View)
	} else if m.EmbedRecord_View != nil {
		return json.Marshal(m.EmbedRecord_View)
	} else if m.EmbedRecordWithMedia_View != nil {
		return json.Marshal(m.EmbedRecordWithMedia_View)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyFeedDefs_ReasonPin_Description = ""

type AppBskyFeedDefs_ReasonPin struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedDefs_ReasonRepost_Description = ""

type AppBskyFeedDefs_ReasonRepost struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	By            *AppBskyActorDefs_ProfileViewBasic `json:"by,omitempty"`
	Cid           *string                            `json:"cid,omitempty"`
	IndexedAt     string                             `json:"indexedAt"`
	Uri           *string                            `json:"uri,omitempty"`
}

const AppBskyFeedDefs_ReplyRef_Description = ""

type AppBskyFeedDefs_ReplyRef struct {
	LexiconTypeID     string                                          `json:"$type,omitempty"`
	GrandparentAuthor *AppBskyActorDefs_ProfileViewBasic              `json:"grandparentAuthor,omitempty"`
	Parent            *AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Parent `json:"parent,omitempty"`
	Root              *AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Root   `json:"root,omitempty"`
}

type AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Parent struct {
	FeedDefs_PostView     *AppBskyFeedDefs_PostView
	FeedDefs_NotFoundPost *AppBskyFeedDefs_NotFoundPost
	FeedDefs_BlockedPost  *AppBskyFeedDefs_BlockedPost
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Parent) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#postView":
		m.FeedDefs_PostView = &AppBskyFeedDefs_PostView{}
		json.Unmarshal(data, m.FeedDefs_PostView)
	case "#notFoundPost":
		m.FeedDefs_NotFoundPost = &AppBskyFeedDefs_NotFoundPost{}
		json.Unmarshal(data, m.FeedDefs_NotFoundPost)
	case "#blockedPost":
		m.FeedDefs_BlockedPost = &AppBskyFeedDefs_BlockedPost{}
		json.Unmarshal(data, m.FeedDefs_BlockedPost)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Parent) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_PostView != nil {
		return json.Marshal(m.FeedDefs_PostView)
	} else if m.FeedDefs_NotFoundPost != nil {
		return json.Marshal(m.FeedDefs_NotFoundPost)
	} else if m.FeedDefs_BlockedPost != nil {
		return json.Marshal(m.FeedDefs_BlockedPost)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Root struct {
	FeedDefs_PostView     *AppBskyFeedDefs_PostView
	FeedDefs_NotFoundPost *AppBskyFeedDefs_NotFoundPost
	FeedDefs_BlockedPost  *AppBskyFeedDefs_BlockedPost
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Root) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#postView":
		m.FeedDefs_PostView = &AppBskyFeedDefs_PostView{}
		json.Unmarshal(data, m.FeedDefs_PostView)
	case "#notFoundPost":
		m.FeedDefs_NotFoundPost = &AppBskyFeedDefs_NotFoundPost{}
		json.Unmarshal(data, m.FeedDefs_NotFoundPost)
	case "#blockedPost":
		m.FeedDefs_BlockedPost = &AppBskyFeedDefs_BlockedPost{}
		json.Unmarshal(data, m.FeedDefs_BlockedPost)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_ReplyRef_Root) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_PostView != nil {
		return json.Marshal(m.FeedDefs_PostView)
	} else if m.FeedDefs_NotFoundPost != nil {
		return json.Marshal(m.FeedDefs_NotFoundPost)
	} else if m.FeedDefs_BlockedPost != nil {
		return json.Marshal(m.FeedDefs_BlockedPost)
	} else {
		return []byte("{}"), nil
	}
}

// Request that less content like the given feed item be shown in the feed
const AppBskyFeedDefs_RequestLess string = "requestLess"

// Request that more content like the given feed item be shown in the feed
const AppBskyFeedDefs_RequestMore string = "requestMore"

const AppBskyFeedDefs_SkeletonFeedPost_Description = ""

type AppBskyFeedDefs_SkeletonFeedPost struct {
	LexiconTypeID string                                                  `json:"$type,omitempty"`
	FeedContext   *string                                                 `json:"feedContext,omitempty"`
	Post          string                                                  `json:"post"`
	Reason        *AppBskyFeedDefsAppBskyFeedDefs_SkeletonFeedPost_Reason `json:"reason,omitempty"`
}

type AppBskyFeedDefsAppBskyFeedDefs_SkeletonFeedPost_Reason struct {
	FeedDefs_SkeletonReasonRepost *AppBskyFeedDefs_SkeletonReasonRepost
	FeedDefs_SkeletonReasonPin    *AppBskyFeedDefs_SkeletonReasonPin
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_SkeletonFeedPost_Reason) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#skeletonReasonRepost":
		m.FeedDefs_SkeletonReasonRepost = &AppBskyFeedDefs_SkeletonReasonRepost{}
		json.Unmarshal(data, m.FeedDefs_SkeletonReasonRepost)
	case "#skeletonReasonPin":
		m.FeedDefs_SkeletonReasonPin = &AppBskyFeedDefs_SkeletonReasonPin{}
		json.Unmarshal(data, m.FeedDefs_SkeletonReasonPin)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_SkeletonFeedPost_Reason) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_SkeletonReasonRepost != nil {
		return json.Marshal(m.FeedDefs_SkeletonReasonRepost)
	} else if m.FeedDefs_SkeletonReasonPin != nil {
		return json.Marshal(m.FeedDefs_SkeletonReasonPin)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyFeedDefs_SkeletonReasonPin_Description = ""

type AppBskyFeedDefs_SkeletonReasonPin struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedDefs_SkeletonReasonRepost_Description = ""

type AppBskyFeedDefs_SkeletonReasonRepost struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Repost        string `json:"repost"`
}

const AppBskyFeedDefs_ThreadContext_Description = "Metadata about this post within the context of the thread it is in."

type AppBskyFeedDefs_ThreadContext struct {
	LexiconTypeID  string  `json:"$type,omitempty"`
	RootAuthorLike *string `json:"rootAuthorLike,omitempty"`
}

const AppBskyFeedDefs_ThreadViewPost_Description = ""

type AppBskyFeedDefs_ThreadViewPost struct {
	LexiconTypeID string                                                        `json:"$type,omitempty"`
	Parent        *AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Parent         `json:"parent,omitempty"`
	Post          *AppBskyFeedDefs_PostView                                     `json:"post,omitempty"`
	Replies       []*AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Replies_Elem `json:"replies,omitempty"`
	ThreadContext *AppBskyFeedDefs_ThreadContext                                `json:"threadContext,omitempty"`
}

type AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Parent struct {
	FeedDefs_ThreadViewPost *AppBskyFeedDefs_ThreadViewPost
	FeedDefs_NotFoundPost   *AppBskyFeedDefs_NotFoundPost
	FeedDefs_BlockedPost    *AppBskyFeedDefs_BlockedPost
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Parent) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "#threadViewPost":
		m.FeedDefs_ThreadViewPost = &AppBskyFeedDefs_ThreadViewPost{}
		json.Unmarshal(data, m.FeedDefs_ThreadViewPost)
	case "#notFoundPost":
		m.FeedDefs_NotFoundPost = &AppBskyFeedDefs_NotFoundPost{}
		json.Unmarshal(data, m.FeedDefs_NotFoundPost)
	case "#blockedPost":
		m.FeedDefs_BlockedPost = &AppBskyFeedDefs_BlockedPost{}
		json.Unmarshal(data, m.FeedDefs_BlockedPost)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Parent) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_ThreadViewPost != nil {
		return json.Marshal(m.FeedDefs_ThreadViewPost)
	} else if m.FeedDefs_NotFoundPost != nil {
		return json.Marshal(m.FeedDefs_NotFoundPost)
	} else if m.FeedDefs_BlockedPost != nil {
		return json.Marshal(m.FeedDefs_BlockedPost)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Replies_Elem struct {
	FeedDefs_ThreadViewPost *AppBskyFeedDefs_ThreadViewPost
	FeedDefs_NotFoundPost   *AppBskyFeedDefs_NotFoundPost
	FeedDefs_BlockedPost    *AppBskyFeedDefs_BlockedPost
}

func (m *AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Replies_Elem) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.feed.defs#threadViewPost":
		m.FeedDefs_ThreadViewPost = &AppBskyFeedDefs_ThreadViewPost{}
		json.Unmarshal(data, m.FeedDefs_ThreadViewPost)
	case "app.bsky.feed.defs#notFoundPost":
		m.FeedDefs_NotFoundPost = &AppBskyFeedDefs_NotFoundPost{}
		json.Unmarshal(data, m.FeedDefs_NotFoundPost)
	case "app.bsky.feed.defs#blockedPost":
		m.FeedDefs_BlockedPost = &AppBskyFeedDefs_BlockedPost{}
		json.Unmarshal(data, m.FeedDefs_BlockedPost)
	}
	return nil
}

func (m AppBskyFeedDefsAppBskyFeedDefs_ThreadViewPost_Replies_Elem) MarshalJSON() ([]byte, error) {
	if m.FeedDefs_ThreadViewPost != nil {
		return json.Marshal(m.FeedDefs_ThreadViewPost)
	} else if m.FeedDefs_NotFoundPost != nil {
		return json.Marshal(m.FeedDefs_NotFoundPost)
	} else if m.FeedDefs_BlockedPost != nil {
		return json.Marshal(m.FeedDefs_BlockedPost)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyFeedDefs_ThreadgateView_Description = ""

type AppBskyFeedDefs_ThreadgateView struct {
	LexiconTypeID string                            `json:"$type,omitempty"`
	Cid           *string                           `json:"cid,omitempty"`
	Lists         []*AppBskyGraphDefs_ListViewBasic `json:"lists,omitempty"`
	Record        *any                              `json:"record,omitempty"`
	Uri           *string                           `json:"uri,omitempty"`
}

const AppBskyFeedDefs_ViewerState_Description = "Metadata about the requesting account's relationship with the subject content. Only has meaningful content for authed requests."

type AppBskyFeedDefs_ViewerState struct {
	LexiconTypeID     string  `json:"$type,omitempty"`
	Bookmarked        *bool   `json:"bookmarked,omitempty"`
	EmbeddingDisabled *bool   `json:"embeddingDisabled,omitempty"`
	Like              *string `json:"like,omitempty"`
	Pinned            *bool   `json:"pinned,omitempty"`
	ReplyDisabled     *bool   `json:"replyDisabled,omitempty"`
	Repost            *string `json:"repost,omitempty"`
	ThreadMuted       *bool   `json:"threadMuted,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.feed.defs",
  "description": "",
  "defs": {
    "blockedAuthor": {
      "type": "object",
      "description": "",
      "required": [
        "did"
      ],
      "properties": {
        "did": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#viewerState"
        }
      }
    },
    "blockedPost": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "blocked",
        "author"
      ],
      "properties": {
        "author": {
          "type": "ref",
          "ref": "#blockedAuthor"
        },
        "blocked": {
          "type": "boolean"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "clickthroughAuthor": {
      "type": "token",
      "description": "User clicked through to the author of the feed item"
    },
    "clickthroughEmbed": {
      "type": "token",
      "description": "User clicked through to the embedded content of the feed item"
    },
    "clickthroughItem": {
      "type": "token",
      "description": "User clicked through to the feed item"
    },
    "clickthroughReposter": {
      "type": "token",
      "description": "User clicked through to the reposter of the feed item"
    },
    "contentModeUnspecified": {
      "type": "token",
      "description": "Declares the feed generator returns any types of posts."
    },
    "contentModeVideo": {
      "type": "token",
      "description": "Declares the feed generator returns posts containing app.bsky.embed.video embeds."
    },
    "feedViewPost": {
      "type": "object",
      "description": "",
      "required": [
        "post"
      ],
      "properties": {
        "feedContext": {
          "type": "string"
        },
        "post": {
          "type": "ref",
          "ref": "#postView"
        },
        "reason": {
          "type": "union",
          "refs": [
            "#reasonRepost",
            "#reasonPin"
          ]
        },
        "reply": {
          "type": "ref",
          "ref": "#replyRef"
        },
        "reqId": {
          "type": "string"
        }
      }
    },
    "generatorView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "did",
        "creator",
        "displayName",
        "indexedAt"
      ],
      "properties": {
        "acceptsInteractions": {
          "type": "boolean"
        },
        "avatar": {
          "type": "string"
        },
        "cid": {
          "type": "string"
        },
        "contentMode": {
          "type": "string"
        },
        "creator": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileView"
        },
        "description": {
          "type": "string"
        },
        "descriptionFacets": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.richtext.facet"
          }
        },
        "did": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "indexedAt": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "likeCount": {
          "type": "integer"
        },
        "uri": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "#generatorViewerState"
        }
      }
    },
    "generatorViewerState": {
      "type": "object",
      "description": "",
      "properties": {
        "like": {
          "type": "string"
        }
      }
    },
    "interaction": {
      "type": "object",
      "description": "",
      "properties": {
        "event": {
          "type": "string"
        },
        "feedContext": {
          "type": "string"
        },
        "item": {
          "type": "string"
        },
        "reqId": {
          "type": "string"
        }
      }
    },
    "interactionLike": {
      "type": "token",
      "description": "User liked the feed item"
    },
    "interactionQuote": {
      "type": "token",
      "description": "User quoted the feed item"
    },
    "interactionReply": {
      "type": "token",
      "description": "User replied to the feed item"
    },
    "interactionRepost": {
      "type": "token",
      "description": "User reposted the feed item"
    },
    "interactionSeen": {
      "type": "token",
      "description": "Feed item was seen by user"
    },
    "interactionShare": {
      "type": "token",
      "description": "User shared the feed item"
    },
    "notFoundPost": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "notFound"
      ],
      "properties": {
        "notFound": {
          "type": "boolean"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "postView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "author",
        "record",
        "indexedAt"
      ],
      "properties": {
        "author": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "bookmarkCount": {
          "type": "integer"
        },
        "cid": {
          "type": "string"
        },
        "debug": {
          "type": "unknown"
        },
        "embed": {
          "type": "union",
          "refs": [
            "app.bsky.embed.images#view",
            "app.bsky.embed.video#view",
            "app.bsky.embed.external#view",
            "app.bsky.embed.record#view",
            "app.bsky.embed.recordWithMedia#view"
          ]
        },
        "indexedAt": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "likeCount": {
          "type": "integer"
        },
        "quoteCount": {
          "type": "integer"
        },
        "record": {
          "type": "unknown"
        },
        "replyCount": {
          "type": "integer"
        },
        "repostCount": {
          "type": "integer"
        },
        "threadgate": {
          "type": "ref",
          "ref": "#threadgateView"
        },
        "uri": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "#viewerState"
        }
      }
    },
    "reasonPin": {
      "type": "object",
      "description": ""
    },
    "reasonRepost": {
      "type": "object",
      "description": "",
      "required": [
        "by",
        "indexedAt"
      ],
      "properties": {
        "by": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "cid": {
          "type": "string"
        },
        "indexedAt": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "replyRef": {
      "type": "object",
      "description": "",
      "required": [
        "root",
        "parent"
      ],
      "properties": {
        "grandparentAuthor": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "parent": {
          "type": "union",
          "refs": [
            "#postView",
            "#notFoundPost",
            "#blockedPost"
          ]
        },
        "root": {
          "type": "union",
          "refs": [
            "#postView",
            "#notFoundPost",
            "#blockedPost"
          ]
        }
      }
    },
    "requestLess": {
      "type": "token",
      "description": "Request that less content like the given feed item be shown in the feed"
    },
    "requestMore": {
      "type": "token",
      "description": "Request that more content like the given feed item be shown in the feed"
    },
    "skeletonFeedPost": {
      "type": "object",
      "description": "",
      "required": [
        "post"
      ],
      "properties": {
        "feedContext": {
          "type": "string"
        },
        "post": {
          "type": "string"
        },
        "reason": {
          "type": "union",
          "refs": [
            "#skeletonReasonRepost",
            "#skeletonReasonPin"
          ]
        }
      }
    },
    "skeletonReasonPin": {
      "type": "object",
      "description": ""
    },
    "skeletonReasonRepost": {
      "type": "object",
      "description": "",
      "required": [
        "repost"
      ],
      "properties": {
        "repost": {
          "type": "string"
        }
      }
    },
    "threadContext": {
      "type": "object",
      "description": "Metadata about this post within the context of the thread it is in.",
      "properties": {
        "rootAuthorLike": {
          "type": "string"
        }
      }
    },
    "threadViewPost": {
      "type": "object",
      "description": "",
      "required": [
        "post"
      ],
      "properties": {
        "parent": {
          "type": "union",
          "refs": [
            "#threadViewPost",
            "#notFoundPost",
            "#blockedPost"
          ]
        },
        "post": {
          "type": "ref",
          "ref": "#postView"
        },
        "replies": {
          "type": "array",
          "items": {
            "type": "union",
            "refs": [
              "#threadViewPost",
              "#notFoundPost",
              "#blockedPost"
            ]
          }
        },
        "threadContext": {
          "type": "ref",
          "ref": "#threadContext"
        }
      }
    },
    "threadgateView": {
      "type": "object",
      "description": "",
      "properties": {
        "cid": {
          "type": "string"
        },
        "lists": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.graph.defs#listViewBasic"
          }
        },
        "record": {
          "type": "unknown"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "viewerState": {
      "type": "object",
      "description": "Metadata about the requesting account's relationship with the subject content. Only has meaningful content for authed requests.",
      "properties": {
        "bookmarked": {
          "type": "boolean"
        },
        "embeddingDisabled": {
          "type": "boolean"
        },
        "like": {
          "type": "string"
        },
        "pinned": {
          "type": "boolean"
        },
        "replyDisabled": {
          "type": "boolean"
        },
        "repost": {
          "type": "string"
        },
        "threadMuted": {
          "type": "boolean"
        }
      }
    }
  }
}
*/
