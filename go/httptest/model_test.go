package demo

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/request.xml")
	assert.NoError(t, err)

	req := &Request{}

	err = xml.Unmarshal(b, &req)
	assert.NoError(t, err)
}

func TestResponse(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/response.xml")
	assert.NoError(t, err)

	res := &Response{}

	err = xml.Unmarshal(b, &res)
	assert.NoError(t, err)
}
