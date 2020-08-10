package types

//OrgType type of organization (organization/person)
type OrgType int

//
const (
	OrgTypeOrg = iota
	OrgTypePerson
)

func (ot OrgType) String() string {
	return [...]string{"ORG", "PER"}[ot]
}

//Val the value used in html template
func (ot OrgType) Val() string {
	return [...]string{"0", "1"}[ot]
}

//OrgTypes ...
func OrgTypes() []OrgType {
	return []OrgType{OrgTypeOrg, OrgTypePerson}
}

//Desc ...
func (ot OrgType) Desc() string {
	return [...]string{"Organisation", "Person"}[ot]
}
