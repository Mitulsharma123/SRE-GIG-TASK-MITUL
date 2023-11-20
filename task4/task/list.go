package todo

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your task",

	Run: func(cmd *cobra.Command, args []string) {
		// read content from file 
		content, error := os.ReadFile("file.txt")
		if error != nil {
			fmt.Println("Error in reading file")
		}
		// str stores in string format , as content is byte slice
  		str := string(content)
		fmt.Printf("List of task: %s", str)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
