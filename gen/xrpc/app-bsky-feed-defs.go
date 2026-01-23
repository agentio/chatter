// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.defs

const AppBskyFeedDefs_BlockedAuthor_Description = ""

type AppBskyFeedDefs_BlockedAuthor struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Did           string                        `json:"did"`
	Viewer        *AppBskyActorDefs_ViewerState `json:"viewer,omitempty"`
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
	Labels              []*ComATProtoLabelDefs_Label          `json:"labels,omitempty"`
	LikeCount           *int64                                `json:"likeCount,omitempty"`
	Uri                 string                                `json:"uri"`
	Viewer              *AppBskyFeedDefs_GeneratorViewerState `json:"viewer,omitempty"`
}

const AppBskyFeedDefs_GeneratorViewerState_Description = ""

type AppBskyFeedDefs_GeneratorViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Like          *string `json:"like,omitempty"`
}
