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

//go:embed gui_templates/template-entity-nav.html
var navTemplate string

//go:embed gui_templates/template-entity-nav_1.html
var navTemplate1 string

type GuiGenerator struct{}

func (gg GuiGenerator) Do(modelsPath string, genPath string) {
	log.Println("Start generating gui templates in ", genPath)
	os.Mkdir(genPath, os.ModePerm)

	genModels := getGenModels(modelsPath)

	gg.generateGuiTemplates(genModels, genPath)
	gg.generateNavTemplate(genModels, genPath)
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

func (gg GuiGenerator) generateNavTemplate(models []genModel, genPath string) {
	t := navTemplate
	t1 := navTemplate1

	b1 := strings.Builder{}
	for _, genModel := range models {
		if genModel.Nav {
			rt1 := strings.ReplaceAll(t1, "{{EntityName}}", genModel.Name)
			b1.WriteString(rt1)
		}
	}

	t = strings.ReplaceAll(t, "{{EntityNav1}}", b1.String())

	entityNavPath := filepath.Join(genPath, "entity_nav.html")
	err := ioutil.WriteFile(entityNavPath, []byte(t), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
