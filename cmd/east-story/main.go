// ls ~/.local/share/east-story/
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const AppName = "east-story"

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
	// TODO -- setup could be moved back here
	return rootCmd.Execute()
}
