package cjson

import (
	"bytes"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"net"
	"testing"

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
