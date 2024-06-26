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

func easyjson3073ac56DecodeGithubComBsmOpenrtbV3(in *jlexer.Lexer, out *Device) {
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
		case "ua":
			out.UA = string(in.String())
		case "sua":
			if in.IsNull() {
				in.Skip()
				out.Sua = nil
			} else {
				if out.Sua == nil {
					out.Sua = new(UserAgent)
				}
				(*out.Sua).UnmarshalEasyJSON(in)
			}
		case "geo":
			if in.IsNull() {
				in.Skip()
				out.Geo = nil
			} else {
				if out.Geo == nil {
					out.Geo = new(Geo)
				}
				(*out.Geo).UnmarshalEasyJSON(in)
			}
		case "dnt":
			out.DNT = int(in.Int())
		case "lmt":
			out.LMT = int(in.Int())
		case "ip":
			out.IP = string(in.String())
		case "ipv6":
			out.IPv6 = string(in.String())
		case "devicetype":
			out.DeviceType = DeviceType(in.Int())
		case "make":
			out.Make = string(in.String())
		case "model":
			out.Model = string(in.String())
		case "os":
			out.OS = string(in.String())
		case "osv":
			out.OSVersion = string(in.String())
		case "hwv":
			out.HWVersion = string(in.String())
		case "h":
			out.Height = int(in.Int())
		case "w":
			out.Width = int(in.Int())
		case "ppi":
			out.PPI = int(in.Int())
		case "pxratio":
			out.PixelRatio = float64(in.Float64())
		case "js":
			out.JS = int(in.Int())
		case "geofetch":
			out.GeoFetch = int(in.Int())
		case "flashver":
			out.FlashVersion = string(in.String())
		case "language":
			out.Language = string(in.String())
		case "langb":
			out.LanguageB = string(in.String())
		case "carrier":
			out.Carrier = string(in.String())
		case "mccmnc":
			out.MCCMNC = string(in.String())
		case "connectiontype":
			out.ConnType = ConnType(in.Int())
		case "ifa":
			out.IFA = string(in.String())
		case "didsha1":
			out.IDSHA1 = string(in.String())
		case "didmd5":
			out.IDMD5 = string(in.String())
		case "dpidsha1":
			out.PIDSHA1 = string(in.String())
		case "dpidmd5":
			out.PIDMD5 = string(in.String())
		case "macsha1":
			out.MacSHA1 = string(in.String())
		case "macmd5":
			out.MacMD5 = string(in.String())
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
func easyjson3073ac56EncodeGithubComBsmOpenrtbV3(out *jwriter.Writer, in Device) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UA != "" {
		const prefix string = ",\"ua\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UA))
	}
	if in.Sua != nil {
		const prefix string = ",\"sua\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Sua).MarshalEasyJSON(out)
	}
	if in.Geo != nil {
		const prefix string = ",\"geo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Geo).MarshalEasyJSON(out)
	}
	if in.DNT != 0 {
		const prefix string = ",\"dnt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.DNT))
	}
	if in.LMT != 0 {
		const prefix string = ",\"lmt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.LMT))
	}
	if in.IP != "" {
		const prefix string = ",\"ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IP))
	}
	if in.IPv6 != "" {
		const prefix string = ",\"ipv6\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IPv6))
	}
	if in.DeviceType != 0 {
		const prefix string = ",\"devicetype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.DeviceType))
	}
	if in.Make != "" {
		const prefix string = ",\"make\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Make))
	}
	if in.Model != "" {
		const prefix string = ",\"model\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Model))
	}
	if in.OS != "" {
		const prefix string = ",\"os\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OS))
	}
	if in.OSVersion != "" {
		const prefix string = ",\"osv\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OSVersion))
	}
	if in.HWVersion != "" {
		const prefix string = ",\"hwv\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HWVersion))
	}
	if in.Height != 0 {
		const prefix string = ",\"h\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Height))
	}
	if in.Width != 0 {
		const prefix string = ",\"w\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Width))
	}
	if in.PPI != 0 {
		const prefix string = ",\"ppi\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PPI))
	}
	if in.PixelRatio != 0 {
		const prefix string = ",\"pxratio\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.PixelRatio))
	}
	if in.JS != 0 {
		const prefix string = ",\"js\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.JS))
	}
	if in.GeoFetch != 0 {
		const prefix string = ",\"geofetch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GeoFetch))
	}
	if in.FlashVersion != "" {
		const prefix string = ",\"flashver\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FlashVersion))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	if in.LanguageB != "" {
		const prefix string = ",\"langb\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LanguageB))
	}
	if in.Carrier != "" {
		const prefix string = ",\"carrier\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Carrier))
	}
	if in.MCCMNC != "" {
		const prefix string = ",\"mccmnc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MCCMNC))
	}
	if in.ConnType != 0 {
		const prefix string = ",\"connectiontype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ConnType))
	}
	if in.IFA != "" {
		const prefix string = ",\"ifa\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IFA))
	}
	if in.IDSHA1 != "" {
		const prefix string = ",\"didsha1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IDSHA1))
	}
	if in.IDMD5 != "" {
		const prefix string = ",\"didmd5\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IDMD5))
	}
	if in.PIDSHA1 != "" {
		const prefix string = ",\"dpidsha1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PIDSHA1))
	}
	if in.PIDMD5 != "" {
		const prefix string = ",\"dpidmd5\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PIDMD5))
	}
	if in.MacSHA1 != "" {
		const prefix string = ",\"macsha1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MacSHA1))
	}
	if in.MacMD5 != "" {
		const prefix string = ",\"macmd5\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MacMD5))
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
func (v Device) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3073ac56EncodeGithubComBsmOpenrtbV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Device) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3073ac56EncodeGithubComBsmOpenrtbV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Device) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3073ac56DecodeGithubComBsmOpenrtbV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Device) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3073ac56DecodeGithubComBsmOpenrtbV3(l, v)
}
