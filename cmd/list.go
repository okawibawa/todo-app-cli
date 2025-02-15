package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  "List available todos from the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		var todoSlice []Todo

		rows, err := DbPool.Query(context.Background(), `select * from todos`)
		if err != nil {
			return fmt.Errorf("error selecting from todos: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var todo Todo
			err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
			if err != nil {
				return fmt.Errorf("error scanning data into struct: %w", err)
			}
			todoSlice = append(todoSlice, todo)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 6, ' ', 0)
		writer.Write([]byte("ID\tTitle\tTime\tCompleted\n"))
		for _, todo := range todoSlice {
			durationSinceCreation := time.Since(todo.CreatedAt).Truncate(time.Hour)

			timeAgo := time.Now().Add(-durationSinceCreation)

			timeDiff := timediff.TimeDiff(timeAgo)

			str := fmt.Sprintf("%d\t%s\t%s\t%t\n", todo.Id, todo.Title, timeDiff, todo.Completed)

			writer.Write([]byte(str))
		}
		writer.Flush()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
