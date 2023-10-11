package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/daveworth/refresher/kprobe"
)

func main() {
	backgroundCtx := context.Background()
	ctx, cancel := context.WithCancel(backgroundCtx)

	done := make(chan os.Signal, 1)
	defer func() {
		fmt.Println("closing original done channel in main()")
		close(done)
	}()

	signal.Notify(done, os.Interrupt, os.Kill)
	go func() {
		for range done {
			fmt.Println("things are done - cancelling channel")
			cancel()
			return
		}
	}()
	kprobe.Trackexecve(ctx)
}
