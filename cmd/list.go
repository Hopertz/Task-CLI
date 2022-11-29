package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
) 

 var listCmd = &cobra.Command{ 
	Use:   "list",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks,err := db.AllTasks()
        if err!= nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		} 
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Take a break and chilax bro")
			return
		}
		fmt.Println("You have the following Tasks.")
		for i,task := range tasks {
			fmt.Printf("\n%d. %s.\n", i+1, task.Value)
		}
	  },
	 
  }


  func init(){
	RootCmd.AddCommand(listCmd)
  }