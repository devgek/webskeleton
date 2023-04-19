package generator

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed gui_templates/template.html
var entityListTemplate string

//go:embed gui_templates/template-edit.html
var entityEditTemplate string

type GuiGenerator struct{}

func (gg GuiGenerator) Do(modelsPath string, genPath string) {
	log.Println("Start generating gui templates in ", genPath)
	os.Mkdir(genPath, os.ModePerm)

	genModels := getGenModels(modelsPath)

	gg.generateGuiTemplates(genModels, genPath)
}

func (gg GuiGenerator) generateGuiTemplates(models []genModel, genPath string) {
	tList := entityListTemplate
	tEdit := entityEditTemplate

	for _, genModel := range models {
		if genModel.Gui {
			log.Println("Generating gui templates for entity", genModel.Name)
			rList := strings.ReplaceAll(tList, "{{EntityName}}", genModel.Name)
			rEdit := strings.ReplaceAll(tEdit, "{{EntityName}}", genModel.Name)

			listPath := filepath.Join(genPath, genModel.Name+"-gen.html")
			err := ioutil.WriteFile(listPath, []byte(rList), os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}

			editPath := filepath.Join(genPath, genModel.Name+"-edit-gen.html")
			err1 := ioutil.WriteFile(editPath, []byte(rEdit), os.ModePerm)
			if err1 != nil {
				log.Fatalln(err)
			}
		}
	}

}
