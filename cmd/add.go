package cmd

import (
	"fmt"
	"log"
	"strings"

	"task/mongo"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var collection mongo.Collection
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your To-Do list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Println()
		fmt.Printf("Added \"%s\" to your To-Do list\n", task)
		err := collection.InsertOne(bson.M{"task": task})
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func init() {
	collection = MongoConnect("taskPlanner", "tasks", true)
}
