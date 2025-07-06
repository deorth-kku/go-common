package common

import (
	"encoding/json"
	"fmt"
	"net"
	"testing"
)

func TestFloat(t *testing.T) {
	f := Inf[JsonFloat32[ToPosInf]](1)
	data, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
		return
	}
	var f2 JsonFloat32[ToPosInf]
	if err := json.Unmarshal(data, &f2); err != nil {
		t.Error(err)
		return
	}
	if !IsInf(f2, 1) {
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
