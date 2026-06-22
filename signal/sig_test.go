package csignal

import (
	"os"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestSig(t *testing.T) {
	var count atomic.Int32
	stop := Callback(func() {
		count.Add(1)
	}, false, syscall.SIGALRM)
	defer stop()
	const tries = 3
	for range tries {
		syscall.Kill(os.Getpid(), syscall.SIGALRM)
		time.Sleep(time.Millisecond)
	}

	if count.Load() != tries {
		t.Errorf("not enough tries, expected: %d, actual: %d", tries, count.Load())
	}
}
