// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	_v3 "github.com/bsm/openrtb/v3"
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

func easyjson3c9ce8c3DecodeGithubComBsmOpenrtbV3NativeRequest(in *jlexer.Lexer, out *Video) {
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
		case "mimes":
			if in.IsNull() {
				in.Skip()
				out.MIMEs = nil
			} else {
				in.Delim('[')
				if out.MIMEs == nil {
					if !in.IsDelim(']') {
						out.MIMEs = make([]string, 0, 4)
					} else {
						out.MIMEs = []string{}
					}
				} else {
					out.MIMEs = (out.MIMEs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.MIMEs = append(out.MIMEs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minduration":
			out.MinDuration = int(in.Int())
		case "maxduration":
			out.MaxDuration = int(in.Int())
		case "protocols":
			if in.IsNull() {
				in.Skip()
				out.Protocols = nil
			} else {
				in.Delim('[')
				if out.Protocols == nil {
					if !in.IsDelim(']') {
						out.Protocols = make([]_v3.Protocol, 0, 8)
					} else {
						out.Protocols = []_v3.Protocol{}
					}
				} else {
					out.Protocols = (out.Protocols)[:0]
				}
				for !in.IsDelim(']') {
					var v2 _v3.Protocol
					v2 = _v3.Protocol(in.Int())
					out.Protocols = append(out.Protocols, v2)
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
func easyjson3c9ce8c3EncodeGithubComBsmOpenrtbV3NativeRequest(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.MIMEs) != 0 {
		const prefix string = ",\"mimes\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v3, v4 := range in.MIMEs {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.MinDuration != 0 {
		const prefix string = ",\"minduration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MinDuration))
	}
	if in.MaxDuration != 0 {
		const prefix string = ",\"maxduration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MaxDuration))
	}
	if len(in.Protocols) != 0 {
		const prefix string = ",\"protocols\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Protocols {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Video) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9ce8c3EncodeGithubComBsmOpenrtbV3NativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Video) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9ce8c3EncodeGithubComBsmOpenrtbV3NativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Video) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9ce8c3DecodeGithubComBsmOpenrtbV3NativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Video) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9ce8c3DecodeGithubComBsmOpenrtbV3NativeRequest(l, v)
}