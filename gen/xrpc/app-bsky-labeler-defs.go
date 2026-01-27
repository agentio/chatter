// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.labeler.defs

const AppBskyLabelerDefs_LabelerView_Description = ""

// AppBskyLabelerDefs_LabelerView is a record with Lexicon type app.bsky.labeler.defs#labelerView
type AppBskyLabelerDefs_LabelerView struct {
	LexiconTypeID string                                 `json:"$type,omitempty"`
	Cid           string                                 `json:"cid"`
	Creator       *AppBskyActorDefs_ProfileView          `json:"creator,omitempty"`
	IndexedAt     string                                 `json:"indexedAt"`
	Labels        []*ComATProtoLabelDefs_Label           `json:"labels,omitempty"`
	LikeCount     *int64                                 `json:"likeCount,omitempty"`
	Uri           string                                 `json:"uri"`
	Viewer        *AppBskyLabelerDefs_LabelerViewerState `json:"viewer,omitempty"`
}

const AppBskyLabelerDefs_LabelerViewerState_Description = ""

// AppBskyLabelerDefs_LabelerViewerState is a record with Lexicon type app.bsky.labeler.defs#labelerViewerState
type AppBskyLabelerDefs_LabelerViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Like          *string `json:"like,omitempty"`
}
