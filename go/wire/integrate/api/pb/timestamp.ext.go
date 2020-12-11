package pb

import (
	"database/sql/driver"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Timestamp) Scan(dest interface{}) error {
	switch t := dest.(type) {
	case time.Time:
		m.Timestamp = timestamppb.New(t)
	default:
		return errors.New("Not Supported")
	}

	return nil
}

func (m *Timestamp) Value() (driver.Value, error) {
	return m.Timestamp.AsTime(), nil
}
