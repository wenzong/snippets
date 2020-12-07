package user

import (
	"context"

	"github.com/pkg/errors"
	"github.com/wenzong/demo/infra/db"
)

type User struct {
	UserID int `db:"user_id" json:"user_id"`
}

type Service struct {
	conn   *db.DefaultConn
	second *db.SecondConn
}

func NewService(conn *db.DefaultConn, second *db.SecondConn) *Service {
	return &Service{
		conn:   conn,
		second: second,
	}
}

// context.Context is required to get information from request
func (u *Service) Get(ctx context.Context, userID int64) (user *User, err error) {
	session := u.conn.NewSession(nil)
	err = session.Select("*").
		From("user").
		Where("user_id = ?", userID).
		LoadOne(&user)

	if err != nil {
		return nil, errors.Wrapf(err, "User = %d Not Found", userID)
	}

	return user, nil
}
