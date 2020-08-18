package cmd

import (
	"github.com/LabKom-Unila/health-checker/cmd/mysql"
	"github.com/LabKom-Unila/health-checker/cmd/pg"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pgCmd = &cobra.Command{
	Use:   "postgres",
	Short: "PostgreSQL",
	Long:  `Postgres service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var mySqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL",
	Long:  `MySQL service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var redisCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL",
	Long:  `MySQL service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	pgCmd.AddCommand(pg.PingCmd)
	mySqlCmd.AddCommand(mysql.PingCmd)
	rootCmd.AddCommand(pgCmd, mySqlCmd)
}
