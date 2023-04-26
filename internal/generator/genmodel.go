package generator

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type genModel struct {
	TypeName string
	Name     string
	Gui      bool
	Nav      bool
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
		genModel := genModel{TypeName: typeParts[1], Name: nameParts[1], Gui: false}
		if len(parts) > 2 {
			guiParts := strings.Split(parts[2], ":")
			if guiParts[1] == "yes" {
				genModel.Gui = true
			}
		}
		if len(parts) > 3 {
			navParts := strings.Split(parts[3], ":")
			if navParts[1] == "yes" {
				genModel.Nav = true
			}
		}

		genModels = append(genModels, genModel)
	}

	return genModels
}
