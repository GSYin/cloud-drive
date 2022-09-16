package models

type User struct {
	Id       int    `xorm:"pk autoincr"`
	Identity string `xorm:"varchar(255) notnull"`
	Username string `xorm:"varchar(255) notnull"`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
}

func (table User) TableName() string {
	return "user"
}
