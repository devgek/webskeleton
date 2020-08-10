package types

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
