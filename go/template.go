package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/spf13/viper"
)

var (
	tpl = `
[database]
{{- range $key, $value := .DB }}
[database.{{ $key }}]
dsn = "{{ $value }}"
{{- end }}

[redis]
{{- range $key, $value := .DB }}
[redis.{{ $key }}]
addr = "{{ $value}}"
{{- end }}
`
)

func main() {
	info := struct {
		DB    map[string]string
		Redis map[string]string
	}{
		DB: map[string]string{
			"default":   "root:pass@tcp(127.0.0.1:3306)/default",
			"secondary": "root:pass@tcp(127.0.0.1:3306)/secondary",
		},
		Redis: map[string]string{
			"default":   "127.0.0.1:6379",
			"secondary": "127.0.0.1:6379",
		},
	}

	t := template.Must(template.New("config").Parse(tpl))

	b := bytes.NewBufferString("")

	t.Execute(b, info)

	fmt.Println(strings.TrimSpace(b.String()))

	viper.SetConfigType("toml")
	viper.ReadConfig(b)

	fmt.Println(viper.AllSettings())
}
