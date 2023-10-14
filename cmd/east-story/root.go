package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:  AppName,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(echoCmd, initCmd, loginCmd)
}
