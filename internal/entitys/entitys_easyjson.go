// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entitys

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

func easyjson8c49de32DecodeAvitotestgo2024InternalEntitys(in *jlexer.Lexer, out *Error) {
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
		case "error":
			out.Message = string(in.String())
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
func easyjson8c49de32EncodeAvitotestgo2024InternalEntitys(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys(l, v)
}
func easyjson8c49de32DecodeAvitotestgo2024InternalEntitys1(in *jlexer.Lexer, out *Content) {
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
		case "title":
			out.Title = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "url":
			out.Url = string(in.String())
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
func easyjson8c49de32EncodeAvitotestgo2024InternalEntitys1(out *jwriter.Writer, in Content) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.Url))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Content) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Content) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Content) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Content) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys1(l, v)
}
func easyjson8c49de32DecodeAvitotestgo2024InternalEntitys2(in *jlexer.Lexer, out *Banners) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Banners, 0, 0)
			} else {
				*out = Banners{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Banner
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
func easyjson8c49de32EncodeAvitotestgo2024InternalEntitys2(out *jwriter.Writer, in Banners) {
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
func (v Banners) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Banners) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Banners) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Banners) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys2(l, v)
}
func easyjson8c49de32DecodeAvitotestgo2024InternalEntitys3(in *jlexer.Lexer, out *Banner) {
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
		case "banner_id":
			out.Id = int(in.Int())
		case "tag_ids":
			if in.IsNull() {
				in.Skip()
				out.Tag_ids = nil
			} else {
				in.Delim('[')
				if out.Tag_ids == nil {
					if !in.IsDelim(']') {
						out.Tag_ids = make([]int, 0, 8)
					} else {
						out.Tag_ids = []int{}
					}
				} else {
					out.Tag_ids = (out.Tag_ids)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int
					v4 = int(in.Int())
					out.Tag_ids = append(out.Tag_ids, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "feature_id":
			out.Feature_ids = int(in.Int())
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
		case "is_active":
			out.Is_active = bool(in.Bool())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created_at).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Updatet_at).UnmarshalJSON(data))
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
func easyjson8c49de32EncodeAvitotestgo2024InternalEntitys3(out *jwriter.Writer, in Banner) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"banner_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix)
		if in.Tag_ids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tag_ids {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"feature_id\":"
		out.RawString(prefix)
		out.Int(int(in.Feature_ids))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.Is_active))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.Created_at).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.Updatet_at).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Banner) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Banner) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Banner) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Banner) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys3(l, v)
}
func easyjson8c49de32DecodeAvitotestgo2024InternalEntitys4(in *jlexer.Lexer, out *Ans201) {
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
		case "banner_id":
			out.Id = int(in.Int())
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
func easyjson8c49de32EncodeAvitotestgo2024InternalEntitys4(out *jwriter.Writer, in Ans201) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"banner_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ans201) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ans201) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8c49de32EncodeAvitotestgo2024InternalEntitys4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ans201) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ans201) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8c49de32DecodeAvitotestgo2024InternalEntitys4(l, v)
}
