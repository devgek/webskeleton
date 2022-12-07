package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devgek/webskeleton/helper/common"
	"github.com/spf13/cobra"
)

type genModel struct {
	TypeName string
	Name     string
}

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
	log.Println("Start generating entity types and factory in ", genPath)
	os.Mkdir(genPath, os.ModePerm)

	genModels := getGenModels(modelsPath)

	templatePath := filepath.Join(currPath, "_template")

	generateEntityTypes(genModels, templatePath, genPath)

	generateEntityFactory(genModels, templatePath, genPath)

	log.Print("Running go fmt ", genPath)
	command := exec.Command("go", "fmt", genPath)
	command.Dir = currPath
	output, _ := command.CombinedOutput()
	log.Print(string(output))
}

func generateEntityFactory(models []genModel, templatePath string, modelsPath string) {
	t := readStringTemplate(templatePath, "factory_entity.template")
	t1 := readStringTemplate(templatePath, "factory_entity_1.template")
	t2 := readStringTemplate(templatePath, "factory_entity_2.template")
	t3 := readStringTemplate(templatePath, "factory_entity_3.template")

	f1 := strings.Builder{}
	f2 := strings.Builder{}
	f3 := strings.Builder{}
	for _, genModel := range models {
		rt1 := strings.ReplaceAll(t1, "{{EntityTypeName}}", genModel.TypeName)
		rt2 := strings.ReplaceAll(t2, "{{EntityTypeName}}", genModel.TypeName)
		rt3 := strings.ReplaceAll(t3, "{{EntityTypeName}}", genModel.TypeName)

		f1.WriteString(rt1)
		f2.WriteString(rt2)
		f3.WriteString(rt3)
	}

	t = strings.ReplaceAll(t, "{{FactoryEntity1}}", f1.String())
	t = strings.ReplaceAll(t, "{{FactoryEntity2}}", f2.String())
	t = strings.ReplaceAll(t, "{{FactoryEntity3}}", f3.String())

	entityFactoryPath := filepath.Join(modelsPath, "entity_factory_impl.go")
	err := ioutil.WriteFile(entityFactoryPath, []byte(t), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateEntityTypes(models []genModel, templatePath string, modelsPath string) {
	t := readStringTemplate(templatePath, "type_entity.template")
	t3 := readStringTemplate(templatePath, "type_entity_3.template")

	f1 := strings.Builder{}
	f2 := strings.Builder{}
	f3 := strings.Builder{}
	f4 := strings.Builder{}
	f5 := strings.Builder{}
	for _, genModel := range models {
		f1.WriteString("EntityType" + genModel.TypeName + "\n")
		f2.WriteString(", EntityType" + genModel.TypeName)

		rt3 := strings.ReplaceAll(t3, "{{EntityTypeName}}", genModel.TypeName)
		f3.WriteString(rt3)

		f4.WriteString(", \"" + genModel.Name + "\"")
		f5.WriteString(", \"" + genModel.TypeName + "\"")
	}

	t = strings.ReplaceAll(t, "{{TypeEntity1}}", f1.String())
	t = strings.ReplaceAll(t, "{{TypeEntity2}}", f2.String())
	t = strings.ReplaceAll(t, "{{TypeEntity3}}", f3.String())
	t = strings.ReplaceAll(t, "{{TypeEntity4}}", f4.String())
	t = strings.ReplaceAll(t, "{{TypeEntity5}}", f5.String())

	typePath := filepath.Join(modelsPath, "entity_types_impl.go")
	err := ioutil.WriteFile(typePath, []byte(t), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func getGenModels(path string) []genModel {
	genModels := []genModel{}
	entityLinesList := []string{}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") {
			contentBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			content := string(contentBytes)
			if strings.Contains(content, "entitymodel.GormEntity") && strings.Contains(content, "entity:") {
				from := strings.Index(content, "entity:")
				sEnd := content[from+8:]
				to := strings.Index(sEnd, "\"")
				sMatch := sEnd[:to]
				entityLinesList = append(entityLinesList, sMatch)
			}

		}
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	for _, el := range entityLinesList {
		parts := strings.Split(el, ";")
		typeParts := strings.Split(parts[0], ":")
		nameParts := strings.Split(parts[1], ":")
		genModel := genModel{TypeName: typeParts[1], Name: nameParts[1]}

		genModels = append(genModels, genModel)
	}

	return genModels
}

func readStringTemplate(templatePath string, templateName string) string {
	filename := filepath.Join(templatePath, templateName)
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return string(contentBytes)
}
