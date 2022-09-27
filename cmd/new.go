/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/todo/methods"
	proto "example.com/todo/todo"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	Client:=methods.Connection()

		var task string
		var status int64 =0
		fmt.Print("Enter the task: ")
		fmt.Scan(&task)
		// fmt.Print("Enter the status: ")
		// fmt.Scan(&status)
		
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	
	//-----------------------------------------------------------------
	
		// Call Create
		req1 := proto.Task{
			Task:   task,
			Status: status,
		}
		a, err := Client.AddTask(ctx, &req1)
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		log.Printf("Create result: ID: %+v--- Task: %+v---Status: %+v\n\n", a.Id,a.Task,a.Status)
	},
}

func init() {
	noteCmd.AddCommand(newCmd)

}

// type promptContent struct {
// 	errorMsg string
// 	label    string
// }

// func promptGetInput(pc promptContent) string {
// 	validate := func(input string) error {
// 		if len(input) <= 0 {
// 			return errors.New(pc.errorMsg)
// 		}
// 		return nil
// 	}

// 	templates := &promptui.PromptTemplates{
// 		Prompt:  "{{ . }} ",
// 		Valid:   "{{ . | green }} ",
// 		Invalid: "{{ . | red }} ",
// 		Success: "{{ . | bold }} ",
// 	}

// 	prompt := promptui.Prompt{
// 		Label:     pc.label,
// 		Templates: templates,
// 		Validate:  validate,
// 	}

// 	result, err := prompt.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Input: %s\n", result)

// 	return result
// }

// func promptGetSelect(pc promptContent) string {
// 	items := []string{"animal", "food", "person", "object"}
// 	index := -1
// 	var result string
// 	var err error

// 	for index < 0 {
// 		prompt := promptui.SelectWithAdd{
// 			Label:    pc.label,
// 			Items:    items,
// 			AddLabel: "Other",
// 		}

// 		index, result, err = prompt.Run()

// 		if index == -1 {
// 			items = append(items, result)
// 		}
// 	}

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Input: %s\n", result)

// 	return result
// }

// func createNewNote() {
// 	wordPromptContent := promptContent{
// 		"Please provide a word.",
// 		"What word would you like to make a note of?",
// 	}
// 	word := promptGetInput(wordPromptContent)

// 	definitionPromptContent := promptContent{
// 		"Please provide a definition.",
// 		fmt.Sprintf("What is the definition of the %s?", word),
// 	}
// 	definition := promptGetInput(definitionPromptContent)

// 	categoryPromptContent := promptContent{
// 		"Please provide a category.",
// 		fmt.Sprintf("What category does %s belong to?", word),
// 	}
// 	category := promptGetSelect(categoryPromptContent)

// 	data.InsertNote(word, definition, category)
// }
