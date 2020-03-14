package cmd

import (
	"github.com/devgek/webskeleton/helper"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	bootstrapCmd.Flags().String("web", "echo", "The web framework to use [echo|mux]")
}

func runBootstrap(cmd *cobra.Command) {
	// start the web server
	repoName, _ := cmd.Flags().GetString("repository")
	repoUser, _ := cmd.Flags().GetString("user")
	projectName, _ := cmd.Flags().GetString("project")
	webFramework, _ := cmd.Flags().GetString("web")

	packageName := repoName + "/" + repoUser + "/" + projectName
	log.Println("Start bootstraping new project for", "'"+packageName+"' with webframework", webFramework)

	// There can be more than one path, separated by colon.
	gopaths := helper.GoPaths()
	if len(gopaths) == 0 {
		log.Fatalln("GOPATH is not set.")
	}
	// By default, we choose the last GOPATH.
	gopath := gopaths[len(gopaths)-1]

	fullpath := filepath.Join(gopath, "src", packageName)
	dbName := projectName
	projectTemplateDir := filepath.Join(gopath, "src", "github.com", "devgek", "webskeleton")

	// 1. Create target directory
	log.Print("Creating " + fullpath + "...")
	err := os.MkdirAll(fullpath, 0755)
	helper.ExitOnError(err, "")

	// 2. Copy everything under project template directory to target directory.
	log.Print("Copying project template directory to " + fullpath + "...")
	currDir, err := os.Getwd()
	helper.ExitOnError(err, "Can't get current path!")

	err = os.Chdir(projectTemplateDir)
	helper.ExitOnError(err, "")

	var command *exec.Cmd
	if helper.IsWindows() {
		command = exec.Command("xcopy", ".", fullpath, "/S", "/E", "/H", "/Y", "/EXCLUDE:exclude.txt")
	} else {
		command = exec.Command("cp", "-rf", ".", fullpath)
	}
	output, err := command.CombinedOutput()
	helper.ExitOnError(err, string(output))

	err = os.Chdir(currDir)
	helper.ExitOnError(err, "")

	// 3. Interpolate placeholder variables on the new project.
	log.Print("Replacing placeholder variables on " + repoUser + "/" + projectName + "...")

	replacers := make(map[string]string)
	replacers["github.com/devgek/webskeleton"] = packageName
	replacers["webskeleton.db"] = dbName + ".db"
	replacers["webskeleton-auth"] = projectName + "-auth"
	replacers["go-webskeleton"] = projectName
	err = helper.RecursiveSearchReplaceFiles(fullpath, replacers)
	helper.ExitOnError(err, "")

	// 4. Setup and bootstrap databases.
	// nothing to do, yet

	// 5. Get all application dependencies for the first time.
	log.Print("Running go get ./...")
	command = exec.Command("go", "get", "./...")
	command.Dir = fullpath
	output, err = command.CombinedOutput()
	helper.ExitOnError(err, string(output))

	// 6. Run tests on newly generated app.
	log.Print("Running go test ./...")
	command = exec.Command("go", "test", "./...")
	command.Dir = fullpath
	output, _ = command.CombinedOutput()
	log.Print(string(output))
}
