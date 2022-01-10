/*
	Type EntityType and function Val must exist even before first generating models, because they are used
	in entityservicesimpl.go
*/
package generated_models

//EntityType type of entity
type EntityType int

//Val the value used in html template
func (et EntityType) Val() string {
	return "undefined"
}

//ParseEntityType ...
func ParseEntityType(s string) EntityType {
	return 0
}
