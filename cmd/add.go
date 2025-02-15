package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var isCompleted bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Long:  "Add a new todo item to the database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		_, err = DbPool.Exec(context.Background(), `insert into todos (title, completed) values ($1, $2)`, args[0], isCompleted)
		if err != nil {
			return fmt.Errorf("Error adding a new todo item: %w", err)
		}

		return nil
	},
}

func init() {
	addCmd.Flags().BoolVarP(&isCompleted, "completed", "c", false, "whether if the todo item is completed or not (default to false)")

	rootCmd.AddCommand(addCmd)
}
