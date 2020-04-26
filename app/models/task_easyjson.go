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

func easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(in *jlexer.Lexer, out *Tasks) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Tasks, 0, 1)
			} else {
				*out = Tasks{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Task
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
func easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(out *jwriter.Writer, in Tasks) {
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
func (v Tasks) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tasks) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tasks) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tasks) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels(l, v)
}
func easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(in *jlexer.Lexer, out *Task) {
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
		case "id":
			out.ID = uint(in.Uint())
		case "title":
			out.Name = string(in.String())
		case "description":
			out.About = string(in.String())
		case "level":
			out.Level = uint(in.Uint())
		case "deadline":
			out.Deadline = string(in.String())
		case "position":
			out.Pos = float64(in.Float64())
		case "cid":
			out.Cid = uint(in.Uint())
		case "members":
			if in.IsNull() {
				in.Skip()
				out.Members = nil
			} else {
				in.Delim('[')
				if out.Members == nil {
					if !in.IsDelim(']') {
						out.Members = make([]User, 0, 1)
					} else {
						out.Members = []User{}
					}
				} else {
					out.Members = (out.Members)[:0]
				}
				for !in.IsDelim(']') {
					var v4 User
					easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(in, &v4)
					out.Members = append(out.Members, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "comments":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]Comment, 0, 1)
					} else {
						out.Comments = []Comment{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v5 Comment
					(v5).UnmarshalEasyJSON(in)
					out.Comments = append(out.Comments, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "checklists":
			if in.IsNull() {
				in.Skip()
				out.Checklists = nil
			} else {
				in.Delim('[')
				if out.Checklists == nil {
					if !in.IsDelim(']') {
						out.Checklists = make([]Checklist, 0, 1)
					} else {
						out.Checklists = []Checklist{}
					}
				} else {
					out.Checklists = (out.Checklists)[:0]
				}
				for !in.IsDelim(']') {
					var v6 Checklist
					(v6).UnmarshalEasyJSON(in)
					out.Checklists = append(out.Checklists, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "labels":
			if in.IsNull() {
				in.Skip()
				out.Labels = nil
			} else {
				in.Delim('[')
				if out.Labels == nil {
					if !in.IsDelim(']') {
						out.Labels = make([]Label, 0, 1)
					} else {
						out.Labels = []Label{}
					}
				} else {
					out.Labels = (out.Labels)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Label
					(v7).UnmarshalEasyJSON(in)
					out.Labels = append(out.Labels, v7)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Level != 0 {
		const prefix string = ",\"level\":"
		out.RawString(prefix)
		out.Uint(uint(in.Level))
	}
	if in.Deadline != "" {
		const prefix string = ",\"deadline\":"
		out.RawString(prefix)
		out.String(string(in.Deadline))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		out.Float64(float64(in.Pos))
	}
	{
		const prefix string = ",\"cid\":"
		out.RawString(prefix)
		out.Uint(uint(in.Cid))
	}
	if len(in.Members) != 0 {
		const prefix string = ",\"members\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Members {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(out, v9)
			}
			out.RawByte(']')
		}
	}
	if len(in.Comments) != 0 {
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v10, v11 := range in.Comments {
				if v10 > 0 {
					out.RawByte(',')
				}
				(v11).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Checklists) != 0 {
		const prefix string = ",\"checklists\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v12, v13 := range in.Checklists {
				if v12 > 0 {
					out.RawByte(',')
				}
				(v13).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"labels\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v14, v15 := range in.Labels {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Task) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Task) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels1(l, v)
}
func easyjson79a0a577DecodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "password":
			if in.IsNull() {
				in.Skip()
				out.Password = nil
			} else {
				out.Password = in.Bytes()
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
func easyjson79a0a577EncodeGithubComGoParkMailRu20201SIBIRSKAYAKORONAAppModels2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	if len(in.Password) != 0 {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Password)
	}
	out.RawByte('}')
}
