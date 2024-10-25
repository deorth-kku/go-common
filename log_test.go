package common

import (
	"log/slog"
	"maps"
	"math/rand/v2"
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
	err := SetLog("", "DEBUG", "JSON", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}

	slog.Debug("test")
}
func TestSetGroup(t *testing.T) {
	err := SetLog("", "DEBUG", "DEFAULT", SlogHideTime{})
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
	err := SetLog("", "DEBUG", "DEFAULT", SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
	}

	slog.Info("test log iter and map", "iter", maps.All(m), "map", m)
}

func TestSlogStruct(t *testing.T) {
	a := mix{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	err := SetLog("", "DEBUG", "DEFAULT", SlogStruct[mix]{}, SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test struct", "a", a)
}

func TestAddSourceFunc(t *testing.T) {
	err := SetLog("", "DEBUG", "", SlogHideTime{}, SlogAddSourceFunc{func() bool { return rand.UintN(2) == 1 }})
	if err != nil {
		t.Error(err)
		return
	}
	for range 100 {
		slog.Info("test AddSourceFunc", "a", "1")
	}
}

func TestQuoteAttr(t *testing.T) {
	err := SetLog("", "DEBUG", "", SlogQuoteAttr{slog.LevelKey, "[", "]"}, SlogQuoteAttr{slog.SourceKey, "(", ")"}, SlogAddSource{})
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
	err := SetLog("", "DEBUG", "", SlogSlice[any]{}, SlogStruct[mix]{})
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test slice", "slice", sli)
	err = SetLog("", "DEBUG", "JSON")
	if err != nil {
		t.Error(err)
		return
	}
	slog.Info("test slice", "slice", sli)
}
