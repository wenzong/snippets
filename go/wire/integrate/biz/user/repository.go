package user

import (
	"github.com/gocraft/dbr/v2"
	"github.com/pkg/errors"
	"github.com/wenzong/demo/api/pb"
	"github.com/wenzong/demo/infra/db"
)

// MySQL Repository
type Repository struct {
	conn *db.DefaultConn
}

func NewRepository(conn *db.DefaultConn) *Repository {
	return &Repository{conn: conn}
}

// Get query MySQL table `user` by userID
func (r *Repository) Get(userID int64) (*pb.User, error) {
	return r.get(r.conn.NewSession(nil), userID)
}

// TxGet query MySQL table `user` using existing transaction by userID
func (r *Repository) TxGet(tx *dbr.Tx, userID int64) (*pb.User, error) {
	return r.get(tx, userID)
}

func (r *Repository) get(sessionOrTx dbr.SessionRunner, userID int64) (user *pb.User, err error) {
	err = sessionOrTx.
		Select("*").
		From("user").
		Where(dbr.Eq("user_id", userID)).
		LoadOne(&user)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return nil, errors.Wrapf(err, "User = %d Not Found", userID)
		} else {
			return nil, errors.Wrapf(err, "Unknown Error")
		}
	}

	return user, nil
}
