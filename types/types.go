package types

//OrgType type of organization (organization/person)
type OrgType int

//
const (
	OrgTypeOrg = iota
	OrgTypePerson
)

//CustomerType type of customer relationship
type CustomerType int

//
const (
	CustomerTypeC = iota
	CustomerTypeS
	CustomerTypeP
	CustomerTypeI
)
