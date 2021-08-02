package tests

import (
	"github.com/l33m4n123/loadEnv"
	"testing"
	"os"
)

func TestLoadSkipFileNotExist(t *testing.T) {
	err := loadEnv.Load(".eanv", true)
	if err != nil {
		t.Fail()
	}
}

func TestLoadFailNonExistingFile(t *testing.T) {
	err := loadEnv.Load(".eanv", false)
	if err == nil {
		t.Fail()
	}
}

func TestLoadExistingFile(t *testing.T) {
	loadEnv.Load(".env", false)
	if os.Getenv("APP_ENV") != "test" {
		t.Fail()
	}
}
