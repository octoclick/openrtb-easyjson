// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6ff3ac1dDecodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ver":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Version).UnmarshalJSON(data))
			}
		case "assets":
			if in.IsNull() {
				in.Skip()
				out.Assets = nil
			} else {
				in.Delim('[')
				if out.Assets == nil {
					if !in.IsDelim(']') {
						out.Assets = make([]Asset, 0, 0)
					} else {
						out.Assets = []Asset{}
					}
				} else {
					out.Assets = (out.Assets)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Asset
					(v1).UnmarshalEasyJSON(in)
					out.Assets = append(out.Assets, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "link":
			(out.Link).UnmarshalEasyJSON(in)
		case "imptrackers":
			if in.IsNull() {
				in.Skip()
				out.ImpTrackers = nil
			} else {
				in.Delim('[')
				if out.ImpTrackers == nil {
					if !in.IsDelim(']') {
						out.ImpTrackers = make([]string, 0, 4)
					} else {
						out.ImpTrackers = []string{}
					}
				} else {
					out.ImpTrackers = (out.ImpTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ImpTrackers = append(out.ImpTrackers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "jstracker":
			out.JSTracker = string(in.String())
		case "eventtrackers":
			if in.IsNull() {
				in.Skip()
				out.EventTrackers = nil
			} else {
				in.Delim('[')
				if out.EventTrackers == nil {
					if !in.IsDelim(']') {
						out.EventTrackers = make([]EventTrackers, 0, 1)
					} else {
						out.EventTrackers = []EventTrackers{}
					}
				} else {
					out.EventTrackers = (out.EventTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v3 EventTrackers
					(v3).UnmarshalEasyJSON(in)
					out.EventTrackers = append(out.EventTrackers, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6ff3ac1dEncodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Version != "" {
		const prefix string = ",\"ver\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"assets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Assets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Assets {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix)
		(in.Link).MarshalEasyJSON(out)
	}
	if len(in.ImpTrackers) != 0 {
		const prefix string = ",\"imptrackers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v6, v7 := range in.ImpTrackers {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if in.JSTracker != "" {
		const prefix string = ",\"jstracker\":"
		out.RawString(prefix)
		out.String(string(in.JSTracker))
	}
	if len(in.EventTrackers) != 0 {
		const prefix string = ",\"eventtrackers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.EventTrackers {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComOctoclickOpenrtbEasyjsonNativeResponse(l, v)
}
