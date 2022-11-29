package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
) 

 var addCmd = &cobra.Command{ 
	Use:   "add",
	Short: "Add task to a list",
	Run: func(cmd *cobra.Command, args []string) {
		
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong",err.Error())
			return 
		}
		fmt.Printf("Added task %s to a Task List. ", task)
		
	  },
	 
  }


  func init(){
	RootCmd.AddCommand(addCmd)
  }