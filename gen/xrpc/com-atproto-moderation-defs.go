// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.moderation.defs

// Appeal a previously taken moderation action
const ComATProtoModerationDefs_ReasonAppeal string = "reasonAppeal"

// Misleading identity, affiliation, or content. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingOther`.
const ComATProtoModerationDefs_ReasonMisleading string = "reasonMisleading"

// Reports not falling under another report category. Prefer new lexicon definition `tools.ozone.report.defs#reasonOther`.
const ComATProtoModerationDefs_ReasonOther string = "reasonOther"

// Rude, harassing, explicit, or otherwise unwelcoming behavior. Prefer new lexicon definition `tools.ozone.report.defs#reasonHarassmentOther`.
const ComATProtoModerationDefs_ReasonRude string = "reasonRude"

// Unwanted or mislabeled sexual content. Prefer new lexicon definition `tools.ozone.report.defs#reasonSexualUnlabeled`.
const ComATProtoModerationDefs_ReasonSexual string = "reasonSexual"

// Spam: frequent unwanted promotion, replies, mentions. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingSpam`.
const ComATProtoModerationDefs_ReasonSpam string = "reasonSpam"

type ComATProtoModerationDefs_ReasonType string

// Direct violation of server rules, laws, terms of service. Prefer new lexicon definition `tools.ozone.report.defs#reasonRuleOther`.
const ComATProtoModerationDefs_ReasonViolation string = "reasonViolation"

type ComATProtoModerationDefs_SubjectType string

/*
{
  "lexicon": 1,
  "id": "com.atproto.moderation.defs",
  "description": "",
  "defs": {
    "reasonAppeal": {
      "type": "token",
      "description": "Appeal a previously taken moderation action"
    },
    "reasonMisleading": {
      "type": "token",
      "description": "Misleading identity, affiliation, or content. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingOther`."
    },
    "reasonOther": {
      "type": "token",
      "description": "Reports not falling under another report category. Prefer new lexicon definition `tools.ozone.report.defs#reasonOther`."
    },
    "reasonRude": {
      "type": "token",
      "description": "Rude, harassing, explicit, or otherwise unwelcoming behavior. Prefer new lexicon definition `tools.ozone.report.defs#reasonHarassmentOther`."
    },
    "reasonSexual": {
      "type": "token",
      "description": "Unwanted or mislabeled sexual content. Prefer new lexicon definition `tools.ozone.report.defs#reasonSexualUnlabeled`."
    },
    "reasonSpam": {
      "type": "token",
      "description": "Spam: frequent unwanted promotion, replies, mentions. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingSpam`."
    },
    "reasonType": {
      "type": "string",
      "description": ""
    },
    "reasonViolation": {
      "type": "token",
      "description": "Direct violation of server rules, laws, terms of service. Prefer new lexicon definition `tools.ozone.report.defs#reasonRuleOther`."
    },
    "subjectType": {
      "type": "string",
      "description": "Tag describing a type of subject that might be reported."
    }
  }
}
*/
