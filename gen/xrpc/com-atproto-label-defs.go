// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.label.defs

const ComATProtoLabelDefs_Label_Description = "Metadata tag on an atproto resource (eg, repo or record)."

type ComATProtoLabelDefs_Label struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Cid           *string `json:"cid,omitempty"`
	Cts           string  `json:"cts"`
	Exp           *string `json:"exp,omitempty"`
	Neg           *bool   `json:"neg,omitempty"`
	Sig           *[]byte `json:"sig,omitempty"`
	Src           string  `json:"src"`
	Uri           string  `json:"uri"`
	Val           string  `json:"val"`
	Ver           *int64  `json:"ver,omitempty"`
}

type ComATProtoLabelDefs_LabelValue string

const ComATProtoLabelDefs_LabelValueDefinition_Description = "Declares a label value and its expected interpretations and behaviors."

type ComATProtoLabelDefs_LabelValueDefinition struct {
	LexiconTypeID  string                                             `json:"$type,omitempty"`
	AdultOnly      *bool                                              `json:"adultOnly,omitempty"`
	Blurs          string                                             `json:"blurs"`
	DefaultSetting *string                                            `json:"defaultSetting,omitempty"`
	Identifier     string                                             `json:"identifier"`
	Locales        []*ComATProtoLabelDefs_LabelValueDefinitionStrings `json:"locales,omitempty"`
	Severity       string                                             `json:"severity"`
}

const ComATProtoLabelDefs_LabelValueDefinitionStrings_Description = "Strings which describe the label in the UI, localized into a specific language."

type ComATProtoLabelDefs_LabelValueDefinitionStrings struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Description   string `json:"description"`
	Lang          string `json:"lang"`
	Name          string `json:"name"`
}

const ComATProtoLabelDefs_SelfLabel_Description = "Metadata tag on an atproto record, published by the author within the record. Note that schemas should use #selfLabels, not #selfLabel."

type ComATProtoLabelDefs_SelfLabel struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Val           string `json:"val"`
}

const ComATProtoLabelDefs_SelfLabels_Description = "Metadata tags on an atproto record, published by the author within the record."

type ComATProtoLabelDefs_SelfLabels struct {
	LexiconTypeID string                           `json:"$type,omitempty"`
	Values        []*ComATProtoLabelDefs_SelfLabel `json:"values,omitempty"`
}

/*
{
  "lexicon": 1,
  "id": "com.atproto.label.defs",
  "description": "",
  "defs": {
    "label": {
      "type": "object",
      "description": "Metadata tag on an atproto resource (eg, repo or record).",
      "required": [
        "src",
        "uri",
        "val",
        "cts"
      ],
      "properties": {
        "cid": {
          "type": "string"
        },
        "cts": {
          "type": "string"
        },
        "exp": {
          "type": "string"
        },
        "neg": {
          "type": "boolean"
        },
        "sig": {
          "type": "bytes"
        },
        "src": {
          "type": "string"
        },
        "uri": {
          "type": "string"
        },
        "val": {
          "type": "string"
        },
        "ver": {
          "type": "integer"
        }
      }
    },
    "labelValue": {
      "type": "string",
      "description": ""
    },
    "labelValueDefinition": {
      "type": "object",
      "description": "Declares a label value and its expected interpretations and behaviors.",
      "required": [
        "identifier",
        "severity",
        "blurs",
        "locales"
      ],
      "properties": {
        "adultOnly": {
          "type": "boolean"
        },
        "blurs": {
          "type": "string"
        },
        "defaultSetting": {
          "type": "string",
          "default": "warn"
        },
        "identifier": {
          "type": "string"
        },
        "locales": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#labelValueDefinitionStrings"
          }
        },
        "severity": {
          "type": "string"
        }
      }
    },
    "labelValueDefinitionStrings": {
      "type": "object",
      "description": "Strings which describe the label in the UI, localized into a specific language.",
      "required": [
        "lang",
        "name",
        "description"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "lang": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "selfLabel": {
      "type": "object",
      "description": "Metadata tag on an atproto record, published by the author within the record. Note that schemas should use #selfLabels, not #selfLabel.",
      "required": [
        "val"
      ],
      "properties": {
        "val": {
          "type": "string"
        }
      }
    },
    "selfLabels": {
      "type": "object",
      "description": "Metadata tags on an atproto record, published by the author within the record.",
      "required": [
        "values"
      ],
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "ref",
            "ref": "#selfLabel"
          }
        }
      }
    }
  }
}
*/
