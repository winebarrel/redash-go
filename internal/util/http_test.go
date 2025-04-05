package util_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go/v2/internal/util"
)

func Test_UnmarshalBody_OK(t *testing.T) {
	assert := assert.New(t)

	res := &http.Response{
		Body: io.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
	}

	var body map[string]string
	err := util.UnmarshalBody(res, &body)
	assert.NoError(err)
	assert.Equal(map[string]string{"foo": "bar"}, body)
}

func Test_UnmarshalBody_IOErr(t *testing.T) {
	assert := assert.New(t)

	res := &http.Response{
		Body: io.NopCloser(iotest.ErrReader(errors.New("IO error"))),
	}

	var body map[string]string
	err := util.UnmarshalBody(res, &body)
	assert.ErrorContains(err, "read response body failed: IO error")
}

func Test_UnmarshalBody_Err(t *testing.T) {
	assert := assert.New(t)

	res := &http.Response{
		Body: io.NopCloser(strings.NewReader(`{"foo":"`)),
	}

	var body map[string]string
	err := util.UnmarshalBody(res, &body)
	assert.Error(err)
}

func Test_CheckStatus_OK(t *testing.T) {
	assert := assert.New(t)
	tt := []int{200, 299}

	for _, t := range tt {
		res := &http.Response{
			StatusCode: t,
		}

		err := util.CheckStatus(res)
		assert.NoError(err)
	}
}

func Test_CheckStatus_Err(t *testing.T) {
	assert := assert.New(t)
	tt := []int{300, 400, 500}

	for _, t := range tt {
		res := &http.Response{
			StatusCode: t,
			Status:     fmt.Sprintf("STATUS CODE %d", t),
		}

		err := util.CheckStatus(res)
		assert.ErrorContains(err, fmt.Sprintf("HTTP status code not OK: STATUS CODE %d", t))
	}
}

func Test_CheckStatus_Err_WithBody(t *testing.T) {
	assert := assert.New(t)
	tt := []int{300, 400, 500}

	for _, t := range tt {
		res := &http.Response{
			StatusCode: t,
			Status:     fmt.Sprintf("STATUS CODE %d", t),
			Body:       io.NopCloser(strings.NewReader(`body`)),
		}

		err := util.CheckStatus(res)
		assert.ErrorContains(err, fmt.Sprintf("HTTP status code not OK: STATUS CODE %d\nbody", t))
	}
}

type testReadCloser struct {
	io.Reader
	isClosed bool
}

func (r *testReadCloser) Close() error {
	r.isClosed = true
	return nil
}

func Test_CloseResponse_OK(t *testing.T) {
	assert := assert.New(t)
	buf := strings.NewReader(`body`)
	body := &testReadCloser{Reader: buf}
	res := &http.Response{Body: body}
	util.CloseResponse(res)
	assert.True(body.isClosed)
	assert.Equal(0, buf.Len())
}

func Test_CloseResponse_WithNil(t *testing.T) {
	util.CloseResponse(nil)
}

func Test_CloseResponse_WithBodyNil(t *testing.T) {
	res := &http.Response{Body: nil}
	util.CloseResponse(res)
}

func Test_ValuesFrom_Map(t *testing.T) {
	assert := assert.New(t)
	params := map[string]string{"foo": "bar", "zoo": "baz"}
	values, err := util.URLValuesFrom(params)
	assert.NoError(err)
	assert.Equal("foo=bar&zoo=baz", values.Encode())
}

type testParams struct {
	Foo string `url:"foo"`
	Bar int    `url:"bar"`
	Zoo string `url:"zoo,omitempty"`
}

func Test_ValuesFrom_Struct(t *testing.T) {
	assert := assert.New(t)
	params := &testParams{
		Foo: "foo",
		Bar: 1,
	}
	values, err := util.URLValuesFrom(params)
	assert.NoError(err)
	assert.Equal("bar=1&foo=foo", values.Encode())
}

func Test_ValuesFrom_Err(t *testing.T) {
	assert := assert.New(t)
	_, err := util.URLValuesFrom("xxx")
	assert.ErrorContains(err, "query: Values() expects struct input. Got string")
}
