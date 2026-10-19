package cjson

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"runtime"

	"github.com/deorth-kku/go-common/cleanup"
)

// JsonWriter is the reverse of JsonReader.
// Bytes written to it are streamed into a decoder that unmarshals into the target.
type JsonWriter struct {
	out     any
	options json.Options
	wt      *io.PipeWriter
	done    chan error
}

// NewJsonWriter creates a new JsonWriter that streams written bytes into out.
func NewJsonWriter(out any, opts ...json.Options) *JsonWriter {
	return &JsonWriter{
		out:     out,
		options: json.JoinOptions(opts...),
	}
}

// Write implements io.Writer. Bytes are streamed to the decoder in real time.
func (j *JsonWriter) Write(p []byte) (int, error) {
	if j.wt == nil {
		var rd *io.PipeReader
		rd, j.wt = io.Pipe()
		j.done = make(chan error, 1)
		go func() {
			dec := jsontext.NewDecoder(rd, j.options)
			err := json.UnmarshalDecode(dec, j.out)
			rd.Close()
			j.done <- err
		}()
		runtime.AddCleanup(j, cleanup.Closer, j.wt)
	}
	return j.wt.Write(p)
}

// Close signals EOF to the decoder and returns the unmarshal error.
func (j *JsonWriter) Close() error {
	if j.wt == nil {
		return nil
	}
	j.wt.Close()
	return <-j.done
}
