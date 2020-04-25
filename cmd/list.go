package cmd

import (
	"fmt"

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
		fmt.Println("list called")
		collection.Find(bson.M{}, bson.M{}, &tasks)
		for i, v := range tasks {
			fmt.Println(i+1, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
