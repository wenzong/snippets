package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type DefaultConn struct {
	*dbr.Connection
}
type SecondConn struct {
	*dbr.Connection
}

func NewDefaultConn(v *viper.Viper) *DefaultConn {
	config := v.Sub("db.default")

	dsn := config.GetString("dsn")

	conn, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		panic(errors.Wrapf(err, "Parse MySQL DSN failed: %s", dsn))
	}

	return &DefaultConn{conn}
}

func NewSecondConn(v *viper.Viper) (conn *SecondConn) {
	return nil
}

var ProviderSet = wire.NewSet(NewDefaultConn, NewSecondConn)
