// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

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

func easyjson526ccff2DecodeGithubComBsmOpenrtbV3(in *jlexer.Lexer, out *SeatBid) {
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
		case "bid":
			if in.IsNull() {
				in.Skip()
				out.Bids = nil
			} else {
				in.Delim('[')
				if out.Bids == nil {
					if !in.IsDelim(']') {
						out.Bids = make([]Bid, 0, 0)
					} else {
						out.Bids = []Bid{}
					}
				} else {
					out.Bids = (out.Bids)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Bid
					(v1).UnmarshalEasyJSON(in)
					out.Bids = append(out.Bids, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "seat":
			out.Seat = string(in.String())
		case "group":
			out.Group = int(in.Int())
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
func easyjson526ccff2EncodeGithubComBsmOpenrtbV3(out *jwriter.Writer, in SeatBid) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"bid\":"
		out.RawString(prefix[1:])
		if in.Bids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Bids {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Seat != "" {
		const prefix string = ",\"seat\":"
		out.RawString(prefix)
		out.String(string(in.Seat))
	}
	if in.Group != 0 {
		const prefix string = ",\"group\":"
		out.RawString(prefix)
		out.Int(int(in.Group))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SeatBid) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson526ccff2EncodeGithubComBsmOpenrtbV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SeatBid) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson526ccff2EncodeGithubComBsmOpenrtbV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SeatBid) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson526ccff2DecodeGithubComBsmOpenrtbV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SeatBid) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson526ccff2DecodeGithubComBsmOpenrtbV3(l, v)
}