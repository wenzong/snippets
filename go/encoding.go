package main

import (
	"log"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	hello = "世界"
)

func main() {
	for _, enc := range unicode.All {
		log.Println(transform.String(enc.NewDecoder(), hello))
	}
	for _, enc := range simplifiedchinese.All {
		log.Println(transform.Bytes(enc.NewDecoder(), []byte(hello)))
	}
}
