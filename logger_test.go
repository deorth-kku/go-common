package common

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"testing"
)

func TestConcurrentLog(t *testing.T) {
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	lg := NewLogger(file, new(slog.LevelVar), TextFormat)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := range 1000 {
			lg.Info("this is thread 1", "i", i)
		}
		wg.Done()
	}()
	go func() {
		cl := lg.Clone()
		for i := range 500 {
			cl.Info("this is thread 2", "i", i)
		}
		cl.Reload(os.Stderr, new(slog.LevelVar), TextFormat)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("now check file %s to see if log is properly written\n", file.Name())
}
