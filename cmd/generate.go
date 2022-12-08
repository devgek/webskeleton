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

	generateCmd.Flags().String("modelsdir", "./models", "The directory with the model files")
}

func runGenerate(cmd *cobra.Command) {
	modelsDir, _ := cmd.Flags().GetString("modelsdir")
	currPath, err := os.Getwd()
	helper.ExitOnError(err, "Can't get current path!")

	modelsPath := filepath.Join(currPath, modelsDir)
	genPath := filepath.Join(modelsPath, "generated")
	templatePath := filepath.Join(currPath, "_template")

	generator := generator.EntityGenerator{}
	generator.Do(modelsPath, genPath, templatePath)

	log.Print("Running go fmt ", genPath)
	command := exec.Command("go", "fmt", genPath)
	command.Dir = currPath
	output, _ := command.CombinedOutput()
	log.Print(string(output))
}
