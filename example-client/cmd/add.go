/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"example-client/internal/client"
	"fmt"
	"time"

	pb "github.com/DmitriiS-dev/basic-backend/proto"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title] [description]",
	Short: "Adds a new TODO item to the list",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewClient(serverFlagURL)
		if err != nil {
			fmt.Println("Error creating client")
			return
		}
		defer c.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		task := &pb.Task{Title: args[0], Description: args[1]}
		output, err := c.AddTask(ctx, task)
		if err != nil {
			fmt.Println("Error Adding A Task", err)
			return
		}
		fmt.Println("Created a task:")
		fmt.Printf("ID: %d | Title %s \n", output.Id, output.Title)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
