package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// cmd1Cmd represents the cmd1 command
var cmd1Cmd = &cobra.Command{
	Use:   "cmd1",
	Short: "run cmd1",
	Long:  `project cmd1`,
	Run: func(cmd *cobra.Command, args []string) {
		runCmd1(cmd)
	},
}

func init() {
	rootCmd.AddCommand(cmd1Cmd)

}

func runCmd1(cmd *cobra.Command) {
	log.Print("cmd1 is running")
}
