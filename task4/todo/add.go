package todo

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your task",

	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("add called")
	},
  }
  
  func int(){
	RootCmd.AddCommand(addCmd)
  }
