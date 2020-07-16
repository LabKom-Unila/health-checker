package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Check if MySQL is alive or not",
	Long:  `Check ping for MySQL.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mysql called")
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mysqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mysqlCmd.Flags().BoolP("host", "h", false, "Host For MySQL")
}
