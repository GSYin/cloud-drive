package test

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	err := os.Setenv("SECRETID", "SECRETID")
	if err != nil {
		return
	}
	secretId := os.Getenv("SECRETID")
	t.Errorf("secretId: %s", secretId)
}
