package redash_test

import (
	"os"
	"testing"
)

var (
	testAcc            = false
	testRedashEndpoint = "http://localhost:5001"
	testRedashAPIKey   = "G1LARLeRTzoWF7asyy32Qdvken2OZ2LhzoOzwA3r"
	testRedashPgHost   = "postgres"
)

func TestMain(m *testing.M) {
	if v := os.Getenv("TEST_ACC"); v == "1" {
		testAcc = true
	}

	if v := os.Getenv("TEST_REDASH_ENDPOINT"); v != "" {
		testRedashEndpoint = v
	}

	if v := os.Getenv("TEST_REDASH_PG_HOST"); v != "" {
		testRedashPgHost = v
	}

	m.Run()
}
