package todo

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do task added in your task list",

	Run: func(cmd *cobra.Command, args []string) {
	// iterate over list of tasks that is stored in slice of int 
	 var ids []int
	 for _, arg := range args{
		id, err := strconv.Atoi(arg)	// convert string into int 
		if err != nil{
			fmt.Println("failed to parse the argument", arg)
		}else{
			ids = append(ids, id)
		}
	 }
	 fmt.Println("Do task number:", ids)
  },
}
  
  func init(){
	RootCmd.AddCommand(doCmd)
  }
