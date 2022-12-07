package genmodels

// EntityType type of entity
type EntityType int

const (
	EntityTypeUndefined EntityType = iota
	EntityTypeContact
	EntityTypeContactAddress
	EntityTypeUser
)

// EntityTypes ...
func EntityTypes() []EntityType {
	return []EntityType{EntityTypeUndefined, EntityTypeContact, EntityTypeContactAddress, EntityTypeUser}
}

// Val the value used in html template
func (et EntityType) Val() string {
	return [...]string{"undefined", "contact", "contactaddress", "user"}[et]
}

// EntityName the value used in html template
func (et EntityType) EntityName() string {
	return [...]string{"Undefined", "Contact", "ContactAddress", "User"}[et]
}

// ParseEntityType ...
func ParseEntityType(s string) EntityType {
	switch s {
	case EntityTypeContact.Val():
		return EntityTypeContact
	case EntityTypeContactAddress.Val():
		return EntityTypeContactAddress
	case EntityTypeUser.Val():
		return EntityTypeUser

	default:
		return EntityTypeUndefined
	}
}
