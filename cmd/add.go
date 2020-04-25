package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/v26/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var collection mongo.Collection
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your tasks list\n", task)
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
