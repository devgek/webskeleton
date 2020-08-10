package types

//
//RoleType type of user role
//
type RoleType int

//
const (
	RoleTypeUser RoleType = iota
	RoleTypeAdmin
)

//RoleTypes ...
func RoleTypes() []RoleType {
	return []RoleType{RoleTypeUser, RoleTypeAdmin}
}

//Val the value used in html template
func (rt RoleType) Val() string {
	return [...]string{"0", "1"}[rt]
}

//Desc ...
func (rt RoleType) Desc() string {
	return [...]string{"Benutzer", "Administrator"}[rt]
}

//ParseRoleType ...
func ParseRoleType(s string) RoleType {
	if s == RoleTypeAdmin.Val() {
		return RoleTypeAdmin
	}
	return RoleTypeUser
}
