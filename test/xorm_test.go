package test

import (
	"bytes"
	"cloud-drive/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "user:password@/database?charset=utf8mb4")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.User, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, marshal, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
