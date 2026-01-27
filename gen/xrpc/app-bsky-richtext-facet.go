// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.richtext.facet

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyRichtextFacet_ByteSlice_Description = "Specifies the sub-string range a facet feature applies to. Start index is inclusive, end index is exclusive. Indices are zero-indexed, counting bytes of the UTF-8 encoded text. NOTE: some languages, like Javascript, use UTF-16 or Unicode codepoints for string slice indexing; in these languages, convert to byte arrays before working with facets."

// AppBskyRichtextFacet_ByteSlice is a record with Lexicon type app.bsky.richtext.facet#byteSlice
type AppBskyRichtextFacet_ByteSlice struct {
	LexiconTypeID string `json:"$type,omitempty"`
	ByteEnd       int64  `json:"byteEnd"`
	ByteStart     int64  `json:"byteStart"`
}

const AppBskyRichtextFacet_Link_Description = "Facet feature for a URL. The text URL may have been simplified or truncated, but the facet reference should be a complete URL."

// AppBskyRichtextFacet_Link is a record with Lexicon type app.bsky.richtext.facet#link
type AppBskyRichtextFacet_Link struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Uri           string `json:"uri"`
}

const AppBskyRichtextFacet_Description = "Annotation of a sub-string within rich text."

// AppBskyRichtextFacet is a record with Lexicon type app.bsky.richtext.facet#main
type AppBskyRichtextFacet struct {
	LexiconTypeID string                                `json:"$type,omitempty"`
	Features      []*AppBskyRichtextFacet_Features_Elem `json:"features,omitempty"`
	Index         *AppBskyRichtextFacet_ByteSlice       `json:"index,omitempty"`
}

// AppBskyRichtextFacet_Features_Elem is a union with these possible values:
// - AppBskyRichtextFacet_Mention (app.bsky.richtext.facet#mention)
// - AppBskyRichtextFacet_Link (app.bsky.richtext.facet#link)
// - AppBskyRichtextFacet_Tag (app.bsky.richtext.facet#tag)
type AppBskyRichtextFacet_Features_Elem struct {
	Wrapper AppBskyRichtextFacet_Features_Elem_Wrapper
}

// Value wrappers must conform to AppBskyRichtextFacet_Features_Elem_Wrapper
type AppBskyRichtextFacet_Features_Elem_Wrapper interface {
	isAppBskyRichtextFacet_Features_Elem()
}

// AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Mention wraps values of type *AppBskyRichtextFacet_Mention
type AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Mention struct {
	Value *AppBskyRichtextFacet_Mention
}

func (t AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Mention) isAppBskyRichtextFacet_Features_Elem() {
}

// AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Link wraps values of type *AppBskyRichtextFacet_Link
type AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Link struct {
	Value *AppBskyRichtextFacet_Link
}

func (t AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Link) isAppBskyRichtextFacet_Features_Elem() {
}

// AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Tag wraps values of type *AppBskyRichtextFacet_Tag
type AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Tag struct {
	Value *AppBskyRichtextFacet_Tag
}

func (t AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Tag) isAppBskyRichtextFacet_Features_Elem() {
}

func (u AppBskyRichtextFacet_Features_Elem) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Mention:
		return slink.MarshalWithLexiconType("app.bsky.richtext.facet#mention", v.Value)
	case AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Link:
		return slink.MarshalWithLexiconType("app.bsky.richtext.facet#link", v.Value)
	case AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Tag:
		return slink.MarshalWithLexiconType("app.bsky.richtext.facet#tag", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyRichtextFacet_Features_Elem) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.richtext.facet#mention":
		var v AppBskyRichtextFacet_Mention
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Mention{Value: &v}
		return nil
	case "app.bsky.richtext.facet#link":
		var v AppBskyRichtextFacet_Link
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Link{Value: &v}
		return nil
	case "app.bsky.richtext.facet#tag":
		var v AppBskyRichtextFacet_Tag
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyRichtextFacet_Features_Elem__AppBskyRichtextFacet_Tag{Value: &v}
		return nil
	default:
		return nil
	}
}

const AppBskyRichtextFacet_Mention_Description = "Facet feature for mention of another account. The text is usually a handle, including a '@' prefix, but the facet reference is a DID."

// AppBskyRichtextFacet_Mention is a record with Lexicon type app.bsky.richtext.facet#mention
type AppBskyRichtextFacet_Mention struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Did           string `json:"did"`
}

const AppBskyRichtextFacet_Tag_Description = "Facet feature for a hashtag. The text usually includes a '#' prefix, but the facet reference should not (except in the case of 'double hash tags')."

// AppBskyRichtextFacet_Tag is a record with Lexicon type app.bsky.richtext.facet#tag
type AppBskyRichtextFacet_Tag struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Tag           string `json:"tag"`
}
