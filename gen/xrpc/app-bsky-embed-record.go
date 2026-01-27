// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.record

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyEmbedRecord_Description = ""

// AppBskyEmbedRecord is a record with Lexicon type app.bsky.embed.record#main
type AppBskyEmbedRecord struct {
	LexiconTypeID string                   `json:"$type,omitempty"`
	Record        *ComATProtoRepoStrongRef `json:"record,omitempty"`
}

const AppBskyEmbedRecord_View_Description = ""

// AppBskyEmbedRecord_View is a record with Lexicon type app.bsky.embed.record#view
type AppBskyEmbedRecord_View struct {
	LexiconTypeID string                          `json:"$type,omitempty"`
	Record        *AppBskyEmbedRecord_View_Record `json:"record,omitempty"`
}

// AppBskyEmbedRecord_View_Record is a union with these possible values:
// - AppBskyEmbedRecord_ViewRecord (app.bsky.embed.record#viewRecord)
// - AppBskyEmbedRecord_ViewNotFound (app.bsky.embed.record#viewNotFound)
// - AppBskyEmbedRecord_ViewBlocked (app.bsky.embed.record#viewBlocked)
// - AppBskyEmbedRecord_ViewDetached (app.bsky.embed.record#viewDetached)
// - AppBskyFeedDefs_GeneratorView (app.bsky.feed.defs#generatorView)
// - AppBskyGraphDefs_ListView (app.bsky.graph.defs#listView)
// - AppBskyLabelerDefs_LabelerView (app.bsky.labeler.defs#labelerView)
// - AppBskyGraphDefs_StarterPackViewBasic (app.bsky.graph.defs#starterPackViewBasic)
type AppBskyEmbedRecord_View_Record struct {
	Wrapper AppBskyEmbedRecord_View_Record_Wrapper
}

// Value wrappers must conform to AppBskyEmbedRecord_View_Record_Wrapper
type AppBskyEmbedRecord_View_Record_Wrapper interface {
	isAppBskyEmbedRecord_View_Record()
}

// AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewRecord wraps values of type *AppBskyEmbedRecord_ViewRecord
type AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewRecord struct {
	Value *AppBskyEmbedRecord_ViewRecord
}

func (t AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewRecord) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewNotFound wraps values of type *AppBskyEmbedRecord_ViewNotFound
type AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewNotFound struct {
	Value *AppBskyEmbedRecord_ViewNotFound
}

func (t AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewNotFound) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewBlocked wraps values of type *AppBskyEmbedRecord_ViewBlocked
type AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewBlocked struct {
	Value *AppBskyEmbedRecord_ViewBlocked
}

func (t AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewBlocked) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewDetached wraps values of type *AppBskyEmbedRecord_ViewDetached
type AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewDetached struct {
	Value *AppBskyEmbedRecord_ViewDetached
}

func (t AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewDetached) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyFeedDefs_GeneratorView wraps values of type *AppBskyFeedDefs_GeneratorView
type AppBskyEmbedRecord_View_Record__AppBskyFeedDefs_GeneratorView struct {
	Value *AppBskyFeedDefs_GeneratorView
}

func (t AppBskyEmbedRecord_View_Record__AppBskyFeedDefs_GeneratorView) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_ListView wraps values of type *AppBskyGraphDefs_ListView
type AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_ListView struct {
	Value *AppBskyGraphDefs_ListView
}

func (t AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_ListView) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyLabelerDefs_LabelerView wraps values of type *AppBskyLabelerDefs_LabelerView
type AppBskyEmbedRecord_View_Record__AppBskyLabelerDefs_LabelerView struct {
	Value *AppBskyLabelerDefs_LabelerView
}

func (t AppBskyEmbedRecord_View_Record__AppBskyLabelerDefs_LabelerView) isAppBskyEmbedRecord_View_Record() {
}

// AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_StarterPackViewBasic wraps values of type *AppBskyGraphDefs_StarterPackViewBasic
type AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_StarterPackViewBasic struct {
	Value *AppBskyGraphDefs_StarterPackViewBasic
}

func (t AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_StarterPackViewBasic) isAppBskyEmbedRecord_View_Record() {
}

func (u AppBskyEmbedRecord_View_Record) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewRecord:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#viewRecord", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewNotFound:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#viewNotFound", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewBlocked:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#viewBlocked", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewDetached:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#viewDetached", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyFeedDefs_GeneratorView:
		return slink.MarshalWithLexiconType("app.bsky.feed.defs#generatorView", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_ListView:
		return slink.MarshalWithLexiconType("app.bsky.graph.defs#listView", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyLabelerDefs_LabelerView:
		return slink.MarshalWithLexiconType("app.bsky.labeler.defs#labelerView", v.Value)
	case AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_StarterPackViewBasic:
		return slink.MarshalWithLexiconType("app.bsky.graph.defs#starterPackViewBasic", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyEmbedRecord_View_Record) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.record#viewRecord":
		var v AppBskyEmbedRecord_ViewRecord
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewRecord{Value: &v}
		return nil
	case "app.bsky.embed.record#viewNotFound":
		var v AppBskyEmbedRecord_ViewNotFound
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewNotFound{Value: &v}
		return nil
	case "app.bsky.embed.record#viewBlocked":
		var v AppBskyEmbedRecord_ViewBlocked
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewBlocked{Value: &v}
		return nil
	case "app.bsky.embed.record#viewDetached":
		var v AppBskyEmbedRecord_ViewDetached
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyEmbedRecord_ViewDetached{Value: &v}
		return nil
	case "app.bsky.feed.defs#generatorView":
		var v AppBskyFeedDefs_GeneratorView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyFeedDefs_GeneratorView{Value: &v}
		return nil
	case "app.bsky.graph.defs#listView":
		var v AppBskyGraphDefs_ListView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_ListView{Value: &v}
		return nil
	case "app.bsky.labeler.defs#labelerView":
		var v AppBskyLabelerDefs_LabelerView
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyLabelerDefs_LabelerView{Value: &v}
		return nil
	case "app.bsky.graph.defs#starterPackViewBasic":
		var v AppBskyGraphDefs_StarterPackViewBasic
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_View_Record__AppBskyGraphDefs_StarterPackViewBasic{Value: &v}
		return nil
	default:
		return nil
	}
}

const AppBskyEmbedRecord_ViewBlocked_Description = ""

// AppBskyEmbedRecord_ViewBlocked is a record with Lexicon type app.bsky.embed.record#viewBlocked
type AppBskyEmbedRecord_ViewBlocked struct {
	LexiconTypeID string                         `json:"$type,omitempty"`
	Author        *AppBskyFeedDefs_BlockedAuthor `json:"author,omitempty"`
	Blocked       bool                           `json:"blocked"`
	Uri           string                         `json:"uri"`
}

const AppBskyEmbedRecord_ViewDetached_Description = ""

// AppBskyEmbedRecord_ViewDetached is a record with Lexicon type app.bsky.embed.record#viewDetached
type AppBskyEmbedRecord_ViewDetached struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Detached      bool   `json:"detached"`
	Uri           string `json:"uri"`
}

const AppBskyEmbedRecord_ViewNotFound_Description = ""

// AppBskyEmbedRecord_ViewNotFound is a record with Lexicon type app.bsky.embed.record#viewNotFound
type AppBskyEmbedRecord_ViewNotFound struct {
	LexiconTypeID string `json:"$type,omitempty"`
	NotFound      bool   `json:"notFound"`
	Uri           string `json:"uri"`
}

const AppBskyEmbedRecord_ViewRecord_Description = ""

// AppBskyEmbedRecord_ViewRecord is a record with Lexicon type app.bsky.embed.record#viewRecord
type AppBskyEmbedRecord_ViewRecord struct {
	LexiconTypeID string                                       `json:"$type,omitempty"`
	Author        *AppBskyActorDefs_ProfileViewBasic           `json:"author,omitempty"`
	Cid           string                                       `json:"cid"`
	Embeds        []*AppBskyEmbedRecord_ViewRecord_Embeds_Elem `json:"embeds,omitempty"`
	IndexedAt     string                                       `json:"indexedAt"`
	Labels        []*ComATProtoLabelDefs_Label                 `json:"labels,omitempty"`
	LikeCount     *int64                                       `json:"likeCount,omitempty"`
	QuoteCount    *int64                                       `json:"quoteCount,omitempty"`
	ReplyCount    *int64                                       `json:"replyCount,omitempty"`
	RepostCount   *int64                                       `json:"repostCount,omitempty"`
	Uri           string                                       `json:"uri"`
	Value         any                                          `json:"value"`
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem is a union with these possible values:
// - AppBskyEmbedImages_View (app.bsky.embed.images#view)
// - AppBskyEmbedVideo_View (app.bsky.embed.video#view)
// - AppBskyEmbedExternal_View (app.bsky.embed.external#view)
// - AppBskyEmbedRecord_View (app.bsky.embed.record#view)
// - AppBskyEmbedRecordWithMedia_View (app.bsky.embed.recordWithMedia#view)
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem struct {
	Wrapper AppBskyEmbedRecord_ViewRecord_Embeds_Elem_Wrapper
}

// Value wrappers must conform to AppBskyEmbedRecord_ViewRecord_Embeds_Elem_Wrapper
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem_Wrapper interface {
	isAppBskyEmbedRecord_ViewRecord_Embeds_Elem()
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedImages_View wraps values of type *AppBskyEmbedImages_View
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedImages_View struct {
	Value *AppBskyEmbedImages_View
}

func (t AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedImages_View) isAppBskyEmbedRecord_ViewRecord_Embeds_Elem() {
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedVideo_View wraps values of type *AppBskyEmbedVideo_View
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedVideo_View struct {
	Value *AppBskyEmbedVideo_View
}

func (t AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedVideo_View) isAppBskyEmbedRecord_ViewRecord_Embeds_Elem() {
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedExternal_View wraps values of type *AppBskyEmbedExternal_View
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedExternal_View struct {
	Value *AppBskyEmbedExternal_View
}

func (t AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedExternal_View) isAppBskyEmbedRecord_ViewRecord_Embeds_Elem() {
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecord_View wraps values of type *AppBskyEmbedRecord_View
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecord_View struct {
	Value *AppBskyEmbedRecord_View
}

func (t AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecord_View) isAppBskyEmbedRecord_ViewRecord_Embeds_Elem() {
}

// AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecordWithMedia_View wraps values of type *AppBskyEmbedRecordWithMedia_View
type AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecordWithMedia_View struct {
	Value *AppBskyEmbedRecordWithMedia_View
}

func (t AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecordWithMedia_View) isAppBskyEmbedRecord_ViewRecord_Embeds_Elem() {
}

func (u AppBskyEmbedRecord_ViewRecord_Embeds_Elem) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedImages_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.images#view", v.Value)
	case AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedVideo_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.video#view", v.Value)
	case AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedExternal_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.external#view", v.Value)
	case AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecord_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.record#view", v.Value)
	case AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecordWithMedia_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.recordWithMedia#view", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyEmbedRecord_ViewRecord_Embeds_Elem) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.images#view":
		var v AppBskyEmbedImages_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedImages_View{Value: &v}
		return nil
	case "app.bsky.embed.video#view":
		var v AppBskyEmbedVideo_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedVideo_View{Value: &v}
		return nil
	case "app.bsky.embed.external#view":
		var v AppBskyEmbedExternal_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedExternal_View{Value: &v}
		return nil
	case "app.bsky.embed.record#view":
		var v AppBskyEmbedRecord_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecord_View{Value: &v}
		return nil
	case "app.bsky.embed.recordWithMedia#view":
		var v AppBskyEmbedRecordWithMedia_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecord_ViewRecord_Embeds_Elem__AppBskyEmbedRecordWithMedia_View{Value: &v}
		return nil
	default:
		return nil
	}
}
