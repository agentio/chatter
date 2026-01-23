// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.graph.defs

type AppBskyGraphDefs_ListPurpose string

const AppBskyGraphDefs_ListView_Description = ""

type AppBskyGraphDefs_ListView struct {
	LexiconTypeID     string                            `json:"$type,omitempty"`
	Avatar            *string                           `json:"avatar,omitempty"`
	Cid               string                            `json:"cid"`
	Creator           *AppBskyActorDefs_ProfileView     `json:"creator,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	DescriptionFacets []*AppBskyRichtextFacet           `json:"descriptionFacets,omitempty"`
	IndexedAt         string                            `json:"indexedAt"`
	Labels            []*ComATProtoLabelDefs_Label      `json:"labels,omitempty"`
	ListItemCount     *int64                            `json:"listItemCount,omitempty"`
	Name              string                            `json:"name"`
	Purpose           *AppBskyGraphDefs_ListPurpose     `json:"purpose,omitempty"`
	Uri               string                            `json:"uri"`
	Viewer            *AppBskyGraphDefs_ListViewerState `json:"viewer,omitempty"`
}

const AppBskyGraphDefs_ListViewBasic_Description = ""

type AppBskyGraphDefs_ListViewBasic struct {
	LexiconTypeID string                            `json:"$type,omitempty"`
	Avatar        *string                           `json:"avatar,omitempty"`
	Cid           string                            `json:"cid"`
	IndexedAt     *string                           `json:"indexedAt,omitempty"`
	Labels        []*ComATProtoLabelDefs_Label      `json:"labels,omitempty"`
	ListItemCount *int64                            `json:"listItemCount,omitempty"`
	Name          string                            `json:"name"`
	Purpose       *AppBskyGraphDefs_ListPurpose     `json:"purpose,omitempty"`
	Uri           string                            `json:"uri"`
	Viewer        *AppBskyGraphDefs_ListViewerState `json:"viewer,omitempty"`
}

const AppBskyGraphDefs_ListViewerState_Description = ""

type AppBskyGraphDefs_ListViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Blocked       *string `json:"blocked,omitempty"`
	Muted         *bool   `json:"muted,omitempty"`
}

const AppBskyGraphDefs_StarterPackViewBasic_Description = ""

type AppBskyGraphDefs_StarterPackViewBasic struct {
	LexiconTypeID      string                             `json:"$type,omitempty"`
	Cid                string                             `json:"cid"`
	Creator            *AppBskyActorDefs_ProfileViewBasic `json:"creator,omitempty"`
	IndexedAt          string                             `json:"indexedAt"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty"`
	Labels             []*ComATProtoLabelDefs_Label       `json:"labels,omitempty"`
	ListItemCount      *int64                             `json:"listItemCount,omitempty"`
	Record             any                                `json:"record"`
	Uri                string                             `json:"uri"`
}
