package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

// MustHash will do hash(input)
//
// Example:
//    MustHash(md5.New(), strings.NewReader("string"))
//    MustHash(hmac.New(sha256.New, []byte("key")), bytes.NewReader([]byte("value")))
//    MustHash(crypto.SHA512.New(), f) // f, _ = os.Open("/path/to/file")
//
// See https://godoc.org/crypto#Hash for all available hash functions
func MustHash(h hash.Hash, r io.Reader) string {
	_, err := io.Copy(h, r)
	if err != nil {
		panic(fmt.Errorf("%T must hash failed: %w", h, err))
	}

	return hex.EncodeToString(h.Sum(nil))
}
