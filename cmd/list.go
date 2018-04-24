package cmd

import (
	"github.com/samparsky/task/db"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "List",
	Short: "lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string){
		tasks , err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong..", err.Error())
			return
		}

		if(len(tasks) == 0){
			fmt.Println("You have no tasks to complete! Why not take a break")
			return
		}
		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init(){
	RootCmd.AddCommand(listCmd)
}