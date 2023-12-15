package todo

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your task",

	Run: func(cmd *cobra.Command, args []string) {
		// iterate over task with index and task name
		/*for i, arg := range args{
			fmt.Printf("%d : %s\n", i, arg)
		}*/

		// create a single string to store tasks
		taskText := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" in your task list\n", taskText)
		
		// create a text file to store the list of task
		f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//if err != nil {
		//	fmt.Println("Not able to open this file")
		//}
    if err != nil {
	    fmt.Fprintf(os.Stderr, "%v\n", err)
    }

		//close the file
		defer f.Close()

		// below logic stores the task in line by line way 
		/*
		for _, arg := range args {
			_, err := f.WriteString(arg + " " + "\n")
			if err != nil{
				fmt.Print("Not able add txt\n")
			}
		}*/


		// write task strings to file 
		if _, err := f.WriteString(taskText + " " + "\n"); err != nil {
			fmt.Print("Not able add txt \n")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
