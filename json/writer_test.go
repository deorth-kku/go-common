package cjson

import (
	json "encoding/json/v2"
	"io"
	"strings"
	"testing"
)

func TestJsonWriterBasic(t *testing.T) {
	var out map[string]any
	w := NewJsonWriter(&out, json.DefaultOptionsV2())

	input := `{"a":"b","c":10}`
	n, err := w.Write([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if n != len(input) {
		t.Fatalf("wrote %d, want %d", n, len(input))
	}

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	if out["a"] != "b" {
		t.Errorf("out[a] = %v, want \"b\"", out["a"])
	}
	if out["c"] != float64(10) {
		t.Errorf("out[c] = %v, want 10", out["c"])
	}
}

func TestJsonWriterChunked(t *testing.T) {
	var out map[string]any
	w := NewJsonWriter(&out, json.DefaultOptionsV2())

	// Write in chunks to verify streaming
	chunks := []string{`{"a":"`, `b","c":`, `10}`}
	for _, chunk := range chunks {
		if _, err := w.Write([]byte(chunk)); err != nil {
			t.Fatal(err)
		}
	}

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	if out["a"] != "b" {
		t.Errorf("out[a] = %v, want \"b\"", out["a"])
	}
	if out["c"] != float64(10) {
		t.Errorf("out[c] = %v, want 10", out["c"])
	}
}

func TestJsonWriterCopy(t *testing.T) {
	var out struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	w := NewJsonWriter(&out, json.DefaultOptionsV2())

	src := strings.NewReader(`{"name":"testuser","age":42}`)
	if _, err := io.Copy(w, src); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	if out.Name != "testuser" {
		t.Errorf("Name = %q, want \"testuser\"", out.Name)
	}
	if out.Age != 42 {
		t.Errorf("Age = %d, want 42", out.Age)
	}
}

func TestJsonWriterInvalidJSON(t *testing.T) {
	var out map[string]any
	w := NewJsonWriter(&out, json.DefaultOptionsV2())

	_, _ = w.Write([]byte(`{"a":`))
	err := w.Close()
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestJsonWriterCloseWithoutWrite(t *testing.T) {
	var out map[string]any
	w := NewJsonWriter(&out, json.DefaultOptionsV2())

	// Close without any write should be a no-op
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
}
