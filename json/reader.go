package cjson

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"runtime"
	"sync"

	"github.com/deorth-kku/go-common/cleanup"
)

type JsonReader struct {
	data    any
	options json.Options
	rd      *io.PipeReader
}

func NewJsonReader(data any, opts ...json.Options) *JsonReader {
	return &JsonReader{
		data:    data,
		options: json.JoinOptions(opts...),
	}
}

func (j *JsonReader) Read(b []byte) (int, error) {
	if j.rd == nil {
		var wt *io.PipeWriter
		j.rd, wt = io.Pipe()
		go func() {
			var err error
			defer func() { wt.CloseWithError(err) }()
			_, err = j.writeto(wt)
		}()
		runtime.AddCleanup(j, cleanup.Closer, wt)
	}
	return j.rd.Read(b)
}

func (j *JsonReader) Close() error {
	if j.rd == nil {
		j.setclose()
		return nil
	}
	return j.rd.Close()
}

func (j JsonReader) writeto(wt io.Writer) (int64, error) {
	enc := jsontext.NewEncoder(wt, j.options)
	err := json.MarshalEncode(enc, j.data)
	return enc.OutputOffset(), err
}

var (
	closedPipe *io.PipeReader
	pipeOnce   sync.Once
)

func (j *JsonReader) setclose() {
	pipeOnce.Do(func() {
		var wt *io.PipeWriter
		closedPipe, wt = io.Pipe()
		wt.Close()
	})
	j.rd = closedPipe
}

func (j *JsonReader) WriteTo(wt io.Writer) (int64, error) {
	if j.rd != nil {
		return io.Copy(wt, j.rd)
	}
	j.setclose()
	return j.writeto(wt)
}

type countWriter struct {
	n int
}

func (c *countWriter) Write(p []byte) (int, error) {
	c.n += len(p)
	return len(p), nil
}

func (j JsonReader) Len() (int64, error) {
	wt := new(countWriter)
	err := json.MarshalWrite(wt, j.data, j.options)
	// [json.MarshalWrite] does not write the tailing \n
	// this is faster than using [jsontext.Encoder.OutputOffset]
	return int64(wt.n + 1), err
}

type ReaderJson struct {
	rd io.Reader
}

func (rj ReaderJson) MarshalJSON() ([]byte, error) {
	return MarshalV1(rj)
}

func (rj ReaderJson) MarshalJSONTo(enc *jsontext.Encoder) error {
	dec := jsontext.NewDecoder(rj.rd, enc.Options())
	var err error
	var token jsontext.Token
	var value jsontext.Value
	for err == nil {
		switch dec.PeekKind() {
		case '[', ']', '{', '}':
			token, err = dec.ReadToken()
			if err != nil {
				break
			}
			err = enc.WriteToken(token)
		default:
			value, err = dec.ReadValue()
			if err != nil {
				break
			}
			err = enc.WriteValue(value)
		}
	}
	if err == io.EOF && dec.StackDepth() == 0 {
		err = nil
	}
	return err
}

func NewReaderJson(rd io.Reader) ReaderJson {
	return ReaderJson{rd}
}

func FormatReader(rd io.Reader, opts ...json.Options) *JsonReader {
	return NewJsonReader(NewReaderJson(rd), opts...)
}
