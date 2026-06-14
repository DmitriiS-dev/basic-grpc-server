/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"example-client/internal/client"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all TODO items inside the list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CONNECTING TO:", serverFlagURL)

		c, err := client.NewClient(serverFlagURL)
		if err != nil {
			fmt.Println("Error creating client")
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		output, err := c.ListTasks(ctx)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}

		if len(output.Tasks) == 0 {
			fmt.Println("Your TODO list is empty!")
			return
		}

		fmt.Println("--- Current Tasks ---")
		for _, task := range output.Tasks {
			fmt.Printf("ID: %d | Title: %s | Description: %s\n", task.Id, task.Title, task.Description)
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
