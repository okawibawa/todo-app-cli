package cmd

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okawibawa/todo-app-cli/internal/database"
	"github.com/spf13/cobra"
)

var DbPool *pgxpool.Pool

var rootCmd = &cobra.Command{
	Use:               "todo",
	Short:             "CLIs for a simple todo list app.",
	Long:              "todo-cli is a command line interface application for a siple todo list app.",
	PersistentPreRunE: connectDb,
	PersistentPostRun: disconnectDb,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func connectDb(cmd *cobra.Command, args []string) error {
	var err error

	DbPool, err = database.InitDb()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	return nil
}

func disconnectDb(cmd *cobra.Command, args []string) {
	if DbPool != nil {
		DbPool.Close()
	}
}
