// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.notification.defs

// Object used to store activity subscription data in stash.
type AppBskyNotificationDefs_SubjectActivitySubscription struct {
	LexiconTypeID        string                                        `json:"$type,omitempty"`
	ActivitySubscription *AppBskyNotificationDefs_ActivitySubscription `json:"activitySubscription,omitempty"`
	Subject              string                                        `json:"subject,omitempty"`
}

type AppBskyNotificationDefs_RecordDeleted struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

type AppBskyNotificationDefs_ChatPreference struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Include       string `json:"include,omitempty"`
	Push          bool   `json:"push"`
}

type AppBskyNotificationDefs_FilterablePreference struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Include       string `json:"include,omitempty"`
	List          bool   `json:"list"`
	Push          bool   `json:"push"`
}

type AppBskyNotificationDefs_Preference struct {
	LexiconTypeID string `json:"$type,omitempty"`
	List          bool   `json:"list"`
	Push          bool   `json:"push"`
}

type AppBskyNotificationDefs_Preferences struct {
	LexiconTypeID     string                                        `json:"$type,omitempty"`
	Chat              *AppBskyNotificationDefs_ChatPreference       `json:"chat,omitempty"`
	Follow            *AppBskyNotificationDefs_FilterablePreference `json:"follow,omitempty"`
	Like              *AppBskyNotificationDefs_FilterablePreference `json:"like,omitempty"`
	LikeViaRepost     *AppBskyNotificationDefs_FilterablePreference `json:"likeViaRepost,omitempty"`
	Mention           *AppBskyNotificationDefs_FilterablePreference `json:"mention,omitempty"`
	Quote             *AppBskyNotificationDefs_FilterablePreference `json:"quote,omitempty"`
	Reply             *AppBskyNotificationDefs_FilterablePreference `json:"reply,omitempty"`
	Repost            *AppBskyNotificationDefs_FilterablePreference `json:"repost,omitempty"`
	RepostViaRepost   *AppBskyNotificationDefs_FilterablePreference `json:"repostViaRepost,omitempty"`
	StarterpackJoined *AppBskyNotificationDefs_Preference           `json:"starterpackJoined,omitempty"`
	SubscribedPost    *AppBskyNotificationDefs_Preference           `json:"subscribedPost,omitempty"`
	Unverified        *AppBskyNotificationDefs_Preference           `json:"unverified,omitempty"`
	Verified          *AppBskyNotificationDefs_Preference           `json:"verified,omitempty"`
}

type AppBskyNotificationDefs_ActivitySubscription struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Post          bool   `json:"post"`
	Reply         bool   `json:"reply"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.notification.defs",
  "description": "",
  "defs": {
    "activitySubscription": {
      "type": "object",
      "description": "",
      "required": [
        "post",
        "reply"
      ],
      "properties": {
        "post": {
          "type": "boolean"
        },
        "reply": {
          "type": "boolean"
        }
      }
    },
    "chatPreference": {
      "type": "object",
      "description": "",
      "required": [
        "include",
        "push"
      ],
      "properties": {
        "include": {
          "type": "string"
        },
        "push": {
          "type": "boolean"
        }
      }
    },
    "filterablePreference": {
      "type": "object",
      "description": "",
      "required": [
        "include",
        "list",
        "push"
      ],
      "properties": {
        "include": {
          "type": "string"
        },
        "list": {
          "type": "boolean"
        },
        "push": {
          "type": "boolean"
        }
      }
    },
    "preference": {
      "type": "object",
      "description": "",
      "required": [
        "list",
        "push"
      ],
      "properties": {
        "list": {
          "type": "boolean"
        },
        "push": {
          "type": "boolean"
        }
      }
    },
    "preferences": {
      "type": "object",
      "description": "",
      "required": [
        "chat",
        "follow",
        "like",
        "likeViaRepost",
        "mention",
        "quote",
        "reply",
        "repost",
        "repostViaRepost",
        "starterpackJoined",
        "subscribedPost",
        "unverified",
        "verified"
      ],
      "properties": {
        "chat": {
          "type": "ref",
          "ref": "#chatPreference"
        },
        "follow": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "like": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "likeViaRepost": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "mention": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "quote": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "reply": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "repost": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "repostViaRepost": {
          "type": "ref",
          "ref": "#filterablePreference"
        },
        "starterpackJoined": {
          "type": "ref",
          "ref": "#preference"
        },
        "subscribedPost": {
          "type": "ref",
          "ref": "#preference"
        },
        "unverified": {
          "type": "ref",
          "ref": "#preference"
        },
        "verified": {
          "type": "ref",
          "ref": "#preference"
        }
      }
    },
    "recordDeleted": {
      "type": "object",
      "description": ""
    },
    "subjectActivitySubscription": {
      "type": "object",
      "description": "Object used to store activity subscription data in stash.",
      "required": [
        "subject",
        "activitySubscription"
      ],
      "properties": {
        "activitySubscription": {
          "type": "ref",
          "ref": "#activitySubscription"
        },
        "subject": {
          "type": "string"
        }
      }
    }
  }
}
*/
