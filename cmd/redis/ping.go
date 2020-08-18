package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/spf13/cobra"
)

// RedisCmd represents the redis command
var RedisCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if Redis is alive or not",
	Long:  `Check ping for Redis.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		password, _ := cmd.Flags().GetString("password")

		pingRedis(address, password)
	},
}

func init() {
	RedisCmd.Flags().StringP("address", "a", "127.0.0.1:6379", "Host for Redis")
	RedisCmd.Flags().StringP("password", "p", "", "Password for Redis")
}

func pingRedis(address, password string) {
	conn := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	_, err := conn.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(("\033[31m"), err)
		return
	}

	fmt.Println(("\033[32m"), "Ping sucessful")
}
