package tests

import (
	"github.com/l33m4n123/loadEnv"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	err := loadEnv.Load(".env")
	if err != nil {
		panic(err)
	}
	if os.Getenv("APP_ENV") != "dev" {
		t.Fail()
	}
}
