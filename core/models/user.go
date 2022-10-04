package models

import "time"

type User struct {
	Id        int       `xorm:"pk autoincr"`
	Identity  string    `xorm:"varchar(255) notnull"`
	Username  string    `xorm:"varchar(255) notnull"`
	Password  string    `xorm:"varchar(255) notnull"`
	Email     string    `xorm:"varchar(255) notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table User) TableName() string {
	return "user"
}
