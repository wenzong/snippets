package helper

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func NewTestConfig(t *testing.T) (*viper.Viper, func()) {
	var (
		mysqlEndpoint = "localhost:3306"
		err           error
		cleanFns      []func()
	)

	if _, ok := os.LookupEnv("CI"); !ok {
		ctx := context.Background()

		mysql := Container(ctx, MySQLContainerReq)

		cleanFns = append(cleanFns, func() { mysql.Terminate(ctx) })

		mysqlEndpoint, err = mysql.PortEndpoint(ctx, "3306", "")
		assert.NoError(t, err)
	}

	CreateDB(t, mysqlEndpoint, "default")
	cleanFns = append(cleanFns, func() { DropDB(t, mysqlEndpoint, "default") })

	dbParams := url.Values{}
	dbParams.Set("parseTime", "true")
	dbParams.Set("charset", "utf8mb4")
	dbParams.Set("multiStatements", "true")
	config := map[string]interface{}{
		"db": map[string]interface{}{
			"default": map[string]interface{}{
				"dsn": fmt.Sprintf("root:pass@tcp(%s)/default?%s", mysqlEndpoint, dbParams.Encode()),
			},
		},
	}

	v := viper.New()
	v.MergeConfigMap(config)
	v.MergeInConfig()

	return v, func() {
		for _, fn := range cleanFns {
			defer fn()
		}
	}
}

func CreateDB(t *testing.T, endpoint string, databases ...string) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(%s)/", endpoint))
	assert.NoError(t, err)
	defer db.Close()

	for _, database := range databases {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", database))
		assert.NoError(t, err)
	}
}

func DropDB(t *testing.T, endpoint string, databases ...string) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(%s)/", endpoint))
	assert.NoError(t, err)
	defer db.Close()

	for _, database := range databases {
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE `%s`;", database))
		assert.NoError(t, err)
	}
}
