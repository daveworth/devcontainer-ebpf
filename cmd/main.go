package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/daveworth/devcontainer-ebpf/kprobe"
)

func main() {
	backgroundCtx := context.Background()
	ctx, cancel := context.WithCancel(backgroundCtx)

	done := make(chan os.Signal, 1)
	defer func() {
		close(done)
	}()

	signal.Notify(done, os.Interrupt, os.Kill)
	go func() {
		for range done {
			cancel()
			return
		}
	}()
	kprobe.Trackexecve(ctx)
}
