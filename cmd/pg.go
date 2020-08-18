package cmd

import (
	"github.com/LabKom-Unila/health-checker/cmd/pg"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pgCmd = &cobra.Command{
	Use:   "postgres",
	Short: "pg",
	Long:  `Postgres service.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	pgCmd.AddCommand(pg.PingCmd)
	rootCmd.AddCommand(pgCmd)
}
