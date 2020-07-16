package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Check if MySQL is alive or not",
	Long:  `Check ping for MySQL.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")

		log.Println(address)
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.Flags().String("address", "127.0.0.1:6379", "Host for MySQL")
}
