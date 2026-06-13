/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"example-client/internal/client"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all TODO items inside the list",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient(serverFlagURL)
		if err != nil {
			fmt.Println("Error creating client")
			return
		}
		output, err := c.ListTasks(cmd.Context())
		defer c.Close()
		if err != nil {
			fmt.Println("Error fetching tasks")
			return
		}
		if output == nil || output.Tasks == nil {
			fmt.Println("No tasks found")
			return
		}
		for _, task := range output.Tasks {
			fmt.Printf("Id: %d | Title: %s | Description: %s\n",
				task.Id, task.Title, task.Description)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
