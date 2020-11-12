package main

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	in        = "helloworld"
	fileIn    = "/etc/hosts"
	md5Out    = "fc5e038d38a57032085441e7fe7010b0"                                 // $ echo -n 'helloworld' | md5
	sha1Out   = "6adfb183a4a2c94a2f92dab5ade762a47889a5a1"                         // $ echo -n 'helloworld' | sha1sum
	sha256Out = "936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af" // $ echo -n 'helloworld' | sha256sum
)

func TestHashFunctions(t *testing.T) {
	testcases := []struct {
		output  string
		hashFn  func() hash.Hash
		hashFn2 func() hash.Hash
	}{
		{output: md5Out, hashFn: md5.New, hashFn2: crypto.MD5.New},
		{output: sha1Out, hashFn: sha1.New, hashFn2: crypto.SHA1.New},
		{output: sha256Out, hashFn: sha256.New, hashFn2: crypto.SHA256.New},
	}

	for _, tc := range testcases {
		assert.EqualValues(t, tc.output, MustHash(tc.hashFn(), strings.NewReader(in)))

		assert.EqualValues(t, tc.output, MustHash(tc.hashFn(), bytes.NewReader([]byte(in))))

		assert.EqualValues(t, tc.output, MustHash(tc.hashFn2(), strings.NewReader(in)))
	}
}

func TestHMACSHA256(t *testing.T) {
	k := []byte("key")
	p := []byte("content")

	assert.NotPanics(t, func() { MustHash(hmac.New(sha256.New, k), bytes.NewReader(p)) })
}

func TestMD5File(t *testing.T) {
	MustOpen := func(f *os.File, err error) *os.File {
		t.Helper()
		assert.NoError(t, err)

		return f
	}

	assert.NotPanics(t, func() { MustHash(md5.New(), MustOpen(os.Open(fileIn))) })
}

func TestMustHash(t *testing.T) {
	assert.Panics(t, func() { MustHash(crypto.MD4.New(), strings.NewReader(in)) }, "import golang.org/x/crypto/md4 required")
	assert.NotPanics(t, func() { MustHash(crypto.SHA224.New(), strings.NewReader(in)) })
	assert.NotPanics(t, func() { MustHash(crypto.SHA512.New(), strings.NewReader(in)) })
}
