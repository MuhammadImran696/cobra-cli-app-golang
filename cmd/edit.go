/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"example.com/todo/methods"
	proto "example.com/todo/todo"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Client := methods.Connection()

		id := getId()
		// fmt.Print(id)
		// t := data.Test2(id)
		// fmt.Print(t)

		// op := Selection()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req := proto.Args{
			Id: id,
		}
		// Option := Selection()
		response, err := Client.MarkTask(ctx, &req)
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		// log.Printf("Create result: ID: %+v--- Task: %+v---Status: %+v\n\n", response.Id,response.Task,response.Status)
		log.Println(response)

		//-----------------------------------------------------------------

	},
}

func init() {
	noteCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type pmtContent struct {
	errorMsg string
	label    string
}

func pmtGetSelect(pc pmtContent) string {
	items := []string{"Done", "Pending"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
func getId() int64 {

	var num int64
	fmt.Println("Enter the ID of Task to mark: ")
	fmt.Scan(&num)
	return num
}
func Selection() string {
	log.Println("Selection() called")

	CategoryPromptContent := pmtContent{
		"Please select an option.",
		fmt.Sprintf("What category does belong to?"),
	}
	result := pmtGetSelect(CategoryPromptContent)
	return result
}
