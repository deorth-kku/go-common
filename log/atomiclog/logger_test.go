package atomiclog

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
	lg := GetLogger(file, slog.LevelInfo, slog.NewTextHandler)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := range 1000 {
			lg.Info("this is thread 1", "i", i)
		}
		wg.Done()
	}()
	go func() {
		defer wg.Done()
		cl, err := Clone(lg)
		if err != nil {
			t.Error(err)
			return
		}
		for i := range 500 {
			cl.Info("this is thread 2", "i", i)
		}
		err = Reload(cl, os.Stderr, slog.LevelInfo, slog.NewTextHandler)
		if err != nil {
			t.Error(err)
		}
	}()
	wg.Wait()
	fmt.Printf("now check file %s to see if log is properly written. it should be all in file, and no output on stderr\n", file.Name())
}

func TestReload(t *testing.T) {
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	lg := GetLogger(file, slog.LevelInfo, slog.NewTextHandler)
	for i := range 100 {
		lg.Info("logging", "i", i)
		if i == 50 {
			err = Reload(lg, os.Stderr, slog.LevelInfo, slog.NewTextHandler)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}
	fmt.Printf("now check file %s to see if log is properly written. half on file and half on stderr\n", file.Name())
}
