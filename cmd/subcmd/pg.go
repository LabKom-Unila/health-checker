package subcmd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// PgCmd represents the pg command
var PgCmd = &cobra.Command{
	Use:   "pg",
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
	PgCmd.Flags().StringP("address", "a", "127.0.0.1:5432", "Host for PostgreSQL")
	PgCmd.Flags().StringP("user", "u", "postgres", "User for PostgreSQL")
	PgCmd.Flags().StringP("password", "p", "", "Password for PostgreSQL")
	PgCmd.Flags().StringP("database", "d", "postgres", "Database for PostgreSQL")
}

func pingPgsql(address, user, password, db string) {
	dsn := fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, user, password, address, db)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = conn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ping sucessful")
}
