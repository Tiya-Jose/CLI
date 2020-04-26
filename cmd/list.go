package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

type task struct {
	Task string `json:"task"`
}

var tasks []task

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		collection.Find(bson.M{}, bson.M{}, &tasks)
		if tasks != nil {
			fmt.Println()
			fmt.Println("You have the following tasks in your To-Do list:")
			for i, v := range tasks {
				fmt.Println(strconv.Itoa(i+1)+".", v.Task)
			}
		} else {
			fmt.Println()
			fmt.Println("You have not added any tasks to your To-Do list.\nPlease add your tasks")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
