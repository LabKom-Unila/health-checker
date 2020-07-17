package cmd

import (
	"github.com/LabKom-Unila/health-checker/cmd/subcmd"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping",
	Long:  `Check various service before you use.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	pingCmd.AddCommand(subcmd.RedisCmd, subcmd.MysqlCmd)
	rootCmd.AddCommand(pingCmd)
}
