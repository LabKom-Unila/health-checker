package pg

import (
	"database/sql"
	"fmt"

	"github.com/LabKom-Unila/health-checker/helper"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// PingCmd represents the pg command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if PostgreSQL is alive or not",
	Long:  `Check ping for PostgreSQL.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		db, _ := cmd.Flags().GetString("db")

		pingPgsql(address, user, password, db)
	},
}

func init() {
	PingCmd.Flags().StringP("address", "a", "127.0.0.1:5432", "Host for PostgreSQL")
	PingCmd.Flags().StringP("user", "u", "postgres", "User for PostgreSQL")
	PingCmd.Flags().StringP("password", "p", "", "Password for PostgreSQL")
	PingCmd.Flags().StringP("database", "d", "postgres", "Database for PostgreSQL")
}

func pingPgsql(address, user, password, db string) {
	dsn := fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, user, password, address, db)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		helper.ErrorOutput(err)
		return
	}
	if err = conn.Ping(); err != nil {
		helper.ErrorOutput(err)
		return
	}

	helper.SuccessOutput()
}
