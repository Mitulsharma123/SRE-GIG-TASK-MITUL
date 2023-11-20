package todo

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Task struct {
	ID   int
	Text string
}

type TaskList struct {
	Tasks []Task
}

var Tasklist TaskList

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your task",

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("List of task")
		for _, task := range Tasklist.Tasks {
			fmt.Printf("%d. %s\n", task.ID, task.Text)
			fmt.Print(task)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
