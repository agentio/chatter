// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.actor.defs

import (
	"encoding/json"

	"github.com/agentio/slink/pkg/slink"
)

type AppBskyActorDefs_AdultContentPref struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Enabled       bool   `json:"enabled"`
}

// If set, an active progress guide. Once completed, can be set to undefined. Should have unspecced fields tracking progress.
type AppBskyActorDefs_BskyAppProgressGuide struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Guide         string `json:"guide"`
}

// A grab bag of state that's specific to the bsky.app program. Third-party apps shouldn't use this.
type AppBskyActorDefs_BskyAppStatePref struct {
	LexiconTypeID       string                                 `json:"$type,omitempty"`
	ActiveProgressGuide *AppBskyActorDefs_BskyAppProgressGuide `json:"activeProgressGuide,omitempty"`
	Nuxs                []*AppBskyActorDefs_Nux                `json:"nuxs,omitempty"`
	QueuedNudges        []string                               `json:"queuedNudges,omitempty"`
}

type AppBskyActorDefs_ContentLabelPref struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Label         string  `json:"label"`
	LabelerDid    *string `json:"labelerDid,omitempty"`
	Visibility    string  `json:"visibility"`
}

// Read-only preference containing value(s) inferred from the user's declared birthdate. Absence of this preference object in the response indicates that the user has not made a declaration.
type AppBskyActorDefs_DeclaredAgePref struct {
	LexiconTypeID string `json:"$type,omitempty"`
	IsOverAge13   *bool  `json:"isOverAge13,omitempty"`
	IsOverAge16   *bool  `json:"isOverAge16,omitempty"`
	IsOverAge18   *bool  `json:"isOverAge18,omitempty"`
}

type AppBskyActorDefs_FeedViewPref struct {
	LexiconTypeID           string `json:"$type,omitempty"`
	Feed                    string `json:"feed"`
	HideQuotePosts          *bool  `json:"hideQuotePosts,omitempty"`
	HideReplies             *bool  `json:"hideReplies,omitempty"`
	HideRepliesByLikeCount  *int64 `json:"hideRepliesByLikeCount,omitempty"`
	HideRepliesByUnfollowed *bool  `json:"hideRepliesByUnfollowed,omitempty"`
	HideReposts             *bool  `json:"hideReposts,omitempty"`
}

type AppBskyActorDefs_HiddenPostsPref struct {
	LexiconTypeID string   `json:"$type,omitempty"`
	Items         []string `json:"items,omitempty"`
}

type AppBskyActorDefs_InterestsPref struct {
	LexiconTypeID string   `json:"$type,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

// The subject's followers whom you also follow
type AppBskyActorDefs_KnownFollowers struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Count         int64                                `json:"count"`
	Followers     []*AppBskyActorDefs_ProfileViewBasic `json:"followers,omitempty"`
}

type AppBskyActorDefs_LabelerPrefItem struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}

type AppBskyActorDefs_LabelersPref struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Labelers      []*AppBskyActorDefs_LabelerPrefItem `json:"labelers,omitempty"`
}

// A word that the account owner has muted.
type AppBskyActorDefs_MutedWord struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	ActorTarget   *string                             `json:"actorTarget,omitempty"`
	ExpiresAt     *string                             `json:"expiresAt,omitempty"`
	Id            *string                             `json:"id,omitempty"`
	Targets       []*AppBskyActorDefs_MutedWordTarget `json:"targets,omitempty"`
	Value         string                              `json:"value"`
}

type AppBskyActorDefs_MutedWordTarget string

type AppBskyActorDefs_MutedWordsPref struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Items         []*AppBskyActorDefs_MutedWord `json:"items,omitempty"`
}

// A new user experiences (NUX) storage object
type AppBskyActorDefs_Nux struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Completed     bool    `json:"completed"`
	Data          *string `json:"data,omitempty"`
	ExpiresAt     *string `json:"expiresAt,omitempty"`
	Id            string  `json:"id"`
}

type AppBskyActorDefs_PersonalDetailsPref struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	BirthDate     *string `json:"birthDate,omitempty"`
}

// Default post interaction settings for the account. These values should be applied as default values when creating new posts. These refs should mirror the threadgate and postgate records exactly.
type AppBskyActorDefs_PostInteractionSettingsPref struct {
	LexiconTypeID          string                                                                                      `json:"$type,omitempty"`
	PostgateEmbeddingRules []*AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_PostgateEmbeddingRules_Elem `json:"postgateEmbeddingRules,omitempty"`
	ThreadgateAllowRules   []*AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_ThreadgateAllowRules_Elem   `json:"threadgateAllowRules,omitempty"`
}

type AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_PostgateEmbeddingRules_Elem struct {
	FeedPostgate_DisableRule *AppBskyFeedPostgate_DisableRule
}

func (m *AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_PostgateEmbeddingRules_Elem) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.feed.postgate#disableRule":
		m.FeedPostgate_DisableRule = &AppBskyFeedPostgate_DisableRule{}
		json.Unmarshal(data, m.FeedPostgate_DisableRule)
	}
	return nil
}

func (m AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_PostgateEmbeddingRules_Elem) MarshalJSON() ([]byte, error) {
	if m.FeedPostgate_DisableRule != nil {
		return json.Marshal(m.FeedPostgate_DisableRule)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_ThreadgateAllowRules_Elem struct {
	FeedThreadgate_MentionRule   *AppBskyFeedThreadgate_MentionRule
	FeedThreadgate_FollowerRule  *AppBskyFeedThreadgate_FollowerRule
	FeedThreadgate_FollowingRule *AppBskyFeedThreadgate_FollowingRule
	FeedThreadgate_ListRule      *AppBskyFeedThreadgate_ListRule
}

func (m *AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_ThreadgateAllowRules_Elem) UnmarshalJSON(data []byte) error {
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

func (m AppBskyActorDefsAppBskyActorDefs_PostInteractionSettingsPref_ThreadgateAllowRules_Elem) MarshalJSON() ([]byte, error) {
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

type AppBskyActorDefs_Preferences_Elem struct {
	ActorDefs_AdultContentPref            *AppBskyActorDefs_AdultContentPref
	ActorDefs_ContentLabelPref            *AppBskyActorDefs_ContentLabelPref
	ActorDefs_SavedFeedsPref              *AppBskyActorDefs_SavedFeedsPref
	ActorDefs_SavedFeedsPrefV2            *AppBskyActorDefs_SavedFeedsPrefV2
	ActorDefs_PersonalDetailsPref         *AppBskyActorDefs_PersonalDetailsPref
	ActorDefs_DeclaredAgePref             *AppBskyActorDefs_DeclaredAgePref
	ActorDefs_FeedViewPref                *AppBskyActorDefs_FeedViewPref
	ActorDefs_ThreadViewPref              *AppBskyActorDefs_ThreadViewPref
	ActorDefs_InterestsPref               *AppBskyActorDefs_InterestsPref
	ActorDefs_MutedWordsPref              *AppBskyActorDefs_MutedWordsPref
	ActorDefs_HiddenPostsPref             *AppBskyActorDefs_HiddenPostsPref
	ActorDefs_BskyAppStatePref            *AppBskyActorDefs_BskyAppStatePref
	ActorDefs_LabelersPref                *AppBskyActorDefs_LabelersPref
	ActorDefs_PostInteractionSettingsPref *AppBskyActorDefs_PostInteractionSettingsPref
	ActorDefs_VerificationPrefs           *AppBskyActorDefs_VerificationPrefs
}

func (m *AppBskyActorDefs_Preferences_Elem) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.actor.defs#adultContentPref":
		m.ActorDefs_AdultContentPref = &AppBskyActorDefs_AdultContentPref{}
		json.Unmarshal(data, m.ActorDefs_AdultContentPref)
	case "app.bsky.actor.defs#contentLabelPref":
		m.ActorDefs_ContentLabelPref = &AppBskyActorDefs_ContentLabelPref{}
		json.Unmarshal(data, m.ActorDefs_ContentLabelPref)
	case "app.bsky.actor.defs#savedFeedsPref":
		m.ActorDefs_SavedFeedsPref = &AppBskyActorDefs_SavedFeedsPref{}
		json.Unmarshal(data, m.ActorDefs_SavedFeedsPref)
	case "app.bsky.actor.defs#savedFeedsPrefV2":
		m.ActorDefs_SavedFeedsPrefV2 = &AppBskyActorDefs_SavedFeedsPrefV2{}
		json.Unmarshal(data, m.ActorDefs_SavedFeedsPrefV2)
	case "app.bsky.actor.defs#personalDetailsPref":
		m.ActorDefs_PersonalDetailsPref = &AppBskyActorDefs_PersonalDetailsPref{}
		json.Unmarshal(data, m.ActorDefs_PersonalDetailsPref)
	case "app.bsky.actor.defs#declaredAgePref":
		m.ActorDefs_DeclaredAgePref = &AppBskyActorDefs_DeclaredAgePref{}
		json.Unmarshal(data, m.ActorDefs_DeclaredAgePref)
	case "app.bsky.actor.defs#feedViewPref":
		m.ActorDefs_FeedViewPref = &AppBskyActorDefs_FeedViewPref{}
		json.Unmarshal(data, m.ActorDefs_FeedViewPref)
	case "app.bsky.actor.defs#threadViewPref":
		m.ActorDefs_ThreadViewPref = &AppBskyActorDefs_ThreadViewPref{}
		json.Unmarshal(data, m.ActorDefs_ThreadViewPref)
	case "app.bsky.actor.defs#interestsPref":
		m.ActorDefs_InterestsPref = &AppBskyActorDefs_InterestsPref{}
		json.Unmarshal(data, m.ActorDefs_InterestsPref)
	case "app.bsky.actor.defs#mutedWordsPref":
		m.ActorDefs_MutedWordsPref = &AppBskyActorDefs_MutedWordsPref{}
		json.Unmarshal(data, m.ActorDefs_MutedWordsPref)
	case "app.bsky.actor.defs#hiddenPostsPref":
		m.ActorDefs_HiddenPostsPref = &AppBskyActorDefs_HiddenPostsPref{}
		json.Unmarshal(data, m.ActorDefs_HiddenPostsPref)
	case "app.bsky.actor.defs#bskyAppStatePref":
		m.ActorDefs_BskyAppStatePref = &AppBskyActorDefs_BskyAppStatePref{}
		json.Unmarshal(data, m.ActorDefs_BskyAppStatePref)
	case "app.bsky.actor.defs#labelersPref":
		m.ActorDefs_LabelersPref = &AppBskyActorDefs_LabelersPref{}
		json.Unmarshal(data, m.ActorDefs_LabelersPref)
	case "app.bsky.actor.defs#postInteractionSettingsPref":
		m.ActorDefs_PostInteractionSettingsPref = &AppBskyActorDefs_PostInteractionSettingsPref{}
		json.Unmarshal(data, m.ActorDefs_PostInteractionSettingsPref)
	case "app.bsky.actor.defs#verificationPrefs":
		m.ActorDefs_VerificationPrefs = &AppBskyActorDefs_VerificationPrefs{}
		json.Unmarshal(data, m.ActorDefs_VerificationPrefs)
	}
	return nil
}

func (m AppBskyActorDefs_Preferences_Elem) MarshalJSON() ([]byte, error) {
	if m.ActorDefs_AdultContentPref != nil {
		return json.Marshal(m.ActorDefs_AdultContentPref)
	} else if m.ActorDefs_ContentLabelPref != nil {
		return json.Marshal(m.ActorDefs_ContentLabelPref)
	} else if m.ActorDefs_SavedFeedsPref != nil {
		return json.Marshal(m.ActorDefs_SavedFeedsPref)
	} else if m.ActorDefs_SavedFeedsPrefV2 != nil {
		return json.Marshal(m.ActorDefs_SavedFeedsPrefV2)
	} else if m.ActorDefs_PersonalDetailsPref != nil {
		return json.Marshal(m.ActorDefs_PersonalDetailsPref)
	} else if m.ActorDefs_DeclaredAgePref != nil {
		return json.Marshal(m.ActorDefs_DeclaredAgePref)
	} else if m.ActorDefs_FeedViewPref != nil {
		return json.Marshal(m.ActorDefs_FeedViewPref)
	} else if m.ActorDefs_ThreadViewPref != nil {
		return json.Marshal(m.ActorDefs_ThreadViewPref)
	} else if m.ActorDefs_InterestsPref != nil {
		return json.Marshal(m.ActorDefs_InterestsPref)
	} else if m.ActorDefs_MutedWordsPref != nil {
		return json.Marshal(m.ActorDefs_MutedWordsPref)
	} else if m.ActorDefs_HiddenPostsPref != nil {
		return json.Marshal(m.ActorDefs_HiddenPostsPref)
	} else if m.ActorDefs_BskyAppStatePref != nil {
		return json.Marshal(m.ActorDefs_BskyAppStatePref)
	} else if m.ActorDefs_LabelersPref != nil {
		return json.Marshal(m.ActorDefs_LabelersPref)
	} else if m.ActorDefs_PostInteractionSettingsPref != nil {
		return json.Marshal(m.ActorDefs_PostInteractionSettingsPref)
	} else if m.ActorDefs_VerificationPrefs != nil {
		return json.Marshal(m.ActorDefs_VerificationPrefs)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyActorDefs_ProfileAssociated struct {
	LexiconTypeID        string                                                  `json:"$type,omitempty"`
	ActivitySubscription *AppBskyActorDefs_ProfileAssociatedActivitySubscription `json:"activitySubscription,omitempty"`
	Chat                 *AppBskyActorDefs_ProfileAssociatedChat                 `json:"chat,omitempty"`
	Feedgens             *int64                                                  `json:"feedgens,omitempty"`
	Labeler              *bool                                                   `json:"labeler,omitempty"`
	Lists                *int64                                                  `json:"lists,omitempty"`
	StarterPacks         *int64                                                  `json:"starterPacks,omitempty"`
}

type AppBskyActorDefs_ProfileAssociatedActivitySubscription struct {
	LexiconTypeID      string `json:"$type,omitempty"`
	AllowSubscriptions string `json:"allowSubscriptions"`
}

type AppBskyActorDefs_ProfileAssociatedChat struct {
	LexiconTypeID string `json:"$type,omitempty"`
	AllowIncoming string `json:"allowIncoming"`
}

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
	Labels        []*LabelDefs_Label                  `json:"labels,omitempty"`
	Pronouns      *string                             `json:"pronouns,omitempty"`
	Status        *AppBskyActorDefs_StatusView        `json:"status,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

type AppBskyActorDefs_ProfileViewBasic struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	CreatedAt     *string                             `json:"createdAt,omitempty"`
	Debug         *any                                `json:"debug,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	Labels        []*LabelDefs_Label                  `json:"labels,omitempty"`
	Pronouns      *string                             `json:"pronouns,omitempty"`
	Status        *AppBskyActorDefs_StatusView        `json:"status,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

type AppBskyActorDefs_ProfileViewDetailed struct {
	LexiconTypeID        string                                 `json:"$type,omitempty"`
	Associated           *AppBskyActorDefs_ProfileAssociated    `json:"associated,omitempty"`
	Avatar               *string                                `json:"avatar,omitempty"`
	Banner               *string                                `json:"banner,omitempty"`
	CreatedAt            *string                                `json:"createdAt,omitempty"`
	Debug                *any                                   `json:"debug,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	Did                  string                                 `json:"did"`
	DisplayName          *string                                `json:"displayName,omitempty"`
	FollowersCount       *int64                                 `json:"followersCount,omitempty"`
	FollowsCount         *int64                                 `json:"followsCount,omitempty"`
	Handle               string                                 `json:"handle"`
	IndexedAt            *string                                `json:"indexedAt,omitempty"`
	JoinedViaStarterPack *AppBskyGraphDefs_StarterPackViewBasic `json:"joinedViaStarterPack,omitempty"`
	Labels               []*LabelDefs_Label                     `json:"labels,omitempty"`
	PinnedPost           *RepoStrongRef                         `json:"pinnedPost,omitempty"`
	PostsCount           *int64                                 `json:"postsCount,omitempty"`
	Pronouns             *string                                `json:"pronouns,omitempty"`
	Status               *AppBskyActorDefs_StatusView           `json:"status,omitempty"`
	Verification         *AppBskyActorDefs_VerificationState    `json:"verification,omitempty"`
	Viewer               *AppBskyActorDefs_ViewerState          `json:"viewer,omitempty"`
	Website              *string                                `json:"website,omitempty"`
}

type AppBskyActorDefs_SavedFeed struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Id            string `json:"id"`
	Pinned        bool   `json:"pinned"`
	Type          string `json:"type"`
	Value         string `json:"value"`
}

type AppBskyActorDefs_SavedFeedsPref struct {
	LexiconTypeID string   `json:"$type,omitempty"`
	Pinned        []string `json:"pinned,omitempty"`
	Saved         []string `json:"saved,omitempty"`
	TimelineIndex *int64   `json:"timelineIndex,omitempty"`
}

type AppBskyActorDefs_SavedFeedsPrefV2 struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Items         []*AppBskyActorDefs_SavedFeed `json:"items,omitempty"`
}

type AppBskyActorDefs_StatusView struct {
	LexiconTypeID string                                             `json:"$type,omitempty"`
	Cid           *string                                            `json:"cid,omitempty"`
	Embed         *AppBskyActorDefsAppBskyActorDefs_StatusView_Embed `json:"embed,omitempty"`
	ExpiresAt     *string                                            `json:"expiresAt,omitempty"`
	IsActive      *bool                                              `json:"isActive,omitempty"`
	IsDisabled    *bool                                              `json:"isDisabled,omitempty"`
	Record        any                                                `json:"record"`
	Status        string                                             `json:"status"`
	Uri           *string                                            `json:"uri,omitempty"`
}

type AppBskyActorDefsAppBskyActorDefs_StatusView_Embed struct {
	EmbedExternal_View *AppBskyEmbedExternal_View
}

func (m *AppBskyActorDefsAppBskyActorDefs_StatusView_Embed) UnmarshalJSON(data []byte) error {
	recordType := slink.LexiconTypeFromJSONBytes(data)
	switch recordType {
	case "app.bsky.embed.external#view":
		m.EmbedExternal_View = &AppBskyEmbedExternal_View{}
		json.Unmarshal(data, m.EmbedExternal_View)
	}
	return nil
}

func (m AppBskyActorDefsAppBskyActorDefs_StatusView_Embed) MarshalJSON() ([]byte, error) {
	if m.EmbedExternal_View != nil {
		return json.Marshal(m.EmbedExternal_View)
	} else {
		return []byte("{}"), nil
	}
}

type AppBskyActorDefs_ThreadViewPref struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Sort          *string `json:"sort,omitempty"`
}

// Preferences for how verified accounts appear in the app.
type AppBskyActorDefs_VerificationPrefs struct {
	LexiconTypeID string `json:"$type,omitempty"`
	HideBadges    *bool  `json:"hideBadges,omitempty"`
}

// Represents the verification information about the user this object is attached to.
type AppBskyActorDefs_VerificationState struct {
	LexiconTypeID         string                               `json:"$type,omitempty"`
	TrustedVerifierStatus string                               `json:"trustedVerifierStatus"`
	Verifications         []*AppBskyActorDefs_VerificationView `json:"verifications,omitempty"`
	VerifiedStatus        string                               `json:"verifiedStatus"`
}

// An individual verification for an associated subject.
type AppBskyActorDefs_VerificationView struct {
	LexiconTypeID string `json:"$type,omitempty"`
	CreatedAt     string `json:"createdAt"`
	IsValid       bool   `json:"isValid"`
	Issuer        string `json:"issuer"`
	Uri           string `json:"uri"`
}

// Metadata about the requesting account's relationship with the subject account. Only has meaningful content for authed requests.
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

/*
{
  "lexicon": 1,
  "id": "app.bsky.actor.defs",
  "description": "",
  "defs": {
    "adultContentPref": {
      "type": "object",
      "description": "",
      "required": [
        "enabled"
      ],
      "properties": {
        "enabled": {
          "type": "boolean",
          "default": false
        }
      }
    },
    "bskyAppProgressGuide": {
      "type": "object",
      "description": "If set, an active progress guide. Once completed, can be set to undefined. Should have unspecced fields tracking progress.",
      "required": [
        "guide"
      ],
      "properties": {
        "guide": {
          "type": "string"
        }
      }
    },
    "bskyAppStatePref": {
      "type": "object",
      "description": "A grab bag of state that's specific to the bsky.app program. Third-party apps shouldn't use this.",
      "properties": {
        "activeProgressGuide": {
          "type": "ref",
          "ref": "#bskyAppProgressGuide"
        },
        "nuxs": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.actor.defs#nux"
          }
        },
        "queuedNudges": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "contentLabelPref": {
      "type": "object",
      "description": "",
      "required": [
        "label",
        "visibility"
      ],
      "properties": {
        "label": {
          "type": "string"
        },
        "labelerDid": {
          "type": "string"
        },
        "visibility": {
          "type": "string"
        }
      }
    },
    "declaredAgePref": {
      "type": "object",
      "description": "Read-only preference containing value(s) inferred from the user's declared birthdate. Absence of this preference object in the response indicates that the user has not made a declaration.",
      "properties": {
        "isOverAge13": {
          "type": "boolean"
        },
        "isOverAge16": {
          "type": "boolean"
        },
        "isOverAge18": {
          "type": "boolean"
        }
      }
    },
    "feedViewPref": {
      "type": "object",
      "description": "",
      "required": [
        "feed"
      ],
      "properties": {
        "feed": {
          "type": "string"
        },
        "hideQuotePosts": {
          "type": "boolean"
        },
        "hideReplies": {
          "type": "boolean"
        },
        "hideRepliesByLikeCount": {
          "type": "integer"
        },
        "hideRepliesByUnfollowed": {
          "type": "boolean",
          "default": true
        },
        "hideReposts": {
          "type": "boolean"
        }
      }
    },
    "hiddenPostsPref": {
      "type": "object",
      "description": "",
      "required": [
        "items"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "interestsPref": {
      "type": "object",
      "description": "",
      "required": [
        "tags"
      ],
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "knownFollowers": {
      "type": "object",
      "description": "The subject's followers whom you also follow",
      "required": [
        "count",
        "followers"
      ],
      "properties": {
        "count": {
          "type": "integer"
        },
        "followers": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#profileViewBasic"
          }
        }
      }
    },
    "labelerPrefItem": {
      "type": "object",
      "description": "",
      "required": [
        "did"
      ],
      "properties": {
        "did": {
          "type": "string"
        }
      }
    },
    "labelersPref": {
      "type": "object",
      "description": "",
      "required": [
        "labelers"
      ],
      "properties": {
        "labelers": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#labelerPrefItem"
          }
        }
      }
    },
    "mutedWord": {
      "type": "object",
      "description": "A word that the account owner has muted.",
      "required": [
        "value",
        "targets"
      ],
      "properties": {
        "actorTarget": {
          "type": "string",
          "default": "all"
        },
        "expiresAt": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "targets": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.actor.defs#mutedWordTarget"
          }
        },
        "value": {
          "type": "string"
        }
      }
    },
    "mutedWordTarget": {
      "type": "string",
      "description": ""
    },
    "mutedWordsPref": {
      "type": "object",
      "description": "",
      "required": [
        "items"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.actor.defs#mutedWord"
          }
        }
      }
    },
    "nux": {
      "type": "object",
      "description": "A new user experiences (NUX) storage object",
      "required": [
        "id",
        "completed"
      ],
      "properties": {
        "completed": {
          "type": "boolean",
          "default": false
        },
        "data": {
          "type": "string"
        },
        "expiresAt": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "personalDetailsPref": {
      "type": "object",
      "description": "",
      "properties": {
        "birthDate": {
          "type": "string"
        }
      }
    },
    "postInteractionSettingsPref": {
      "type": "object",
      "description": "Default post interaction settings for the account. These values should be applied as default values when creating new posts. These refs should mirror the threadgate and postgate records exactly.",
      "properties": {
        "postgateEmbeddingRules": {
          "type": "array",
          "items": {
            "type": "union",
            "refs": [
              "app.bsky.feed.postgate#disableRule"
            ]
          }
        },
        "threadgateAllowRules": {
          "type": "array",
          "items": {
            "type": "union",
            "refs": [
              "app.bsky.feed.threadgate#mentionRule",
              "app.bsky.feed.threadgate#followerRule",
              "app.bsky.feed.threadgate#followingRule",
              "app.bsky.feed.threadgate#listRule"
            ]
          }
        }
      }
    },
    "preferences": {
      "type": "array",
      "description": "",
      "items": {
        "type": "union",
        "refs": [
          "#adultContentPref",
          "#contentLabelPref",
          "#savedFeedsPref",
          "#savedFeedsPrefV2",
          "#personalDetailsPref",
          "#declaredAgePref",
          "#feedViewPref",
          "#threadViewPref",
          "#interestsPref",
          "#mutedWordsPref",
          "#hiddenPostsPref",
          "#bskyAppStatePref",
          "#labelersPref",
          "#postInteractionSettingsPref",
          "#verificationPrefs"
        ]
      }
    },
    "profileAssociated": {
      "type": "object",
      "description": "",
      "properties": {
        "activitySubscription": {
          "type": "ref",
          "ref": "#profileAssociatedActivitySubscription"
        },
        "chat": {
          "type": "ref",
          "ref": "#profileAssociatedChat"
        },
        "feedgens": {
          "type": "integer"
        },
        "labeler": {
          "type": "boolean"
        },
        "lists": {
          "type": "integer"
        },
        "starterPacks": {
          "type": "integer"
        }
      }
    },
    "profileAssociatedActivitySubscription": {
      "type": "object",
      "description": "",
      "required": [
        "allowSubscriptions"
      ],
      "properties": {
        "allowSubscriptions": {
          "type": "string"
        }
      }
    },
    "profileAssociatedChat": {
      "type": "object",
      "description": "",
      "required": [
        "allowIncoming"
      ],
      "properties": {
        "allowIncoming": {
          "type": "string"
        }
      }
    },
    "profileView": {
      "type": "object",
      "description": "",
      "required": [
        "did",
        "handle"
      ],
      "properties": {
        "associated": {
          "type": "ref",
          "ref": "#profileAssociated"
        },
        "avatar": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "debug": {
          "type": "unknown"
        },
        "description": {
          "type": "string"
        },
        "did": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "handle": {
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
        "pronouns": {
          "type": "string"
        },
        "status": {
          "type": "ref",
          "ref": "#statusView"
        },
        "verification": {
          "type": "ref",
          "ref": "#verificationState"
        },
        "viewer": {
          "type": "ref",
          "ref": "#viewerState"
        }
      }
    },
    "profileViewBasic": {
      "type": "object",
      "description": "",
      "required": [
        "did",
        "handle"
      ],
      "properties": {
        "associated": {
          "type": "ref",
          "ref": "#profileAssociated"
        },
        "avatar": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "debug": {
          "type": "unknown"
        },
        "did": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "handle": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "pronouns": {
          "type": "string"
        },
        "status": {
          "type": "ref",
          "ref": "#statusView"
        },
        "verification": {
          "type": "ref",
          "ref": "#verificationState"
        },
        "viewer": {
          "type": "ref",
          "ref": "#viewerState"
        }
      }
    },
    "profileViewDetailed": {
      "type": "object",
      "description": "",
      "required": [
        "did",
        "handle"
      ],
      "properties": {
        "associated": {
          "type": "ref",
          "ref": "#profileAssociated"
        },
        "avatar": {
          "type": "string"
        },
        "banner": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "debug": {
          "type": "unknown"
        },
        "description": {
          "type": "string"
        },
        "did": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "followersCount": {
          "type": "integer"
        },
        "followsCount": {
          "type": "integer"
        },
        "handle": {
          "type": "string"
        },
        "indexedAt": {
          "type": "string"
        },
        "joinedViaStarterPack": {
          "type": "ref",
          "ref": "app.bsky.graph.defs#starterPackViewBasic"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "pinnedPost": {
          "type": "ref",
          "ref": "com.atproto.repo.strongRef"
        },
        "postsCount": {
          "type": "integer"
        },
        "pronouns": {
          "type": "string"
        },
        "status": {
          "type": "ref",
          "ref": "#statusView"
        },
        "verification": {
          "type": "ref",
          "ref": "#verificationState"
        },
        "viewer": {
          "type": "ref",
          "ref": "#viewerState"
        },
        "website": {
          "type": "string"
        }
      }
    },
    "savedFeed": {
      "type": "object",
      "description": "",
      "required": [
        "id",
        "type",
        "value",
        "pinned"
      ],
      "properties": {
        "id": {
          "type": "string"
        },
        "pinned": {
          "type": "boolean"
        },
        "type": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "savedFeedsPref": {
      "type": "object",
      "description": "",
      "required": [
        "pinned",
        "saved"
      ],
      "properties": {
        "pinned": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "saved": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timelineIndex": {
          "type": "integer"
        }
      }
    },
    "savedFeedsPrefV2": {
      "type": "object",
      "description": "",
      "required": [
        "items"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.actor.defs#savedFeed"
          }
        }
      }
    },
    "statusView": {
      "type": "object",
      "description": "",
      "required": [
        "status",
        "record"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "embed": {
          "type": "union",
          "refs": [
            "app.bsky.embed.external#view"
          ]
        },
        "expiresAt": {
          "type": "string"
        },
        "isActive": {
          "type": "boolean"
        },
        "isDisabled": {
          "type": "boolean"
        },
        "record": {
          "type": "unknown"
        },
        "status": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "threadViewPref": {
      "type": "object",
      "description": "",
      "properties": {
        "sort": {
          "type": "string"
        }
      }
    },
    "verificationPrefs": {
      "type": "object",
      "description": "Preferences for how verified accounts appear in the app.",
      "properties": {
        "hideBadges": {
          "type": "boolean",
          "default": false
        }
      }
    },
    "verificationState": {
      "type": "object",
      "description": "Represents the verification information about the user this object is attached to.",
      "required": [
        "verifications",
        "verifiedStatus",
        "trustedVerifierStatus"
      ],
      "properties": {
        "trustedVerifierStatus": {
          "type": "string"
        },
        "verifications": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#verificationView"
          }
        },
        "verifiedStatus": {
          "type": "string"
        }
      }
    },
    "verificationView": {
      "type": "object",
      "description": "An individual verification for an associated subject.",
      "required": [
        "issuer",
        "uri",
        "isValid",
        "createdAt"
      ],
      "properties": {
        "createdAt": {
          "type": "string"
        },
        "isValid": {
          "type": "boolean"
        },
        "issuer": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "viewerState": {
      "type": "object",
      "description": "Metadata about the requesting account's relationship with the subject account. Only has meaningful content for authed requests.",
      "properties": {
        "activitySubscription": {
          "type": "ref",
          "ref": "app.bsky.notification.defs#activitySubscription"
        },
        "blockedBy": {
          "type": "boolean"
        },
        "blocking": {
          "type": "string"
        },
        "blockingByList": {
          "type": "ref",
          "ref": "app.bsky.graph.defs#listViewBasic"
        },
        "followedBy": {
          "type": "string"
        },
        "following": {
          "type": "string"
        },
        "knownFollowers": {
          "type": "ref",
          "ref": "#knownFollowers"
        },
        "muted": {
          "type": "boolean"
        },
        "mutedByList": {
          "type": "ref",
          "ref": "app.bsky.graph.defs#listViewBasic"
        }
      }
    }
  }
}
*/
