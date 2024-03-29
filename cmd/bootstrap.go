package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devgek/webskeleton/helper/common"
	"github.com/devgek/webskeleton/helper/fileutil"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap a new go web project",
	Long:  `webskeleton bootstrap; Bootstraps a typical go web app using sqlite database, a layout template + login form`,
	Run: func(cmd *cobra.Command, args []string) {
		runBootstrap(cmd)
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	bootstrapCmd.Flags().String("templatedir", "", "The directory with the template files")
	bootstrapCmd.Flags().String("type", "web", "The type of project you want to bootstrap [api|web|cli]")
	bootstrapCmd.Flags().String("repository", "github.com", "The git repository for the new project")
	bootstrapCmd.Flags().String("user", "theuser", "The git user for the new project")
	bootstrapCmd.Flags().String("project", "theproject", "The project name for the new project")
	bootstrapCmd.Flags().String("title", "", "The title for this project, default is project name")
}

func runBootstrap(cmd *cobra.Command) {
	templateDir, _ := cmd.Flags().GetString("templatedir")
	projectType, _ := cmd.Flags().GetString("type")
	repoName, _ := cmd.Flags().GetString("repository")
	repoUser, _ := cmd.Flags().GetString("user")
	projectName, _ := cmd.Flags().GetString("project")
	projectTitle, _ := cmd.Flags().GetString("title")
	if projectTitle == "" {
		projectTitle = projectName
	}

	packageName := repoName + "/" + repoUser + "/" + projectName
	currPath, err := os.Getwd()
	helper.ExitOnError(err, "Can't get current path!")

	log.Println("Start bootstraping new project for", "'"+packageName+"' with title", projectTitle, "in", currPath)

	projectTemplateDir := templateDir
	if templateDir == "" {
		projectTemplateDir = templateDirFromGoPath()
	}

	// projectPath
	projectPath := filepath.Join(currPath, projectName)

	// 1. Create project directory
	log.Print("Creating " + projectPath + "...")
	err = os.MkdirAll(projectPath, 0755)
	helper.ExitOnError(err, "")

	// 2. Copy everything under project template directory to target directory.
	log.Print("Copying project template files and directories to " + projectPath + "...")
	// currDir, err := os.Getwd()
	// helper.ExitOnError(err, "Can't get current path!")

	// err = os.Chdir(projectTemplateDir)
	// helper.ExitOnError(err, "")

	copySources(getSources(projectTemplateDir, projectType), projectTemplateDir, projectPath)

	// err = os.Chdir(currPath)
	// helper.ExitOnError(err, "")

	// 3. Interpolate placeholder variables on the new project.
	log.Print("Replacing placeholder variables on " + repoUser + "/" + projectName + "...")

	replacers := make(map[string]string)
	replacers["github.com/devgek/webskeleton"] = packageName
	replacers["webskeleton.db"] = projectName + ".db"
	replacers["webskeleton-auth"] = projectName + "-auth"
	replacers[`= "webskeleton" //do not change`] = `= "` + projectName + `"`
	replacers["go-webskeleton"] = projectTitle
	replacers[".webskeleton.yaml"] = "." + projectName + ".yaml"
	replacers["webskeleton_T_"] = projectName + "_T_"
	replacers["webskeleton-types"] = projectName + "-types"
	replacers["webskeleton.com"] = projectName + ".com"
	replacers["../webskeleton"] = "../" + projectName
	replacers[`"ProjectName":"webskeleton"`] = `"ProjectName":"` + projectName + `"`
	err = recursiveSearchReplaceFiles(projectPath, replacers)
	helper.ExitOnError(err, "")

	// 4. Setup and bootstrap databases.
	// nothing to do, yet

	// // 5. Get all application dependencies for the first time.
	// log.Print("Running go get ./...")
	// command = exec.Command("go", "get", "./...")
	// command.Dir = fullpath
	// output, err = command.CombinedOutput()
	// helper.ExitOnError(err, string(output))
	err = os.Chdir(projectPath)
	helper.ExitOnError(err, "")

	// 5. Initialize a go module project
	//log.Print("Running go mod init ", packageName)
	//command := exec.Command("go", "mod", "init", packageName)
	//command.Dir = projectPath
	//output, _ := command.CombinedOutput()
	//log.Print(string(output))

	log.Print("Running go mod tidy")
	command := exec.Command("go", "mod", "tidy")
	command.Dir = projectPath
	output, _ := command.CombinedOutput()
	log.Print(string(output))

	//6. Run tests on newly generated app.
	log.Print("Running go test ./...")
	command = exec.Command("go", "test", "./...")
	command.Dir = projectPath
	output, _ = command.CombinedOutput()
	log.Print(string(output))
}

func templateDirFromGoPath() string {
	// There can be more than one path, separated by colon.
	gopaths := helper.GoPaths()
	if len(gopaths) == 0 {
		log.Fatalln("GOPATH is not set.")
	}
	// By default, we choose the last GOPATH.
	gopath := gopaths[len(gopaths)-1]
	projectTemplateDir := filepath.Join(gopath, "src", "github.com", "devgek", "webskeleton")
	return projectTemplateDir
}

func copySources(sourceLines []string, sourceRoot, destinationRoot string) {
	for _, line := range sourceLines {
		parts := strings.Split(line, ";")
		cmd := parts[0]
		source := parts[1]
		sourcePath := filepath.Join(sourceRoot, source)
		destinationPath := destinationRoot
		if len(parts) == 3 {
			destinationPath = filepath.Join(destinationRoot, parts[2])
		}

		log.Print(cmd, " ", sourcePath, "--->", destinationPath)
		err := copy.Copy(sourcePath, destinationPath)
		helper.ExitOnError(err, "Error while copying source [files]cd")
	}
}

func getSources(rootPath string, projectType string) []string {
	fileName := filepath.Join(rootPath, "_test", "copy-"+projectType+".txt")
	fileName = filepath.Clean(fileName)

	sources, err := fileutil.ReadLines(fileName)
	helper.ExitOnError(err, "getSources")

	return sources
}

func recursiveSearchReplaceFiles(fullpath string, replacers map[string]string) error {
	fileOrDirList := []string{}
	err := filepath.Walk(fullpath, func(path string, f os.FileInfo, err error) error {
		fileOrDirList = append(fileOrDirList, path)
		return nil
	})

	if err != nil {
		return err
	}

	for _, fileOrDir := range fileOrDirList {
		fileInfo, _ := os.Stat(fileOrDir)
		if !fileInfo.IsDir() {
			for oldString, newString := range replacers {
				contentBytes, _ := ioutil.ReadFile(fileOrDir)
				newContentBytes := bytes.Replace(contentBytes, []byte(oldString), []byte(newString), -1)

				err := ioutil.WriteFile(fileOrDir, newContentBytes, fileInfo.Mode())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
