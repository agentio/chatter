// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.recordWithMedia

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyEmbedRecordWithMedia_View_Description = ""

// AppBskyEmbedRecordWithMedia_View is a record with Lexicon type app.bsky.embed.recordWithMedia#view
type AppBskyEmbedRecordWithMedia_View struct {
	LexiconTypeID string                                  `json:"$type,omitempty"`
	Media         *AppBskyEmbedRecordWithMedia_View_Media `json:"media,omitempty"`
	Record        *AppBskyEmbedRecord_View                `json:"record,omitempty"`
}

// AppBskyEmbedRecordWithMedia_View_Media is a union with these possible values:
// - AppBskyEmbedImages_View (app.bsky.embed.images#view)
// - AppBskyEmbedVideo_View (app.bsky.embed.video#view)
// - AppBskyEmbedExternal_View (app.bsky.embed.external#view)
type AppBskyEmbedRecordWithMedia_View_Media struct {
	Wrapper AppBskyEmbedRecordWithMedia_View_Media_Wrapper
}

// Value wrappers must conform to AppBskyEmbedRecordWithMedia_View_Media_Wrapper
type AppBskyEmbedRecordWithMedia_View_Media_Wrapper interface {
	isAppBskyEmbedRecordWithMedia_View_Media()
}

// AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedImages_View wraps values of type *AppBskyEmbedImages_View
type AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedImages_View struct {
	Value *AppBskyEmbedImages_View
}

func (t AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedImages_View) isAppBskyEmbedRecordWithMedia_View_Media() {
}

// AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedVideo_View wraps values of type *AppBskyEmbedVideo_View
type AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedVideo_View struct {
	Value *AppBskyEmbedVideo_View
}

func (t AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedVideo_View) isAppBskyEmbedRecordWithMedia_View_Media() {
}

// AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedExternal_View wraps values of type *AppBskyEmbedExternal_View
type AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedExternal_View struct {
	Value *AppBskyEmbedExternal_View
}

func (t AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedExternal_View) isAppBskyEmbedRecordWithMedia_View_Media() {
}

func (u AppBskyEmbedRecordWithMedia_View_Media) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedImages_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.images#view", v.Value)
	case AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedVideo_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.video#view", v.Value)
	case AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedExternal_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.external#view", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyEmbedRecordWithMedia_View_Media) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.images#view":
		var v AppBskyEmbedImages_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedImages_View{Value: &v}
		return nil
	case "app.bsky.embed.video#view":
		var v AppBskyEmbedVideo_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedVideo_View{Value: &v}
		return nil
	case "app.bsky.embed.external#view":
		var v AppBskyEmbedExternal_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyEmbedRecordWithMedia_View_Media__AppBskyEmbedExternal_View{Value: &v}
		return nil
	default:
		return nil
	}
}
