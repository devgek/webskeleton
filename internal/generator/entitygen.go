package generator

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed entity_templates/factory_entity.template
var feTemplate string

//go:embed entity_templates/factory_entity_1.template
var feTemplate1 string

//go:embed entity_templates/factory_entity_2.template
var feTemplate2 string

//go:embed entity_templates/type_entity.template
var teTemplate string

//go:embed entity_templates/type_entity_3.template
var teTemplate3 string

type EntityGenerator struct{}

func (eg EntityGenerator) Do(modelsPath string, genPath string, modulePath string) {
	log.Println("Start generating entity types and factory in ", genPath, " for module "+modulePath)
	os.Mkdir(genPath, os.ModePerm)

	genModels := getGenModels(modelsPath)

	eg.generateEntityTypes(genModels, genPath)

	eg.generateEntityFactory(genModels, genPath, modulePath)
}

func (eg EntityGenerator) generateEntityFactory(models []genModel, modelsPath string, modulePath string) {
	t := feTemplate
	t1 := feTemplate1
	t2 := feTemplate2

	b1 := strings.Builder{}
	b2 := strings.Builder{}
	for _, genModel := range models {
		rt1 := strings.ReplaceAll(t1, "{{EntityName}}", genModel.Name)
		rt1 = strings.ReplaceAll(rt1, "{{EntityTypeName}}", genModel.TypeName)
		b1.WriteString(rt1)

		rt2 := strings.ReplaceAll(t2, "{{EntityName}}", genModel.Name)
		rt2 = strings.ReplaceAll(rt2, "{{EntityTypeName}}", genModel.TypeName)
		b2.WriteString(rt2)
	}

	t = strings.ReplaceAll(t, "{{ModulePath}}", modulePath)
	t = strings.ReplaceAll(t, "{{FactoryEntity1}}", b1.String())
	t = strings.ReplaceAll(t, "{{FactoryEntity2}}", b2.String())

	entityFactoryPath := filepath.Join(modelsPath, "entity_factory_impl.go")
	err := ioutil.WriteFile(entityFactoryPath, []byte(t), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

func (eg EntityGenerator) generateEntityTypes(models []genModel, modelsPath string) {
	t := teTemplate
	t3 := teTemplate3

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
