package test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	uid := uuid.NewV4()
	t.Log(uid)
}
