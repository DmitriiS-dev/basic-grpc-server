/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"example-client/internal/client"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a TODO item by its ID",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

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

		output, err := c.GetTask(ctx, id)
		if err != nil {
			fmt.Println("Error fetching task", err)
		}
		if output.Id == 0 {
			fmt.Println("No task found with id %d", id)
			return
		}
		fmt.Println("Found a Task:")
		fmt.Printf("Title: %d | Description %s \n", output.Title, output.Description)
	},
}

func convertStringToInt(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Could not convert %s to int: %w", s, err)
	}
	return int32(i), nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
