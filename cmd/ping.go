package cmd

import (
	"github.com/LabKom-Unila/health-checker/cmd/subcmd"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	pingCmd.AddCommand(subcmd.RedisCmd, subcmd.MysqlCmd)
	rootCmd.AddCommand(pingCmd)
}
