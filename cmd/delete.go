package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single todo",
	Long:  "Delete a single todo from the database",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("argument must be an integer")
		}

		commandTag, err := DbPool.Exec(context.Background(), "delete from todos where id = $1", args[0])
		if err != nil {
			return fmt.Errorf("error deleting data: %w", err)
		}

		if commandTag.RowsAffected() == 0 {
			return fmt.Errorf("todo item with the id of %s is not found", args[0])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
