package models

//
//EntityType type of entity
//
type EntityType int

//
const (
	EntityTypeUndefined EntityType = iota
	EntityTypeUser
	EntityTypeContact
	EntityTypeContactAddress
)

//EntityTypes ...
func EntityTypes() []EntityType {
	return []EntityType{EntityTypeUndefined, EntityTypeUser, EntityTypeContact, EntityTypeContactAddress}
}

//Val the value used in html template
func (et EntityType) Val() string {
	return [...]string{"undefined", "user", "contact", "contactaddress"}[et]
}

//ParseEntityType ...
func ParseEntityType(s string) EntityType {
	if s == EntityTypeUser.Val() {
		return EntityTypeUser
	} else if s == EntityTypeContact.Val() {
		return EntityTypeContact
	} else if s == EntityTypeContactAddress.Val() {
		return EntityTypeContactAddress
	}
	return EntityTypeUndefined
}
