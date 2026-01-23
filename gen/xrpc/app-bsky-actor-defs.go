// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.actor.defs

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyActorDefs_KnownFollowers_Description = "The subject's followers whom you also follow"

type AppBskyActorDefs_KnownFollowers struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Count         int64                                `json:"count"`
	Followers     []*AppBskyActorDefs_ProfileViewBasic `json:"followers,omitempty"`
}

const AppBskyActorDefs_ProfileAssociated_Description = ""

type AppBskyActorDefs_ProfileAssociated struct {
	LexiconTypeID        string                                                  `json:"$type,omitempty"`
	ActivitySubscription *AppBskyActorDefs_ProfileAssociatedActivitySubscription `json:"activitySubscription,omitempty"`
	Chat                 *AppBskyActorDefs_ProfileAssociatedChat                 `json:"chat,omitempty"`
	Feedgens             *int64                                                  `json:"feedgens,omitempty"`
	Germ                 *AppBskyActorDefs_ProfileAssociatedGerm                 `json:"germ,omitempty"`
	Labeler              *bool                                                   `json:"labeler,omitempty"`
	Lists                *int64                                                  `json:"lists,omitempty"`
	StarterPacks         *int64                                                  `json:"starterPacks,omitempty"`
}

const AppBskyActorDefs_ProfileAssociatedActivitySubscription_Description = ""

type AppBskyActorDefs_ProfileAssociatedActivitySubscription struct {
	LexiconTypeID      string `json:"$type,omitempty"`
	AllowSubscriptions string `json:"allowSubscriptions"`
}

const AppBskyActorDefs_ProfileAssociatedChat_Description = ""

type AppBskyActorDefs_ProfileAssociatedChat struct {
	LexiconTypeID string `json:"$type,omitempty"`
	AllowIncoming string `json:"allowIncoming"`
}

const AppBskyActorDefs_ProfileAssociatedGerm_Description = ""

type AppBskyActorDefs_ProfileAssociatedGerm struct {
	LexiconTypeID string `json:"$type,omitempty"`
	MessageMeUrl  string `json:"messageMeUrl"`
	ShowButtonTo  string `json:"showButtonTo"`
}

const AppBskyActorDefs_ProfileView_Description = ""

type AppBskyActorDefs_ProfileView struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	CreatedAt     *string                             `json:"createdAt,omitempty"`
	Debug         *any                                `json:"debug,omitempty"`
	Description   *string                             `json:"description,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	IndexedAt     *string                             `json:"indexedAt,omitempty"`
	Labels        []*ComATProtoLabelDefs_Label        `json:"labels,omitempty"`
	Pronouns      *string                             `json:"pronouns,omitempty"`
	Status        *AppBskyActorDefs_StatusView        `json:"status,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

const AppBskyActorDefs_ProfileViewBasic_Description = ""

type AppBskyActorDefs_ProfileViewBasic struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	CreatedAt     *string                             `json:"createdAt,omitempty"`
	Debug         *any                                `json:"debug,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	Labels        []*ComATProtoLabelDefs_Label        `json:"labels,omitempty"`
	Pronouns      *string                             `json:"pronouns,omitempty"`
	Status        *AppBskyActorDefs_StatusView        `json:"status,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

const AppBskyActorDefs_StatusView_Description = ""

type AppBskyActorDefs_StatusView struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	Cid           *string                            `json:"cid,omitempty"`
	Embed         *AppBskyActorDefs_StatusView_Embed `json:"embed,omitempty"`
	ExpiresAt     *string                            `json:"expiresAt,omitempty"`
	IsActive      *bool                              `json:"isActive,omitempty"`
	IsDisabled    *bool                              `json:"isDisabled,omitempty"`
	Record        any                                `json:"record"`
	Status        string                             `json:"status"`
	Uri           *string                            `json:"uri,omitempty"`
}

type AppBskyActorDefs_StatusView_Embed struct {
	EmbedExternal_View *AppBskyEmbedExternal_View
}

func (m *AppBskyActorDefs_StatusView_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.external#view":
		m.EmbedExternal_View = &AppBskyEmbedExternal_View{}
		json.Unmarshal(data, m.EmbedExternal_View)
	}
	return nil
}

func (m AppBskyActorDefs_StatusView_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedExternal_View != nil {
		return json.Marshal(m.EmbedExternal_View)
	} else {
		return []byte("{}"), nil
	}
}

const AppBskyActorDefs_VerificationState_Description = "Represents the verification information about the user this object is attached to."

type AppBskyActorDefs_VerificationState struct {
	LexiconTypeID         string                               `json:"$type,omitempty"`
	TrustedVerifierStatus string                               `json:"trustedVerifierStatus"`
	Verifications         []*AppBskyActorDefs_VerificationView `json:"verifications,omitempty"`
	VerifiedStatus        string                               `json:"verifiedStatus"`
}

const AppBskyActorDefs_VerificationView_Description = "An individual verification for an associated subject."

type AppBskyActorDefs_VerificationView struct {
	LexiconTypeID string `json:"$type,omitempty"`
	CreatedAt     string `json:"createdAt"`
	IsValid       bool   `json:"isValid"`
	Issuer        string `json:"issuer"`
	Uri           string `json:"uri"`
}

const AppBskyActorDefs_ViewerState_Description = "Metadata about the requesting account's relationship with the subject account. Only has meaningful content for authed requests."

type AppBskyActorDefs_ViewerState struct {
	LexiconTypeID        string                                        `json:"$type,omitempty"`
	ActivitySubscription *AppBskyNotificationDefs_ActivitySubscription `json:"activitySubscription,omitempty"`
	BlockedBy            *bool                                         `json:"blockedBy,omitempty"`
	Blocking             *string                                       `json:"blocking,omitempty"`
	BlockingByList       *AppBskyGraphDefs_ListViewBasic               `json:"blockingByList,omitempty"`
	FollowedBy           *string                                       `json:"followedBy,omitempty"`
	Following            *string                                       `json:"following,omitempty"`
	KnownFollowers       *AppBskyActorDefs_KnownFollowers              `json:"knownFollowers,omitempty"`
	Muted                *bool                                         `json:"muted,omitempty"`
	MutedByList          *AppBskyGraphDefs_ListViewBasic               `json:"mutedByList,omitempty"`
}
