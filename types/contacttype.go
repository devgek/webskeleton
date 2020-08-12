package types

//ContactType type of customer relationship
type ContactType int

//
const (
	ContactTypeK = iota
	ContactTypeL
	ContactTypeP
	ContactTypeI
	ContactTypeW
)

//ContactTypes ...
func ContactTypes() []ContactType {
	return []ContactType{ContactTypeK, ContactTypeL, ContactTypeP, ContactTypeI, ContactTypeW}
}

func (ct ContactType) String() string {
	return [...]string{"K", "L", "P", "I", "W"}[ct]
}

//Val the value used in html template
func (ct ContactType) Val() string {
	return [...]string{"0", "1", "2", "3", "4"}[ct]
}

//Desc ...
func (ct ContactType) Desc() string {
	return [...]string{"Kunde", "Lieferant", "Partner", "Interessent", "Werbung"}[ct]
}
