package models

import "time"

type RepositoryPool struct {
	Id        int
	Identity  string
	Filehash  string
	Filename  string
	Fileext   string
	Filesize  int64
	Filepath  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
