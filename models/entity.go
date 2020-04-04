package models

//EntityFactory create Entities by name
type EntityFactory struct {
}

//Get return entity struct by name
func (ef EntityFactory) Get(entityName string) interface{} {
	switch entityName {
	case "user":
		return &User{}
	case "contact":
		return &Contact{}
	default:
		panic("Undefind entity " + entityName)
	}
}

//GetSlice return slice of entity struct by name
func (ef EntityFactory) GetSlice(entityName string) interface{} {
	switch entityName {
	case "user":
		return &[]User{}
	case "contact":
		return &[]Contact{}
	default:
		panic("Undefind entity " + entityName)
	}
}
