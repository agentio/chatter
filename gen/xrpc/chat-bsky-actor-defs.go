// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.actor.defs

const ChatBskyActorDefs_ProfileViewBasic_Description = ""

type ChatBskyActorDefs_ProfileViewBasic struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	ChatDisabled  *bool                               `json:"chatDisabled,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	Labels        []*ComATProtoLabelDefs_Label        `json:"labels,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}
