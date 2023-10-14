package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGTERM, os.Interrupt)

	go func() {
		<-q
		cancel()
	}()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) (err error) {
	fmt.Fprintln(os.Stdout, "Hello, World")
	return
}
