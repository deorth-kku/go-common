package cjson

import (
	"bytes"
	v1 "encoding/json"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"net"
	"strings"
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

func TestReadSeq(t *testing.T) {
	data := strings.NewReader("[1,2,3]")
	it, errf := ReaderSeq[int](data)
	for line := range it {
		fmt.Println(line)
	}
	if errf() != nil {
		t.Error(errf())
		return
	}
}

func TestReadSeq2(t *testing.T) {
	data := strings.NewReader(`{"test":1,"hello":2}`)
	it, errf := ReaderSeq2[string, int](data)
	for k, v := range it {
		fmt.Println(k, v)
	}
	if errf() != nil {
		t.Error(errf())
		return
	}
}

func TestReadSeq2Err(t *testing.T) {
	data := strings.NewReader(`{"test":1,"hello":2}`)
	it, errf := ReaderSeq2[string, int](data)
	for k, v := range it {
		fmt.Println(k, v)
		break
	}
	if errf() == nil {
		t.Error("no error found")
		return
	}
	fmt.Println(errf())
}

func TestConsumer2(t *testing.T) {
	data := strings.NewReader(`{"test":1,"hello":2}`)
	var co Consumer2[string, int]
	co = func(s string, i int) bool {
		fmt.Println(s, i)
		return true
	}
	err := json.UnmarshalRead(data, &co)
	if err != nil {
		t.Error(err)
		return
	}
}
