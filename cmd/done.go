package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// doCmd represents the do command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Remove the task from your task list ",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := strconv.Atoi(task)
		if err != nil {
			log.Println(err)
		}
		collection.Find(bson.M{}, bson.M{}, &tasks)
		for i, v := range tasks {
			if i == id-1 {
				collection.DeleteOne(bson.M{"task": v.Task})
			}
		}
		fmt.Println("done called")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
