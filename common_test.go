package common

import (
	"cmp"
	"errors"
	"fmt"
	"maps"
	"math"
	"math/rand/v2"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {
	server := NewHttpServer()
	server.ListenAndServe("/tmp/123.sock,0666")
}

func TestAbstrace(t *testing.T) {
	server := NewHttpServer()
	server.ListenAndServe("@123.sock")
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
	f := NaN32()
	if !IsNaN(f) {
		t.Error("no!")
	}
}

func TestInf32(t *testing.T) {
	f := Inf32(1)
	if !IsInf(f, 1) {
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

func TestRand(t *testing.T) {
	s := []string{
		"a",
		"b",
		"c",
	}

	for i, a := range SliceRandom(s) {
		fmt.Println(i, a)
	}
	fmt.Println(s)
	SliceShuffle(s)
	fmt.Println(s)
}

func BenchmarkInf(b *testing.B) {
	num := math.Inf(1)
	for range b.N {
		IsInf(num, 1)
	}
}

func BenchmarkInfStd(b *testing.B) {
	num := math.Inf(1)
	for range b.N {
		math.IsInf(num, 1)
	}
}

func BenchmarkInf32(b *testing.B) {
	num := Inf32(1)
	for range b.N {
		IsInf(num, 1)
	}
}

func BenchmarkInf32Std(b *testing.B) {
	num := Inf32(1)
	for range b.N {
		math.IsInf(float64(num), 1)
	}
}

type interror int

func (i interror) Error() string {
	return strconv.Itoa(int(i))
}

func TestUnwraps(t *testing.T) {
	err := errors.New("start")
	for i := range interror(10) {
		err = fmt.Errorf("this is %w, %w", err, error(i))
	}
	for i := range Unwraps(err) {
		fmt.Println(i.Error())
	}
	fmt.Println(err.Error())
}

func TestMod(t *testing.T) {
	var b int
	fmt.Println(1 / b)
}

func BenchmarkCeil(b *testing.B) {
	for range b.N {
		DevidedCeil(rand.Int(), rand.Int())
	}
}

func BenchmarkReflectIsZero(b *testing.B) {
	var t time.Time
	for range b.N {
		reflect.ValueOf(t).IsZero()
	}
}

func BenchmarkGenericsIsZero(b *testing.B) {
	var t time.Time
	for range b.N {
		IsZero(t)
	}
}

func TestSeq2(t *testing.T) {
	keys := slices.Collect(Seq2K(maps.All(map[string]int{
		"1": 2,
		"3": 4,
	})))
	fmt.Println(keys)
}

var ips = []string{
	"192.168.1.1",
	"10.0.0.0",
	"172.16.254.1",
	"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	"::1",
	"fe80::1ff:fe23:4567:890a",
	"2001:db8::",
	"::ffff:192.0.2.128",
}

func getip(count int) string {
	return ips[count%len(ips)]
}

func BenchmarkOldParseIP(b *testing.B) {
	for i := range b.N {
		ShortIP(net.ParseIP(getip(i)))
	}
}

func BenchmarkParseIP(b *testing.B) {
	for i := range b.N {
		ParseIP(getip(i))
	}
}

func TestParseIP(t *testing.T) {
	for i, ip := range ips {
		parsed := ParseIP(ip)
		if parsed == nil {
			t.Errorf("failed to parse IP %s at index %d", ip, i)
		}
	}
}

func TestBSMap(t *testing.T) {
	m := NewBSMap[int, Empty]()
	for range 100 {
		m.Store(rand.Int(), Empty{})
	}
	if !slices.IsSorted(m.SliceKeys()) {
		t.Fatal("not sorted")
	}
	for key := range m.Keys {
		_, ok := m.Load(key)
		if !ok {
			t.Fatal("not found")
		}
	}
	for _, key := range SliceRandom(m.SliceKeys()) {
		m.Delete(key)
		_, ok := m.Load(key)
		if ok {
			t.Fatal("found after delete")
		}
	}
}

type fakeint int

func (a fakeint) Compare(b fakeint) int {
	return cmp.Compare(a, b)
}

func genfakeint() fakeint {
	return rand.N[fakeint](math.MaxInt)
}

func TestBSMapT(t *testing.T) {
	m := NewBSMapT[fakeint, Empty]()
	for range 100 {
		m.Store(genfakeint(), Empty{})
	}
	if !slices.IsSorted(m.SliceKeys()) {
		t.Fatal("not sorted")
	}
	for key := range m.Keys {
		_, ok := m.Load(key)
		if !ok {
			t.Fatal("not found")
		}
	}
	for _, key := range SliceRandom(m.SliceKeys()) {
		m.Delete(key)
		_, ok := m.Load(key)
		if ok {
			t.Fatal("found after delete")
		}
	}
}

func BenchmarkBSMap(b *testing.B) {
	m := NewBSMap(make(PairSlice[fakeint, Empty], 0, b.N)...)
	for range b.N {
		m.Store(genfakeint(), Empty{})
	}
}

func BenchmarkBSMapT(b *testing.B) {
	m := NewBSMapT(make(PairSlice[fakeint, Empty], 0, b.N)...)
	for range b.N {
		m.Store(genfakeint(), Empty{})
	}
}

func BenchmarkBuiltinMap(b *testing.B) {
	m := make(map[fakeint]Empty, b.N)
	for range b.N {
		m[genfakeint()] = Empty{}
	}
}

func BenchmarkParse(b *testing.B) {
	type ii int
	str := "-1"
	for range b.N {
		_, err := Parse[ii](str, 10)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkIsZero(b *testing.B) {
	var a time.Time
	for range b.N {
		IsZero(a)
	}
}

func BenchmarkIsZeroSlow(b *testing.B) {
	a := time.Now()
	for range b.N {
		IsZeroSlow(a)
	}
}

func TestClear(t *testing.T) {
	m := NewBSMap[int, string]()
	m.Store(1, "1")
	m.Clear()
	fmt.Print(m)
}

func TestSig(t *testing.T) {
	var count atomic.Int32
	stop := SignalsCallback(func() {
		count.Add(1)
	}, false, syscall.SIGALRM)
	defer stop()
	const tries = 3
	for range tries {
		syscall.Kill(os.Getpid(), syscall.SIGALRM)
		time.Sleep(time.Millisecond)
	}

	if count.Load() != tries {
		t.Error(fmt.Errorf("not enough tries, expected: %d, actual: %d", tries, count.Load()))
	}
}
