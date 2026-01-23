// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // chat.bsky.actor.defs

type ChatBskyActorDefs_ProfileViewBasic struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	ChatDisabled  *bool                               `json:"chatDisabled,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	Labels        []*LabelDefs_Label                  `json:"labels,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "chat.bsky.actor.defs",
  "description": "",
  "defs": {
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
          "ref": "app.bsky.actor.defs#profileAssociated"
        },
        "avatar": {
          "type": "string"
        },
        "chatDisabled": {
          "type": "boolean"
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
        "verification": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#verificationState"
        },
        "viewer": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#viewerState"
        }
      }
    }
  }
}
*/
