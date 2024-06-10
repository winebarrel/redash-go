package redash_test

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"testing/iotest"
)

var (
	testAcc       = false
	testIOErrResp = &http.Response{
		Status:     strconv.Itoa(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(iotest.ErrReader(errors.New("IO error"))),
	}
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
