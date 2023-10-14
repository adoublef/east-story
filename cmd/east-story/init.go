package main

import (
	"os"

	gap "github.com/muesli/go-app-paths"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize tooling",
	Long:  `Initialize tooling by setting up location of app data`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// use XDG to create necessary data dirs for the program
		s := gap.NewScope(gap.User, AppName)
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
		return nil
	},
}
