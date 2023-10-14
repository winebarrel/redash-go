package redash_test

import (
	"os"
	"testing"
)

var (
	testAcc = false
)

const (
	testRedashEndpoint = "http://localhost:5001"
	testRedashAPIKey   = "6nh64ZsT66WeVJvNZ6WB5D2JKZULeC2VBdSD68wt"
)

func TestMain(m *testing.M) {
	if v := os.Getenv("TEST_ACC"); v == "1" {
		testAcc = true
	}

	m.Run()
}
