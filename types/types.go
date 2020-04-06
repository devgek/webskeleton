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

//Desc ...
func (ot OrgType) Desc() string {
	return [...]string{"Organisation", "Person"}[ot]
}

//CustomerType type of customer relationship
type CustomerType int

//
const (
	CustomerTypeK = iota
	CustomerTypeL
	CustomerTypeP
	CustomerTypeI
	CustomerTypeW
)

func (ct CustomerType) String() string {
	return [...]string{"K", "L", "P", "I", "W"}[ct]
}

//Desc ...
func (ct CustomerType) Desc() string {
	return [...]string{"Kunde", "Lieferant", "Partner", "Interessent", "Werbung"}[ct]
}
