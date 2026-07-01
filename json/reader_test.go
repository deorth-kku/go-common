package cjson

import (
	v1 "encoding/json"
	"encoding/json/jsontext"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

func getreq() map[string]any {
	list := make([]string, 1)
	return map[string]any{
		"version": "2.0",
		"id":      1,
		"name":    "test",
		"list":    list,
	}
}

func BenchmarkDiscard(b *testing.B) {
	rd := NewJsonReader(getreq())
	b.ResetTimer()
	for b.Loop() {
		_, _ = rd.WriteTo(io.Discard)
	}
}

func BenchmarkLen(b *testing.B) {
	rd := NewJsonReader(getreq())
	b.ResetTimer()
	for b.Loop() {
		_, _ = rd.Len()
	}
}

func TestRead(t *testing.T) {
	rd := NewJsonReader(getreq())
	discard := make([]byte, 10)
	var err error
	count := 0
	for {
		var l int
		l, err = rd.Read(discard)
		if err != nil {
			break
		}
		count += l
	}
	fmt.Println(count, err)
	fmt.Println(rd.Len())
}

func TestJsonReaderClose(t *testing.T) {
	rd := NewJsonReader(getreq())
	io.ReadAll(rd)
	n, _ := io.ReadAll(rd)
	if len(n) != 0 {
		t.Error("still read")
	}

	rd = NewJsonReader(getreq())
	io.Copy(io.Discard, rd)
	l, _ := io.Copy(io.Discard, rd)
	if l != 0 {
		t.Error("still read")
	}

	rd = NewJsonReader(getreq())
	io.ReadAll(rd)
	l, _ = io.Copy(io.Discard, rd)
	if l != 0 {
		t.Error("still read")
	}

	rd = NewJsonReader(getreq())
	io.Copy(io.Discard, rd)
	n, _ = io.ReadAll(rd)
	if len(n) != 0 {
		t.Error("still read")
	}

	runtime.GC()
}

func TestFormatReader(t *testing.T) {
	data := strings.NewReader(`{"test":{"a":[1]}}`)
	rd := FormatReader(data, v1.DefaultOptionsV1(), WithOmitTopLevelNewline(true), jsontext.WithIndent("  "))
	_, err := io.Copy(os.Stdout, rd)
	if err != nil {
		t.Error(err)
	}
}

func TestJsonReader(t *testing.T) {
	file := map[string]any{
		"a": "b",
		"c": 10,
		"d": struct {
			Test string
		}{
			Test: "1",
		},
	}
	jr := NewJsonReader(file, v1.DefaultOptionsV1())
	n, err := io.Copy(
		os.Stdout,
		jr,
	)
	if err != nil {
		t.Error(err)
		return
	}

	jr = NewJsonReader(file, v1.DefaultOptionsV1())
	data, err := io.ReadAll(jr)
	if err != nil {
		t.Error(err)
		return
	}
	if int(n) != len(data) {
		t.Error("data not equal len")
	}
}
