/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"example-client/internal/client"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "Deletes a TODO by its Id",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// args
		c, err := client.NewClient(serverFlagURL)
		if err != nil {
			println("Error, could not create new Client", err)
			return
		}
		defer c.Close()
		id, err := convertStringToInt(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		output, err := c.DeleteTask(ctx, id)
		if err != nil {
			fmt.Println("Error Deleting a task", err)
			return
		}
		fmt.Println("Found a Task & Deleted it (Here is the info):")
		fmt.Printf(
			"ID: %d | Title: %s | Description: %s\n",
			output.Id,
			output.Title,
			output.Description,
		)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
