package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to authenticate",
	Long:  `echo is for authenticating the user`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		rp := oauth2.Config{
			// NOTE -- will use infisical to access these
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Endpoint:     github.Endpoint,
		}

		rs, err := rp.DeviceAuth(context.Background())
		if err != nil {
			return fmt.Errorf("failed to complete device flow: %w", err)
		}
		fmt.Printf("\nPlease browse to %s and enter code %s\n", rs.VerificationURI, rs.UserCode)

		token, err := rp.DeviceAccessToken(context.Background(), rs)
		if err != nil {
			return err
		}

		log.Printf("successfully obtained token: %v", token)
		// login to server with auth token
		return nil
	},
}
