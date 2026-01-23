// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.repo.strongRef

const ComATProtoRepoStrongRef_Description = ""

type ComATProtoRepoStrongRef struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Cid           string `json:"cid"`
	Uri           string `json:"uri"`
}

/*
{
  "lexicon": 1,
  "id": "com.atproto.repo.strongRef",
  "description": "A URI with a content-hash fingerprint.",
  "defs": {
    "main": {
      "type": "object",
      "description": "",
      "required": [
        "uri",
        "cid"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        }
      }
    }
  }
}
*/
