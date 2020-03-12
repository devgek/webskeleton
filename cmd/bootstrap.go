package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap -repository ",
	Short: "bootstrap a new go web project",
	Long:  `webskeleton bootstrap; a typical go web app`,
	Run: func(cmd *cobra.Command, args []string) {
		runBootstrap(cmd)
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bootstrapCmd.Flags().String("repository", "github.com", "The git repository for the new project")
	bootstrapCmd.Flags().String("user", "theuser", "The git user for the new project")
	bootstrapCmd.Flags().String("project", "theproject", "The project name for the new project")
}

func runBootstrap(cmd *cobra.Command) {
	// start the web server
	repository, _ := cmd.Flags().GetString("repository")
	user, _ := cmd.Flags().GetString("user")
	project, _ := cmd.Flags().GetString("project")

	log.Println("Start bootstraping new project for", "'"+repository+"/"+user+"/"+project+"'")

}
