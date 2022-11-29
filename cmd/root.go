package cmd 

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Long: `This is a task manager to manage,guide and remind me to do tasks.`,
	 
  }