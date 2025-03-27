package cmd

import (
	"context"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/spf13/cobra"
	"os"
)

var userCmd = &cobra.Command{
	Use:   "user [email or id]",
	Short: "Retrieve user details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, client, err := okta.NewClient(context.TODO(), okta.WithOrgUrl(os.Getenv("OV_OKTA_ORG_URL")), okta.WithToken(os.Getenv("OV_OKTA_API_TOKEN")))
		if err != nil {
			fmt.Println("Failed to create Okta client:", err)
			return
		}
		user, _, err := client.User.GetUser(context.Background(), args[0])
		if err != nil {
			fmt.Println("Failed to retrieve user:", err)
			return
		}

		fmt.Printf(`User Details for %s:
	ID: %s
	Created: %s
	Activated: %s
	LastLogin: %s
	LastPasswordChange: %s
	LastUpdated: %s
	Login: %s
	Status: %s
	StatusChanged: %s
`,
			args[0], user.Id, user.Created, user.Activated, user.LastLogin,
			user.LastUpdated, user.PasswordChanged, (*user.Profile)["login"], user.Status, user.StatusChanged)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
