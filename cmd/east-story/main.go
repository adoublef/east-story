// ls ~/.local/share/east-story/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	gap "github.com/muesli/go-app-paths"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "east-story",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var echoCmd = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}

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
	// init cobra
	return rootCmd.Execute()
}
