package clog

import (
	"log/slog"
	"maps"
	"math/rand/v2"
	"os"
	"testing"
)

var (
	_ SlogOption = SlogAddSource{}
	_ SlogOption = SlogHideTime{}
	_ SlogOption = SlogIter{}
	_ SlogOption = SlogMap{}
	_ SlogOption = SlogStruct[any]{}
	_ SlogOption = SlogAddSourceFunc{}
	_ SlogOption = SlogQuoteAttr{}
	_ SlogOption = SlogSlice[any]{}
)

func TestSetLog(t *testing.T) {
	err := Set("", "DEBUG", "JSON", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}

	slog.Debug("test")
}
func TestSetGroup(t *testing.T) {
	err := Set("", "DEBUG", "DEFAULT", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}
	slog.Info("test", "test", Map2Group(map[string]any{
		"a": 1,
		"b": 2,
		"g": map[string]any{
			"a": 1,
			"b": 2,
		},
	}))

}

func TestSlogIterMap(t *testing.T) {
	m := map[string]any{
		"a": 1,
		"b": 2,
	}
	err := Set("", "DEBUG", "DEFAULT", SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
	}

	slog.Info("test log iter and map", "iter", maps.All(m), "map", m)
}

type mix struct {
	A int
	B string
	M map[string]any
}

func TestSlogStruct(t *testing.T) {
	a := mix{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	err := Set("", "DEBUG", "DEFAULT", SlogStruct[mix]{}, SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test struct", "a", a)
}

func TestAddSourceFunc(t *testing.T) {
	err := Set("", "DEBUG", "", SlogHideTime{}, SlogAddSourceFunc{func() bool { return rand.UintN(2) == 1 }})
	if err != nil {
		t.Error(err)
		return
	}
	for range 100 {
		slog.Info("test AddSourceFunc", "a", "1")
	}
}

func TestQuoteAttr(t *testing.T) {
	err := Set("", "DEBUG", "", SlogQuoteAttr{slog.LevelKey, "[", "]"}, SlogQuoteAttr{slog.SourceKey, "(", ")"}, SlogAddSource{})
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test quote attr")
}

func TestSlice(t *testing.T) {
	a := mix{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	sli := []any{"a", a, 1}
	err := Set("", "DEBUG", "", SlogSlice[any]{}, SlogStruct[mix]{})
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test slice", "slice", sli)
	err = Set("", "DEBUG", "JSON")
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test slice", "slice", sli)
}

type testvaluer struct{}

func (testvaluer) LogValue() slog.Value {
	return slog.StringValue("test")
}

var _ slog.LogValuer = testvaluer{}

func TestLogValuer(t *testing.T) {
	SetRaw(os.Stderr, slog.LevelDebug, DefaultFormat)
	slog.Debug("test valuer", "test", testvaluer{})
}
