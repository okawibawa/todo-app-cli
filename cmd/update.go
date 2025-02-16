package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a todo item",
	Long:  "Update a todo item's title or completed status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return fmt.Errorf("error getting flag value: %w", err)
		}

		isCompleted, err := cmd.Flags().GetBool("completed")
		if err != nil {
			return fmt.Errorf("error getting flag value: %w", err)
		}

		query := `update todos set `
		var updates []string
		var queryArgs []interface{}
		argIndex := 1

		if title != "" {
			updates = append(updates, fmt.Sprintf(`title = $%d`, argIndex))
			queryArgs = append(queryArgs, title)
			argIndex++
		}

		if cmd.Flags().Changed("completed") {
			updates = append(updates, fmt.Sprintf(`completed = $%d`, argIndex))
			queryArgs = append(queryArgs, isCompleted)
			argIndex++
		}

		query += strings.Join(updates, ", ") + fmt.Sprintf(` where id = $%d`, argIndex)
		queryArgs = append(queryArgs, args[0])

		_, err = DbPool.Exec(context.Background(), query, queryArgs...)
		if err != nil {
			return fmt.Errorf("error updating data to the database: %w", err)
		}

		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("title", "t", "", "new todo item title")
	updateCmd.Flags().BoolP("completed", "c", false, "whether if the todo item is completed or not (default to false)")
	updateCmd.MarkFlagsOneRequired("title", "completed")
	rootCmd.AddCommand(updateCmd)
}
