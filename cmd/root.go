package cmd

import (
	"fmt"
	"os"

	"CLI-Task-Planner/mongo"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task Planner that is similar to creating a To-Do list and cutting off the done tasks",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func MongoConnect(dbName, cName string, flag bool) mongo.Collection {
	c := mongo.NewClient("", "", "localhost", 27017)
	c.Connection()
	return c.NewCollection(dbName, cName, flag)
}
