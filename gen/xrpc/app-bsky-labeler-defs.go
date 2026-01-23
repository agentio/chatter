// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.labeler.defs

type AppBskyLabelerDefs_LabelerView struct {
	LexiconTypeID string                                 `json:"$type,omitempty"`
	Cid           string                                 `json:"cid,omitempty"`
	Creator       *AppBskyActorDefs_ProfileView          `json:"creator,omitempty"`
	IndexedAt     string                                 `json:"indexedAt,omitempty"`
	Labels        []*LabelDefs_Label                     `json:"labels,omitempty"`
	LikeCount     *int64                                 `json:"likeCount,omitempty"`
	Uri           string                                 `json:"uri,omitempty"`
	Viewer        *AppBskyLabelerDefs_LabelerViewerState `json:"viewer,omitempty"`
}

type AppBskyLabelerDefs_LabelerViewDetailed struct {
	LexiconTypeID      string                                 `json:"$type,omitempty"`
	Cid                string                                 `json:"cid,omitempty"`
	Creator            *AppBskyActorDefs_ProfileView          `json:"creator,omitempty"`
	IndexedAt          string                                 `json:"indexedAt,omitempty"`
	Labels             []*LabelDefs_Label                     `json:"labels,omitempty"`
	LikeCount          *int64                                 `json:"likeCount,omitempty"`
	Policies           *AppBskyLabelerDefs_LabelerPolicies    `json:"policies,omitempty"`
	ReasonTypes        []*ModerationDefs_ReasonType           `json:"reasonTypes,omitempty"`
	SubjectCollections []string                               `json:"subjectCollections,omitempty"`
	SubjectTypes       []*ModerationDefs_SubjectType          `json:"subjectTypes,omitempty"`
	Uri                string                                 `json:"uri,omitempty"`
	Viewer             *AppBskyLabelerDefs_LabelerViewerState `json:"viewer,omitempty"`
}

type AppBskyLabelerDefs_LabelerViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Like          *string `json:"like,omitempty"`
}

type AppBskyLabelerDefs_LabelerPolicies struct {
	LexiconTypeID         string                            `json:"$type,omitempty"`
	LabelValueDefinitions []*LabelDefs_LabelValueDefinition `json:"labelValueDefinitions,omitempty"`
	LabelValues           []*LabelDefs_LabelValue           `json:"labelValues,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "app.bsky.labeler.defs",
  "description": "",
  "defs": {
    "labelerPolicies": {
      "type": "object",
      "description": "",
      "required": [
        "labelValues"
      ],
      "properties": {
        "labelValueDefinitions": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#labelValueDefinition"
          }
        },
        "labelValues": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#labelValue"
          }
        }
      }
    },
    "labelerView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "creator",
        "indexedAt"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "creator": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileView"
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
          "ref": "#labelerViewerState"
        }
      }
    },
    "labelerViewDetailed": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "creator",
        "policies",
        "indexedAt"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "creator": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileView"
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
        "policies": {
          "type": "ref",
          "ref": "app.bsky.labeler.defs#labelerPolicies"
        },
        "reasonTypes": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.moderation.defs#reasonType"
          }
        },
        "subjectCollections": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "subjectTypes": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.moderation.defs#subjectType"
          }
        },
        "uri": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "#labelerViewerState"
        }
      }
    },
    "labelerViewerState": {
      "type": "object",
      "description": "",
      "properties": {
        "like": {
          "type": "string"
        }
      }
    }
  }
}
*/
