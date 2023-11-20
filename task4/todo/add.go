package todo

import (
	"fmt"
	//"strings"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your task",

	Run: func(cmd *cobra.Command, args []string) {
		//taskText := strings.Join(args, " ")
		//fmt.Printf("Added \"%s\" in your task list\n", t)
		taskText := args[0]
		newTask := Task{
			ID:   len(Tasklist.Tasks) + 1,
			Text: taskText,
		}
		Tasklist.Tasks = append(Tasklist.Tasks, newTask)
		fmt.Sprintf("Task %s added successfully! in your task list\n", Tasklist.Tasks)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
