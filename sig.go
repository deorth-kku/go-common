package common

import (
	"context"
	"os"
	"os/signal"
)

type stopfunc = func() bool

// SignalsCallback registers a callback to be invoked when specified signals are received.
// If once is true, the callback is invoked at most once; otherwise it is invoked each time
// a signal is received. The returned function stops listening for signals and returns true
// if the context was successfully stopped, or false if it was already stopped.
func SignalsCallback(cb func(), once bool, sigs ...os.Signal) stopfunc {
	if once {
		return signalsCallbackOnce(cb, sigs...)
	}
	return signalsCallback(cb, sigs...)
}

func signalsCallback(cb func(), sigs ...os.Signal) stopfunc {
	var cancel context.CancelFunc
	var stop func() bool

	var register func()
	register = func() {
		var ctx context.Context
		ctx, cancel = signal.NotifyContext(context.Background(), sigs...)
		stop = context.AfterFunc(ctx, func() {
			register()
			cb()
		})
	}

	register()

	return func() (stopped bool) {
		stopped = stop()
		cancel()
		return
	}
}

func signalsCallbackOnce(cb func(), sigs ...os.Signal) stopfunc {
	ctx, cancel := signal.NotifyContext(context.Background(), sigs...)
	stop := context.AfterFunc(ctx, cb)
	return func() (stopped bool) {
		stopped = stop()
		cancel()
		return
	}
}
