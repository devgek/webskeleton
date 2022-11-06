package cmd

import (
	"github.com/devgek/webskeleton/types"
	"github.com/devgek/webskeleton/web/api/env"
	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "initialize the datastore with data",
	Long:  `webskeleton data`,
	Run: func(cmd *cobra.Command, args []string) {
		runData(cmd)
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)

}

func runData(cmd *cobra.Command) {
	env := env.GetApiEnv(false)
	env.Services.CreateUser("user1", "secret", "user1@webskeleton.com", types.RoleTypeUser)
	env.Services.CreateUser("user2", "secret", "user2@webskeleton.com", types.RoleTypeUser)
	env.Services.CreateUser("user3", "secret", "user3@webskeleton.com", types.RoleTypeUser)
	env.Services.CreateUser("user4", "secret", "user4@webskeleton.com", types.RoleTypeUser)
	env.Services.CreateUser("user5", "secret", "user5@webskeleton.com", types.RoleTypeUser)
}
