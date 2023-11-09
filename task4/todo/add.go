package todo

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your task",

	Run: func(cmd *cobra.Command, args []string) {
	  task := strings.Join(args, " ")
	  fmt.Printf("Added \"%s\" in your task list\n", task)
	},
  }
  
  func init(){
	RootCmd.AddCommand(addCmd)
  }
