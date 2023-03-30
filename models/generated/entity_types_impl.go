package genmodels

// EntityType type of entity
type EntityType int

const (
	EntityTypeUndefined EntityType = iota
	EntityTypeAccount
	EntityTypeContact
	EntityTypeContactAddress
	EntityTypeUser
)

// EntityTypes ...
func EntityTypes() []EntityType {
	return []EntityType{EntityTypeUndefined, EntityTypeAccount, EntityTypeContact, EntityTypeContactAddress, EntityTypeUser}
}

// Val the value used in html template
func (et EntityType) Val() string {
	return [...]string{"undefined", "account", "contact", "contactaddress", "user"}[et]
}

// EntityName the value used in html template
func (et EntityType) EntityName() string {
	return [...]string{"Undefined", "Account", "Contact", "ContactAddress", "User"}[et]
}

// ParseEntityType ...
func ParseEntityType(s string) EntityType {
	switch s {
	case EntityTypeAccount.Val():
		return EntityTypeAccount
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
