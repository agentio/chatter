// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.graph.defs

// A list of actors to apply an aggregate moderation action (mute/block) on.
const AppBskyGraphDefs_Modlist string = "modlist"

// A list of actors used for only for reference purposes such as within a starter pack.
const AppBskyGraphDefs_Referencelist string = "referencelist"

type AppBskyGraphDefs_ListViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Blocked       *string `json:"blocked,omitempty"`
	Muted         *bool   `json:"muted,omitempty"`
}

// indicates that a handle or DID could not be resolved
type AppBskyGraphDefs_NotFoundActor struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Actor         string `json:"actor,omitempty"`
	NotFound      bool   `json:"notFound"`
}

// lists the bi-directional graph relationships between one actor (not indicated in the object), and the target actors (the DID included in the object)
type AppBskyGraphDefs_Relationship struct {
	LexiconTypeID  string  `json:"$type,omitempty"`
	BlockedBy      *string `json:"blockedBy,omitempty"`
	BlockedByList  *string `json:"blockedByList,omitempty"`
	Blocking       *string `json:"blocking,omitempty"`
	BlockingByList *string `json:"blockingByList,omitempty"`
	Did            string  `json:"did,omitempty"`
	FollowedBy     *string `json:"followedBy,omitempty"`
	Following      *string `json:"following,omitempty"`
}

type AppBskyGraphDefs_ListItemView struct {
	LexiconTypeID string                        `json:"$type,omitempty"`
	Subject       *AppBskyActorDefs_ProfileView `json:"subject,omitempty"`
	Uri           string                        `json:"uri,omitempty"`
}

type AppBskyGraphDefs_StarterPackView struct {
	LexiconTypeID      string                             `json:"$type,omitempty"`
	Cid                string                             `json:"cid,omitempty"`
	Creator            *AppBskyActorDefs_ProfileViewBasic `json:"creator,omitempty"`
	Feeds              []*AppBskyFeedDefs_GeneratorView   `json:"feeds,omitempty"`
	IndexedAt          string                             `json:"indexedAt,omitempty"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty"`
	Labels             []*LabelDefs_Label                 `json:"labels,omitempty"`
	List               *AppBskyGraphDefs_ListViewBasic    `json:"list,omitempty"`
	ListItemsSample    []*AppBskyGraphDefs_ListItemView   `json:"listItemsSample,omitempty"`
	Record             any                                `json:"record,omitempty"`
	Uri                string                             `json:"uri,omitempty"`
}

type AppBskyGraphDefs_StarterPackViewBasic struct {
	LexiconTypeID      string                             `json:"$type,omitempty"`
	Cid                string                             `json:"cid,omitempty"`
	Creator            *AppBskyActorDefs_ProfileViewBasic `json:"creator,omitempty"`
	IndexedAt          string                             `json:"indexedAt,omitempty"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty"`
	Labels             []*LabelDefs_Label                 `json:"labels,omitempty"`
	ListItemCount      *int64                             `json:"listItemCount,omitempty"`
	Record             any                                `json:"record,omitempty"`
	Uri                string                             `json:"uri,omitempty"`
}

// A list of actors used for curation purposes such as list feeds or interaction gating.
const AppBskyGraphDefs_Curatelist string = "curatelist"

type AppBskyGraphDefs_ListViewBasic struct {
	LexiconTypeID string                            `json:"$type,omitempty"`
	Avatar        *string                           `json:"avatar,omitempty"`
	Cid           string                            `json:"cid,omitempty"`
	IndexedAt     *string                           `json:"indexedAt,omitempty"`
	Labels        []*LabelDefs_Label                `json:"labels,omitempty"`
	ListItemCount *int64                            `json:"listItemCount,omitempty"`
	Name          string                            `json:"name,omitempty"`
	Purpose       *AppBskyGraphDefs_ListPurpose     `json:"purpose,omitempty"`
	Uri           string                            `json:"uri,omitempty"`
	Viewer        *AppBskyGraphDefs_ListViewerState `json:"viewer,omitempty"`
}

type AppBskyGraphDefs_ListView struct {
	LexiconTypeID     string                            `json:"$type,omitempty"`
	Avatar            *string                           `json:"avatar,omitempty"`
	Cid               string                            `json:"cid,omitempty"`
	Creator           *AppBskyActorDefs_ProfileView     `json:"creator,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	DescriptionFacets []*AppBskyRichtextFacet           `json:"descriptionFacets,omitempty"`
	IndexedAt         string                            `json:"indexedAt,omitempty"`
	Labels            []*LabelDefs_Label                `json:"labels,omitempty"`
	ListItemCount     *int64                            `json:"listItemCount,omitempty"`
	Name              string                            `json:"name,omitempty"`
	Purpose           *AppBskyGraphDefs_ListPurpose     `json:"purpose,omitempty"`
	Uri               string                            `json:"uri,omitempty"`
	Viewer            *AppBskyGraphDefs_ListViewerState `json:"viewer,omitempty"`
}

type AppBskyGraphDefs_ListPurpose string

/*
{
  "lexicon": 1,
  "id": "app.bsky.graph.defs",
  "description": "",
  "defs": {
    "curatelist": {
      "type": "token",
      "description": "A list of actors used for curation purposes such as list feeds or interaction gating."
    },
    "listItemView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "subject"
      ],
      "properties": {
        "subject": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileView"
        },
        "uri": {
          "type": "string"
        }
      }
    },
    "listPurpose": {
      "type": "string",
      "description": ""
    },
    "listView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "creator",
        "name",
        "purpose",
        "indexedAt"
      ],
      "properties": {
        "avatar": {
          "type": "string"
        },
        "cid": {
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
        "listItemCount": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "purpose": {
          "type": "ref",
          "ref": "#listPurpose"
        },
        "uri": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "#listViewerState"
        }
      }
    },
    "listViewBasic": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "name",
        "purpose"
      ],
      "properties": {
        "avatar": {
          "type": "string"
        },
        "cid": {
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
        "listItemCount": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "purpose": {
          "type": "ref",
          "ref": "#listPurpose"
        },
        "uri": {
          "type": "string"
        },
        "viewer": {
          "type": "ref",
          "ref": "#listViewerState"
        }
      }
    },
    "listViewerState": {
      "type": "object",
      "description": "",
      "properties": {
        "blocked": {
          "type": "string"
        },
        "muted": {
          "type": "boolean"
        }
      }
    },
    "modlist": {
      "type": "token",
      "description": "A list of actors to apply an aggregate moderation action (mute/block) on."
    },
    "notFoundActor": {
      "type": "object",
      "description": "indicates that a handle or DID could not be resolved",
      "required": [
        "actor",
        "notFound"
      ],
      "properties": {
        "actor": {
          "type": "string"
        },
        "notFound": {
          "type": "boolean"
        }
      }
    },
    "referencelist": {
      "type": "token",
      "description": "A list of actors used for only for reference purposes such as within a starter pack."
    },
    "relationship": {
      "type": "object",
      "description": "lists the bi-directional graph relationships between one actor (not indicated in the object), and the target actors (the DID included in the object)",
      "required": [
        "did"
      ],
      "properties": {
        "blockedBy": {
          "type": "string"
        },
        "blockedByList": {
          "type": "string"
        },
        "blocking": {
          "type": "string"
        },
        "blockingByList": {
          "type": "string"
        },
        "did": {
          "type": "string"
        },
        "followedBy": {
          "type": "string"
        },
        "following": {
          "type": "string"
        }
      }
    },
    "starterPackView": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "record",
        "creator",
        "indexedAt"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "creator": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "feeds": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "app.bsky.feed.defs#generatorView"
          }
        },
        "indexedAt": {
          "type": "string"
        },
        "joinedAllTimeCount": {
          "type": "integer"
        },
        "joinedWeekCount": {
          "type": "integer"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "list": {
          "type": "ref",
          "ref": "#listViewBasic"
        },
        "listItemsSample": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#listItemView"
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
    "starterPackViewBasic": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid",
        "record",
        "creator",
        "indexedAt"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "creator": {
          "type": "ref",
          "ref": "app.bsky.actor.defs#profileViewBasic"
        },
        "indexedAt": {
          "type": "string"
        },
        "joinedAllTimeCount": {
          "type": "integer"
        },
        "joinedWeekCount": {
          "type": "integer"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "com.atproto.label.defs#label"
          }
        },
        "listItemCount": {
          "type": "integer"
        },
        "record": {
          "type": "unknown"
        },
        "uri": {
          "type": "string"
        }
      }
    }
  }
}
*/
