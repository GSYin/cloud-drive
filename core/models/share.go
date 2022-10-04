package models

import "time"

type Share struct {
	Id                     int
	Identity               string
	UserIdentity           string
	RepositoryIdentity     string
	UserRepositoryIdentity string
	ExpiredTime            int
	ClickNum               int
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"updated"`
	DeletedAt              time.Time `xorm:"deleted"`
}

func (table Share) TableName() string {
	return "share"
}
