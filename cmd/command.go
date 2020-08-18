package cmd

import (
	"github.com/LabKom-Unila/health-checker/cmd/mysql"
	"github.com/LabKom-Unila/health-checker/cmd/pg"
	"github.com/LabKom-Unila/health-checker/cmd/redis"

	"github.com/spf13/cobra"
)

// pgCmd for postgres utilty
var pgCmd = &cobra.Command{
	Use:   "postgres",
	Short: "PostgreSQL",
	Long:  `Postgres service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// mySQLCmd for MySQL utilty
var mySQLCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL",
	Long:  `MySQL service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// redisCmd for Redis utilty
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Redis",
	Long:  `Redis service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// main command for CLI
	pgCmd.AddCommand(pg.PingCmd)
	mySQLCmd.AddCommand(mysql.PingCmd)
	redisCmd.AddCommand(redis.RedisCmd)

	// root command
	rootCmd.AddCommand(pgCmd, mySQLCmd, redisCmd)
}
