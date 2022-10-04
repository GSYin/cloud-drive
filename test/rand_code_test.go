package test

import (
	"cloud-drive/core/utils"
	"testing"
)

func TestRandCode(t *testing.T) {
	t.Logf(utils.GenerateRandCode())
}
