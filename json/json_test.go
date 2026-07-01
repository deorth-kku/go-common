package cjson

import (
	"bytes"
	v1 "encoding/json"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"net"
	"testing"

	citer "github.com/deorth-kku/go-common/iter"
	cmath "github.com/deorth-kku/go-common/math"
)

func TestFloat(t *testing.T) {
	f := cmath.Inf[Float32[ToPosInf]](1)
	data, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
		return
	}
	var f2 Float32[ToPosInf]
	if err := json.Unmarshal(data, &f2); err != nil {
		t.Error(err)
		return
	}
	if !cmath.IsInf(f2, 1) {
		t.Error("not positive infinity")
		return
	}
	if err := json.Unmarshal([]byte("3"), &f2); err != nil {
		t.Error(err)
		return
	}
	if f2 != 3 {
		t.Error("wrong number")
	}
}

func TestSqlString(t *testing.T) {
	var n SqlString[net.IP]
	err := n.Scan("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n.Raw)
}

func TestOmitTopLevelNewline(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	enc := jsontext.NewEncoder(buf, WithOmitTopLevelNewline(true))
	err := enc.WriteValue(jsontext.Value("{}"))
	if err != nil {
		t.Error(err)
		return
	}
	if buf.Bytes()[buf.Len()-1] == 'n' {
		t.Error("the tailing new line was added")
	}
	fmt.Print(buf.String())
}

func BenchmarkToUintptr(b *testing.B) {
	for b.Loop() {
		toUintptr(func() {})
	}
}

func TestToSlice(t *testing.T) {
	data := `["test"]`
	var it SeqJson[string]
	err := json.Unmarshal([]byte(data), &it)
	if err != nil {
		t.Error(err)
		return
	}
	ptr := get(it)
	if ptr == nil {
		t.Error("failed to recover")
		return
	}
	fmt.Println(len(*ptr))
}

func TestToPairSlice(t *testing.T) {
	data := `{"key":"value"}`
	var it Seq2Json[string, string]
	err := json.Unmarshal([]byte(data), &it)
	if err != nil {
		t.Error(err)
		return
	}
	ptr := get2(it)
	if ptr == nil {
		t.Error("failed to recover")
		return
	}
	fmt.Println(len(*ptr))
}

func TestV1Behavior(t *testing.T) {
	data, err := v1.Marshal(ToSeq2(citer.EmptyRange2[string, struct{}]))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}
