// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.moderation.defs

// Spam: frequent unwanted promotion, replies, mentions. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingSpam`.
const ModerationDefs_ReasonSpam string = "reasonSpam"

// Direct violation of server rules, laws, terms of service. Prefer new lexicon definition `tools.ozone.report.defs#reasonRuleOther`.
const ModerationDefs_ReasonViolation string = "reasonViolation"

// Unwanted or mislabeled sexual content. Prefer new lexicon definition `tools.ozone.report.defs#reasonSexualUnlabeled`.
const ModerationDefs_ReasonSexual string = "reasonSexual"

// Rude, harassing, explicit, or otherwise unwelcoming behavior. Prefer new lexicon definition `tools.ozone.report.defs#reasonHarassmentOther`.
const ModerationDefs_ReasonRude string = "reasonRude"

// Appeal a previously taken moderation action
const ModerationDefs_ReasonAppeal string = "reasonAppeal"

type ModerationDefs_ReasonType string

// Misleading identity, affiliation, or content. Prefer new lexicon definition `tools.ozone.report.defs#reasonMisleadingOther`.
const ModerationDefs_ReasonMisleading string = "reasonMisleading"

// Reports not falling under another report category. Prefer new lexicon definition `tools.ozone.report.defs#reasonOther`.
const ModerationDefs_ReasonOther string = "reasonOther"

type ModerationDefs_SubjectType string

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
