package common

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestHttp(t *testing.T) {
	server := NewHttpServer()
	server.ListenAndServe("/tmp/123.sock,0666")
}

func TestCutSlice(t *testing.T) {
	longslice := make([]int, 65535)
	for i := range longslice {
		longslice[i] = i
	}
	last := -1
	for _, subslice := range CutSlice(longslice, 100) {
		if subslice[0] != last+1 {
			t.Error("no!")
		}
		last = subslice[len(subslice)-1]
	}
}

func TestNaN32(t *testing.T) {
	f := Nan32()
	if !IsNaN(f) {
		t.Error("no!")
	}
}

type mix struct {
	A int
	B string
	M map[string]any
}

func TestStruct(t *testing.T) {
	a := mix{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	m, err := Struct2Map(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}

func TestParseMode(t *testing.T) {
	str, m, _, err := FileWithMode("te\\,st,0666")
	if err != nil {
		t.Error(err)
		return
	}
	if m != 0666 {
		t.Error("wrong")
	}
	fmt.Println(str)
}

func TestCheckDirPerm(t *testing.T) {
	dirname := "/tmp/test.12313"
	err := CheckDirWritePermission(dirname)
	if err == nil {
		t.Error("not exist but no error")
	} else {
		fmt.Println("expected no exist error:", err)
	}
	err = os.Mkdir(dirname, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(dirname)
	err = CheckDirWritePermission(dirname)
	if err == nil {
		t.Error("not permissoned but no error")
	} else {
		fmt.Println("expected no permission error:", err)
	}
	err = os.Chmod(dirname, 0755)
	if err != nil {
		t.Error(err)
		return
	}
	err = CheckDirWritePermission(dirname)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckFilePerm(t *testing.T) {
	filename := "/tmp/abc/123"
	err := CheckFileWritePermission(filename)
	if err == nil {
		t.Error("parent not exist but no error")
	} else {
		fmt.Println("expect parent not exist error:", err)
	}

	dirname := filepath.Dir(filename)
	err = os.Mkdir(dirname, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(dirname)
	err = CheckFileWritePermission(filename)
	if err == nil {
		t.Error("parent no permisson but no error")
	} else {
		fmt.Println("expected parent no permisson error:", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	f.Close()
	defer os.Remove(filename)

	err = CheckFileWritePermission(filename)
	if err != nil {
		t.Error(err)
		return
	}

	err = os.Chmod(filename, 0000)
	if err != nil {
		t.Error(err)
		return
	}
	err = CheckFileWritePermission(filename)
	if err == nil {
		t.Error("no permisson but no error")
	} else {
		fmt.Println("expected no permisson error:", err)
	}

}

func TestParse(t *testing.T) {
	type ii int
	str := "-1"
	i, err := Parse[ii](str, 10)
	if err != nil {
		t.Error(err)
		return
	}
	if str != fmt.Sprint(i) {
		t.Error("not match")
		return
	}

	type uu uint16
	str = "443"
	u8, err := Parse[uu](str, 10)
	if err != nil {
		t.Error(err)
		return
	}
	if str != fmt.Sprint(u8) {
		t.Error("not match")
		return
	}
	_, err = Parse[uu]("65536", 10)
	if err == nil {
		t.Error("not overflow when it should")
		return
	}

	str = "1.2345"
	fstr := "%.04f"
	f, err := Parse[float32](str, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if str != fmt.Sprintf(fstr, f) {
		t.Error("not match")
		fmt.Printf("%f", f)
		return
	}

	f64, err := Parse[float64](str, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if str != fmt.Sprintf(fstr, f64) {
		fmt.Println(f64)
		t.Error("not match")
		return
	}

}

func TestMaxInt(t *testing.T) {
	type uu uint
	if MaxInt[uu]() != math.MaxUint {
		t.Error("not eq")
	}
	type ii int
	if MaxInt[ii]() != math.MaxInt {
		t.Error("not eq")
	}
}
