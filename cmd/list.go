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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Client:=methods.Connection()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

		req2:= proto.List{
			Message: "return list",
		}
		 response,err := Client.ListTask(ctx, &req2)
			if err != nil {
				log.Fatalf("Create failed: %v", err)
			}
		
			//  fmt.Println(response)
			
				// fmt.Print("hi ",response.Items[])
				for i := 0; i < len(response.Items); i++ {
					if response.Items[i].Status== 0 {
				fmt.Printf("ReadAll result: ID: %+v   ---- Task: %+v   ---- Status: %+v\n\n", response.Items[i].Id,response.Items[i].Task,"Pending")
					}else{
				fmt.Printf("ReadAll result: ID: %+v   ---- Task: %+v   ---- Status: %+v\n\n", response.Items[i].Id,response.Items[i].Task,"Done")
		
					}
					
				}
	},
}

func init() {
	noteCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
