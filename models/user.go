package models

//User ...
type User struct {
	Entity
	Name  string `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass  string `form:"gkvPass"`
	Email string `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Admin bool   `gorm:"not null" form:"gkvAdmin"`
}
