// ls ~/.local/share/east-story/
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	gap "github.com/muesli/go-app-paths"
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
	// use XDG to create necessary data dirs for the program
	s := gap.NewScope(gap.User, "east-story")
	dirs, err := s.DataDirs()
	if err != nil {
		return err
	}
	// create the app base dir, if it doesn't exist
	var dir string
	if len(dirs) > 0 {
		dir = dirs[0]
	} else {
		dir, _ = os.UserHomeDir()
	}
	// create directory with 0o770 (504) permission
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dir, 0o770)
		}
		return err
	}
	return
}
