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

//CustomerTypes ...
func CustomerTypes() []CustomerType {
	return []CustomerType{CustomerTypeK, CustomerTypeL, CustomerTypeP, CustomerTypeI, CustomerTypeW}
}

func (ct CustomerType) String() string {
	return [...]string{"K", "L", "P", "I", "W"}[ct]
}

//Val the value used in html template
func (ct CustomerType) Val() string {
	return [...]string{"0", "1", "2", "3", "4"}[ct]
}

//Desc ...
func (ct CustomerType) Desc() string {
	return [...]string{"Kunde", "Lieferant", "Partner", "Interessent", "Werbung"}[ct]
}
