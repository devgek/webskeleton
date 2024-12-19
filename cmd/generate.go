package cmd

import (
	"github.com/devgek/webskeleton/helper/common"
	"github.com/devgek/webskeleton/internal/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate entity source code",
	Long:  `webskeleton generate; Generates source code needed for "entity-handling"`,
	Run: func(cmd *cobra.Command, args []string) {
		runGenerate(cmd)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().String("type", "db", "The type of generation you want to run [db|gui]")
	generateCmd.Flags().String("modelsdir", "./models", "The directory with the model files")
	generateCmd.Flags().String("path", "repository/user/project", "The module path, if type is db")
}

func runGenerate(cmd *cobra.Command) {
	modelsDir, _ := cmd.Flags().GetString("modelsdir")
	generationType, _ := cmd.Flags().GetString("type")
	modulePath, _ := cmd.Flags().GetString("path")

	currPath, err := os.Getwd()
	helper.ExitOnError(err, "Can't get current path!")

	modelsPath := filepath.Join(currPath, modelsDir)

	var genPath string
	if "db" == generationType {
		theGenerator := generator.EntityGenerator{}
		genPath = filepath.Join(modelsPath, "generated")
		theGenerator.Do(modelsPath, genPath, modulePath)
	} else if "gui" == generationType {
		generator := generator.GuiGenerator{}
		genPath = filepath.Join(currPath, "web", "app", "template", "templates")
		generator.Do(modelsPath, genPath)
	} else {
		log.Print("missing type=[db|gui]")
	}

	if "db" == generationType {
		log.Print("Running go fmt ", genPath)
		command := exec.Command("go", "fmt", genPath)
		command.Dir = currPath
		output, _ := command.CombinedOutput()
		log.Print(string(output))
	}
}
