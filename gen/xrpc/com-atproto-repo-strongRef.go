// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.repo.strongRef

type RepoStrongRef struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Cid           string `json:"cid,omitempty"`
	Uri           string `json:"uri,omitempty"`
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
