// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(in *jlexer.Lexer, out *EventMetaData) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user":
			if in.IsNull() {
				in.Skip()
				out.Usr = nil
			} else {
				if out.Usr == nil {
					out.Usr = new(User)
				}
				(*out.Usr).UnmarshalEasyJSON(in)
			}
		case "bid":
			out.Bid = uint(in.Uint())
		case "cid":
			out.Cid = uint(in.Uint())
		case "tid":
			out.Tid = uint(in.Uint())
		case "entityData":
			out.EntityData = string(in.String())
		case "about":
			out.About = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(out *jwriter.Writer, in EventMetaData) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Usr != nil {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Usr).MarshalEasyJSON(out)
	}
	if in.Bid != 0 {
		const prefix string = ",\"bid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.Bid))
	}
	if in.Cid != 0 {
		const prefix string = ",\"cid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.Cid))
	}
	if in.Tid != 0 {
		const prefix string = ",\"tid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.Tid))
	}
	if in.EntityData != "" {
		const prefix string = ",\"entityData\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EntityData))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.About))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventMetaData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventMetaData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventMetaData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventMetaData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(l, v)
}
func easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(in *jlexer.Lexer, out *Events) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Events, 0, 1)
			} else {
				*out = Events{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Event
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(out *jwriter.Writer, in Events) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Events) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Events) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Events) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Events) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(l, v)
}
func easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(in *jlexer.Lexer, out *Event) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventType":
			out.EventType = string(in.String())
		case "createAt":
			out.CreateAt = int64(in.Int64())
		case "isRead":
			out.IsRead = bool(in.Bool())
		case "uid":
			out.Uid = uint(in.Uint())
		case "makeUser":
			if in.IsNull() {
				in.Skip()
				out.MakeUsr = nil
			} else {
				if out.MakeUsr == nil {
					out.MakeUsr = new(User)
				}
				(*out.MakeUsr).UnmarshalEasyJSON(in)
			}
		case "metaData":
			(out.MetaData).UnmarshalEasyJSON(in)
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
func easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EventType))
	}
	{
		const prefix string = ",\"createAt\":"
		out.RawString(prefix)
		out.Int64(int64(in.CreateAt))
	}
	{
		const prefix string = ",\"isRead\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRead))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Uint(uint(in.Uid))
	}
	{
		const prefix string = ",\"makeUser\":"
		out.RawString(prefix)
		if in.MakeUsr == nil {
			out.RawString("null")
		} else {
			(*in.MakeUsr).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"metaData\":"
		out.RawString(prefix)
		(in.MetaData).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(l, v)
}
