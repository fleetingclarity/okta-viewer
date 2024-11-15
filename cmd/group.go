package cmd

import (
	"context"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/spf13/cobra"
	"os"
)

var groupCmd = &cobra.Command{
	Use:   "groups [user ID]",
	Short: "List groups for a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, client, err := okta.NewClient(context.TODO(), okta.WithOrgUrl(os.Getenv("OV_OKTA_ORG_URL")), okta.WithToken(os.Getenv("OV_OKTA_API_TOKEN")))
		if err != nil {
			fmt.Println("Failed to create the Okta client:", err)
			return
		}
		groups, _, err := client.User.ListUserGroups(context.Background(), args[0])
		if err != nil {
			fmt.Println("Failed to retrieve groups:", err)
			return
		}
		fmt.Println("Groups:")
		for _, group := range groups {
			fmt.Printf("- %s (%s)\n", group.Profile.Name, group.Id)
		}
	},
}

var groupUsersCmd = &cobra.Command{
	Use:   "group-users [group name]",
	Short: "List all users in a specific group by group name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, client, err := okta.NewClient(context.TODO(), okta.WithOrgUrl(os.Getenv("OV_OKTA_ORG_URL")), okta.WithToken(os.Getenv("OV_OKTA_API_TOKEN")))
		if err != nil {
			fmt.Println("Failed to create the Okta client:", err)
			return
		}
		groupName := args[0]
		groups, _, err := client.Group.ListGroups(context.TODO(), &query.Params{Q: groupName})
		if err != nil || len(groups) == 0 {
			fmt.Printf("Group with name %q not found.\n", groupName)
			return
		}

		groupID := groups[0].Id
		allUsers := make([]*okta.User, 0)
		users, resp, err := client.Group.ListGroupUsers(context.TODO(), groupID, nil)
		if err != nil {
			fmt.Println("Failed to retrieve group members: ", err)
			return
		}
		allUsers = append(allUsers, users...)

		for resp.HasNextPage() {
			nextPage := []*okta.User{}
			resp, err = resp.Next(context.TODO(), &nextPage)
			if err != nil {
				fmt.Println("Error retrieving next page of results:", err)
				return
			}
			allUsers = append(allUsers, nextPage...)
		}

		fmt.Printf("Users in group %q:\n", groupName)
		for _, user := range users {
			email, ok := (*user.Profile)["email"]
			if !ok {
				fmt.Println("- Email field is missing or not a string for userId=%q", user.Id)
				continue
			}
			fmt.Printf("- %s (ID: %s, Status: %s)\n", email, user.Id, user.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	rootCmd.AddCommand(groupUsersCmd)
}
